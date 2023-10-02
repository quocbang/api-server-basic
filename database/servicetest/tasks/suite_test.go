package tasks

import (
	"testing"

	s "github.com/stretchr/testify/suite"

	"github.com/quocbang/api-server-basic/database/orm/models"
	"github.com/quocbang/api-server-basic/database/servicetest/internal/suite"
)

type Suite struct {
	suite.SuiteConfig
}

// NewSuite create all necessary.
func NewSuite() *Suite {
	s := suite.NewSuiteTest(suite.NewSuiteParameters{
		RelativeModels:       []models.Models{&models.Tasks{}},
		ClearDataForEachTest: true,
	})
	return &Suite{
		SuiteConfig: *s,
	}
}

func TestSuite(t *testing.T) {
	// Run the test suite using suite.Run.
	s.Run(t, NewSuite())
}
