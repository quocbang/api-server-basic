package suite

import (
	"log"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/quocbang/api-server-basic/utils/roles"
)

type BasicSuite struct {
	*suite.Suite
	E     *echo.Echo
	roles map[string][]string
}

type Setup struct {
	Roles map[string][]string
}

func NewSuite(s Setup) *BasicSuite {
	field := []zap.Field{
		zap.String("random seed", fake.UUID()),
	}
	logger.Info("start service test", field...)
	return &BasicSuite{
		roles: s.Roles,
		Suite: &suite.Suite{},
	}
}

func (b *BasicSuite) SetupSuite() {
	// setup roles.
	roles.InitializePermission(b.roles)
	b.E = echo.New()
}

func (b *BasicSuite) TearDownSuite() {

}

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}
