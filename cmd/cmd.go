package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v3"

	"github.com/quocbang/api-server-basic/api"
	"github.com/quocbang/api-server-basic/config"
	"github.com/quocbang/api-server-basic/database/impl"
	serviceImpl "github.com/quocbang/api-server-basic/impl"
	"github.com/quocbang/api-server-basic/middleware/logging"
	"github.com/quocbang/api-server-basic/utils/roles"
)

type Config struct {
	Options    config.Options
	TLSOptions config.TLSOptions
}

func parseFlags() *Config {
	var conf Config
	configurations := []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Server Configuration",
			LongDescription:  "Server Configuration",
			Options:          &conf.Options,
		},
		{
			ShortDescription: "TLS handshake Configuration",
			LongDescription:  "TLS handshake Configuration",
			Options:          &conf.TLSOptions,
		},
	}

	parser := flags.NewParser(nil, flags.Default)
	for _, opt := range configurations {
		if _, err := parser.AddGroup(opt.ShortDescription, opt.LongDescription, opt.Options); err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok && fe.Type == flags.ErrHelp {
			code = 0
		}
		os.Exit(code)
	}
	return &conf
}

func loadConfig(filePath string) (*config.Configs, error) {
	var conf config.Configs

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func RunServer() {
	server := echo.New()

	// parse flags.
	flags := parseFlags()

	// custom logger.
	logger, err := logging.InitializeLogger(flags.Options.DevMode)
	if err != nil {
		log.Fatalf("failed to initialize logger, error: %v", err)
	}
	defer logger.Sync()
	server.Use(logging.CustomLogger(logger))

	// load config in config.yaml file
	configs, err := loadConfig(flags.Options.Config)
	if err != nil {
		log.Fatal(err)
	}

	// initialize permission.
	roles.InitializePermission(configs.Roles)

	// data manager and initialize the database.
	dm, err := impl.NewDataManager(impl.DBConfig{
		Address:  configs.Postgres.Address,
		Port:     configs.Postgres.Port,
		DBName:   configs.Postgres.Name,
		DBSchema: configs.Postgres.Schema,
		UserName: configs.Postgres.UserName,
		Password: configs.Postgres.Password,
	}, impl.RedisConfig{
		Address:  configs.Redis.Address,
		Password: configs.Redis.Password,
		Database: configs.Redis.Database,
	}, impl.DemoAccount{
		Admin: configs.DemoAccount,
	})
	if err != nil {
		log.Fatal(err)
	}

	// register service.
	service := serviceImpl.RegisterService(dm)

	// Enable CORS middleware
	server.Use(middleware.CORS())

	// config API.
	api.RegisterAPI(server, service)

	address := fmt.Sprintf("%s:%d", flags.Options.Host, flags.Options.Port)
	if !flags.TLSOptions.UseTLS() {
		log.Println("Server handshake without TLS")
		server.Logger.Fatal(server.Start(address))
	} else {
		log.Println("Server handshake with TLS")
		server.Logger.Fatal(server.StartTLS(address, flags.TLSOptions.Cert, flags.TLSOptions.Key))
	}
}
