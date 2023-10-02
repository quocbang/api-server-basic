package testutil

import (
	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/mocks"
	"github.com/quocbang/api-server-basic/impl"
	"github.com/quocbang/api-server-basic/impl/service"
	"github.com/stretchr/testify/mock"
)

// mockStruct should is mock.Mock{}
// mockInterface should is interface that contian your test method.
type Scripts[mockStruct, mockInterface any] struct {
	MethodName string
	Input      []any
	Output     []any
	// should be return interface
	ParentReturn func(mockStruct) mockInterface
	IsParent     bool
}

// mockStruct should is *mock.Mock{}, mockInterface should is interface contained method.
func CreateMockService[mockStruct, mockInterface any](s ...Scripts[mockStruct, mockInterface]) *mocks.Services {
	var (
		mockStructs any
	)
	mock := &mock.Mock{}

	for i := len(s) - 1; i >= 0; i-- {
		if !s[i].IsParent {
			mock.On(s[i].MethodName, s[i].Input...).Return(s[i].Output...)
			mockStructs = mock
		} else {
			mock.On(s[i].MethodName).Return(s[i].ParentReturn(mockStructs.(mockStruct)))
		}
	}
	return &mocks.Services{Mock: mock}
}

// NewMock return all service that was registered.
func NewMock[mockStruct, mockInterface any](s []Scripts[mockStruct, mockInterface]) service.DataManagerService {
	var service *mocks.Services
	if len(s) != 0 {
		// create mock scripts.
		service = CreateMockService[mockStruct, mockInterface](s...)
	}
	dm := registerMockDB(service)
	return registerService(dm)
}

// RegisterService is register and list all service.
func registerService(dm database.DataManager) service.DataManagerService {
	return &impl.ServiceConfig{
		DM: dm,
	}
}

func registerMockDB(m *mocks.Services) database.DataManager {
	return m
}
