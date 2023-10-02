package suite

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"gorm.io/gorm"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/impl"
)

var postGresTest struct {
	DBSchema   string `long:"db_schema" description:"The test DB schema" env:"DB_SCHEMA"`
	DBAddress  string `long:"db_address" description:"The test DB address" env:"DB_ADDRESS"`
	DBPort     int    `long:"db_port" description:"The test DB port" env:"DB_PORT"`
	DBUsername string `long:"db_username" description:"The test DB username" env:"DB_USERNAME"`
	DBPassword string `long:"db_password" description:"The test DB password" env:"DB_PASSWORD"`
	DBDatabase string `long:"db_database" description:"The test DB database" env:"DB_DATABASE"`
}

var redisTest struct {
	RedisAddress  string `long:"redis_address" description:"The test Redis address" env:"REDIS_ADDRESS"`
	RedisPassword string `long:"redis_password" description:"The test Redis password " env:"REDIS_PASSWORD"`
	RedisDatabase int    `long:"redis_databse" description:"The test Redis databse" env:"REDIS_DATABASE"`
}

func InitializeDB() (dm database.DataManager, db *gorm.DB, err error) {
	if err = parseFlags(); err != nil {
		return // return when parse failed.
	}

	dm, db, err = newTestDataManager()
	return
}

func parseFlags() error {
	configurations := []swag.CommandLineOptionsGroup{
		{
			LongDescription:  "PostGres Configuration",
			ShortDescription: "PostGres Configuration",
			Options:          &postGresTest,
		},
		{
			LongDescription:  "Redis Configuration",
			ShortDescription: "Redis Configuration",
			Options:          &redisTest,
		},
	}

	parse := flags.NewParser(nil, flags.IgnoreUnknown)
	//  NewParser(databaseParse, flags.IgnoreUnknown)
	for _, opt := range configurations {
		if _, err := parse.AddGroup(opt.LongDescription, opt.LongDescription, opt.Options); err != nil {
			return err
		}
	}

	if _, err := parse.Parse(); err != nil {
		return fmt.Errorf("failed to parse postgres flags")
	}

	return nil
}

func newTestDataManager() (dm database.DataManager, db *gorm.DB, err error) {
	pgConfig := impl.DBConfig{
		Address:  postGresTest.DBAddress,
		Port:     postGresTest.DBPort,
		UserName: postGresTest.DBUsername,
		Password: postGresTest.DBPassword,
		DBName:   postGresTest.DBDatabase,
		DBSchema: postGresTest.DBSchema,
	}
	redisConfig := impl.RedisConfig{
		Address:  redisTest.RedisAddress,
		Password: redisTest.RedisPassword,
		Database: redisTest.RedisDatabase,
	}

	dm, err = impl.NewDataManager(pgConfig, redisConfig, impl.DemoAccount{})
	if err != nil {
		return nil, nil, err
	}

	db, err = impl.NewDB(pgConfig)
	if err != nil {
		return nil, nil, err
	}
	return
}
