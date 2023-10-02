package impl

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/logging"
	"github.com/quocbang/api-server-basic/database/orm/models"
	"github.com/quocbang/api-server-basic/database/services/users"
	postgresErr "github.com/quocbang/api-server-basic/database/utils/gorm/postgres"
	"github.com/quocbang/api-server-basic/utils/roles"
)

type DM struct {
	db    *gorm.DB
	redis *redis.Client
}

type DBConfig struct {
	Address  string
	Port     int
	DBName   string
	DBSchema string
	UserName string
	Password string
}

func NewDB(config DBConfig) (*gorm.DB, error) {
	stringConnect := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Address, config.UserName, config.Password, config.DBName, config.Port)

	if config.DBSchema != "" {
		stringConnect += fmt.Sprintf(" search_path=%s", config.DBSchema)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: stringConnect,
	}), &gorm.Config{Logger: logging.NewLogger()})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDM(dbConfig DBConfig, redisConfig RedisConfig) (*DM, error) {
	// connect with POstgres database.
	db, err := NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	// connect with Redis database.
	redis, err := NewRedis(redisConfig)
	if err != nil {
		return nil, err
	}
	return &DM{
		db:    db,
		redis: redis,
	}, nil
}

type DemoAccount struct {
	Admin []string
}

func NewDataManager(dbConfig DBConfig, redisConfig RedisConfig, demo DemoAccount) (database.DataManager, error) {
	// connect database
	dm, err := NewDM(dbConfig, redisConfig)
	if err != nil {
		return nil, err
	}

	// migrate table here.
	if err := dm.maybeMigrate(); err != nil {
		return nil, err
	}

	// create first leader and admin.
	// TODO: should create by devops in real project.
	dm.createDemoAccount(demo)

	return dm, nil
}

func (dm *DM) maybeMigrate() error {
	models := models.ListModels()
	dst := []any{}
	for _, model := range models {
		dst = append(dst, model)
	}

	if err := dm.db.AutoMigrate(dst...); err != nil {
		return err
	}
	return nil
}

func (dm *DM) Close() error {
	db, err := dm.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// createDemoAccount only for demo.
func (dm *DM) createDemoAccount(acc DemoAccount) {
	if len(acc.Admin) == 0 {
		log.Println("not found demo account")
		return
	}

	users := make([]models.Users, len(acc.Admin))
	for idx, v := range acc.Admin {
		account := strings.Split(v, "|")
		if len(account) != 2 {
			log.Printf("wrong format expected: xxx@xxx|your_password but got: %v \n", v)
		}
		password, err := hashPassword(account[1])
		if err != nil {
			log.Fatalf("create password for demo account failed, error: %v", err)
		}
		users[idx] = models.Users{
			Email:    account[0],
			Password: password,
			Roles:    pq.Int64Array{int64(roles.Roles_ADMINISTRATOR)}, // admin.
		}
	}

	reply := dm.db.Create(&users)
	if err := reply.Error; err != nil {
		if postgresErr.ErrorIs(err, postgresErr.UniqueViolation) {
			log.Printf("account already exists")
			return
		}
		log.Fatalf("create demo account failed, error: %v", err)
	}

	log.Println("create demo user successfully")
}

// HashPassword hashes the password using bcrypt
func hashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

type RedisConfig struct {
	Address  string
	Password string
	Database int
}

// NewRedis is connect to redis databse.
func NewRedis(config RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Database,
	})

	if err := rdb.Ping().Err(); err != nil {
		return nil, err
	}

	users.InitializeCheckBlackList(rdb)
	return rdb, nil
}
