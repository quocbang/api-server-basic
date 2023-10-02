package suite

import (
	"log"
	"os"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/orm/models"
)

type SuiteConfig struct {
	*suite.Suite
	DB *gorm.DB
	DM database.DataManager

	relativeModels       []models.Models
	clearDataForEachTest bool
}

type NewSuiteParameters struct {
	// just do test or clear data in those models also data in database.
	RelativeModels []models.Models
	// clear all data in RelativeModels.
	ClearDataForEachTest bool
}

func NewSuiteTest(params NewSuiteParameters) *SuiteConfig {
	return &SuiteConfig{
		relativeModels:       params.RelativeModels,
		Suite:                &suite.Suite{},
		clearDataForEachTest: params.ClearDataForEachTest,
	}
}

var logger *zap.Logger

func (s *SuiteConfig) SetupSuite() {
	ranDomSeed := fake.UUID()
	field := []zap.Field{zap.String("random_seed", ranDomSeed)}
	logger.Info("sedd", field...)
	// initialize database.
	dm, db, err := InitializeDB()
	if err != nil {
		os.Exit(1)
	}
	s.DB = db
	s.DM = dm
}

func (suite *SuiteConfig) TearDownSuite() {
	err := suite.DM.Close()
	if err != nil {
		os.Exit(1)
	}
}

func (s *SuiteConfig) ClearData() error {
	for _, m := range s.relativeModels {
		if err := s.DB.Where("1=1").Delete(m).Error; err != nil {
			return err
		}
	}
	return nil
}

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}
