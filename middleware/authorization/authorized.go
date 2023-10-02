package authorization

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"

	"github.com/quocbang/api-server-basic/database/services/users"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/utils/roles"
)

const (
	// AuthorizationKey use for authentication of each secure services.
	AuthorizationKey = "x-server-auth-key"
	UserPrincipalKey = "user-principal"
)

var secretKey = os.Getenv("SECRET_KEY")

// Principal principal.
type Principal struct {
	Email      string        `json:"email"`
	Roles      []roles.Roles `json:"roles"`
	ExpiryTime time.Time     `json:"expiry_time"`
}

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get(AuthorizationKey)
		principal, err := verifyToken(token)
		if err != nil {
			if localErr.ErrorIs(err, localErr.Code_MISSING_TOKEN) {
				return ctx.JSON(http.StatusNetworkAuthenticationRequired, err)
			}
			return ctx.JSON(http.StatusForbidden, err)
		}
		// check black list.
		if users.IsTokenBLocked(token) {
			return ctx.JSON(http.StatusForbidden, localErr.Error{
				Code:   localErr.Code_TOKEN_BLOCKED,
				Detail: "token blocked",
			})
		}
		ctx.Set(UserPrincipalKey, principal)
		return next(ctx)
	}
}

// VerifyToken is check the validity of the token.
func verifyToken(token string) (*Principal, error) {
	if token == "" {
		return nil, localErr.Error{
			Code:   localErr.Code_MISSING_TOKEN,
			Detail: "authorize token required",
		}
	}
	if secretKey == "" {
		return nil, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "secret key not found",
		}
	}

	JWTToken, err := jwt.ParseWithClaims(token, &users.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, localErr.Error{
				Code:   localErr.Code_FORBIDDEN,
				Detail: "invalid signing method",
			}
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: fmt.Sprintf("failed to parse token: %v", err),
		}
	}

	var principal Principal
	if claims, ok := JWTToken.Claims.(*users.JwtCustomClaims); !ok || !JWTToken.Valid {
		return nil, localErr.Error{
			Code:   localErr.Code_INVALID_TOKEN,
			Detail: "token invalid",
		}
	} else {
		if claims.ExpiryTime.Before(time.Now()) {
			return nil, localErr.Error{
				Code:   localErr.Code_TOKEN_EXPIRED,
				Detail: "the token was expired",
			}
		}
		principal = Principal{
			Email:      claims.Email,
			Roles:      claims.Roles,
			ExpiryTime: claims.ExpiryTime,
		}
	}

	return &principal, nil
}
