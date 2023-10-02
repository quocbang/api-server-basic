package impl

import (
	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/impl/service"
)

// ServiceConfig is config necessary data for each method of service.
type ServiceConfig struct {
	DM database.DataManager
	// more service config here....
}

// RegisterService is register and list all service.
func RegisterService(dm database.DataManager) service.DataManagerService {
	return &ServiceConfig{
		DM: dm,
	}
}
