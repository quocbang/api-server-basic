package users

import (
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/orm/models"
	"github.com/quocbang/api-server-basic/database/utils/gorm/postgres"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
	"github.com/quocbang/api-server-basic/utils/roles"
)

var (
	secretKey       = os.Getenv("SECRET_KEY")
	blackListConfig *service
)

type service struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewService(db *gorm.DB, redis *redis.Client) database.Users {
	return service{db: db, redis: redis}
}

// CreateUser definition.
func (s service) CreateAccount(ctx echo.Context, req requests.CreateAccountRequest) (requests.CreateAccountReply, error) {
	// generate password.
	pwd, err := hashPassword(req.Password)
	if err != nil {
		return requests.CreateAccountReply{}, err
	}

	// create roles of user.
	roles := make(pq.Int64Array, len(req.Roles))
	for idx, role := range req.Roles {
		roles[idx] = int64(role)
	}

	// create user.
	user := models.Users{
		Email:    req.Email,
		Password: pwd,
		Roles:    roles,
	}
	reply := s.db.Create(&user)
	if err := reply.Error; err != nil {
		if postgres.ErrorIs(err, postgres.UniqueViolation) {
			return requests.CreateAccountReply{}, localErr.Error{
				Code:   localErr.Code_ERR_DATA_EXISTED,
				Detail: "user already exists",
			}
		}
		return requests.CreateAccountReply{}, err
	}

	return requests.CreateAccountReply{RowsAffected: reply.RowsAffected}, nil
}

func (s service) DeleteAccounts(ctx echo.Context, req requests.DeleteAccountRequest) (requests.DeleteAccountReply, error) {
	reply := s.db.Where(`email in ?`, req.Emails).Delete(&models.Users{})
	if err := reply.Error; err != nil {
		return requests.DeleteAccountReply{}, err
	}
	return requests.DeleteAccountReply{RowsAffected: reply.RowsAffected}, nil
}

func (s service) Login(ctx echo.Context, req requests.LoginRequest) (requests.LoginReply, error) {
	// get hashed password.
	userInfo, err := getUser(ctx, req.Email, s.db)
	if err != nil {
		return requests.LoginReply{}, err
	}

	// compare given password and stored password.
	if err := bcrypt.CompareHashAndPassword(userInfo.Password, []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Passwords don't match, handle the invalid login
			return requests.LoginReply{}, localErr.Error{
				Code:   localErr.Code_WRONG_PASSWORD,
				Detail: "wrong password",
			}
		} else {
			// Handle the error
			return requests.LoginReply{}, err
		}
	}

	// create JWT.
	token, err := s.generateJWT(ctx, userInfo)
	if err != nil {
		return requests.LoginReply{}, err
	}

	return requests.LoginReply{
		Token: token,
	}, nil
}

func getUser(ctx echo.Context, email string, db *gorm.DB) (models.Users, error) {
	var user models.Users

	if err := db.Where("email=?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Users{}, localErr.Error{
				Code:   localErr.Code_ERR_DATA_NOT_FOUND,
				Detail: "user not found",
			}
		}
		return models.Users{}, err
	}

	return user, nil
}

// HashPassword hashes the password using bcrypt
func hashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

type JwtCustomClaims struct {
	Email      string        `json:"email"`
	ExpiryTime time.Time     `json:"expiry_time"`
	Roles      []roles.Roles `json:"roles"`
	jwt.StandardClaims
}

// generate JWT.
func (s *service) generateJWT(ctx echo.Context, userInfo models.Users) (string, error) {
	if secretKey == "" {
		return "", localErr.Error{
			Detail: "secret key not found",
		}
	}

	r := make([]roles.Roles, len(userInfo.Roles))
	for idx, role := range userInfo.Roles {
		r[idx] = roles.Roles(role)
	}
	claims := &JwtCustomClaims{
		Email:      userInfo.Email,
		ExpiryTime: time.Now().Add(time.Hour * 8),
		Roles:      r,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Logout implementation.
func (s service) SignOut(ctx echo.Context, token string, timeRemaining time.Duration) error {
	// check black list whether logout or not.
	if err := s.checkBlackList(token); err != nil {
		return err
	}
	reply := s.redis.Set(token, nil, timeRemaining)
	if err := reply.Err(); err != nil {
		return err
	}
	return nil
}

// check black list.
func (s *service) checkBlackList(token string) error {
	reply := s.redis.Get(token)
	// return reply.Err() == nil then the token key was stored
	if err := reply.Err(); err != nil {
		return nil
	}
	return localErr.Error{
		Code:   localErr.Code_TOKEN_BLOCKED,
		Detail: "this token was logged out",
	}
}

// check black list in external package without database connection.
func IsTokenBLocked(token string) bool {
	if blackListConfig != nil {
		if err := blackListConfig.checkBlackList(token); err != nil {
			return true // token was stored in redis => blocked.
		}
		return false // token not found in black list.
	}
	return true // should connect fail.
}

func InitializeCheckBlackList(redis *redis.Client) {
	blackListConfig = &service{redis: redis}
}
