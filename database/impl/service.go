package impl

import (
	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/services/tasks"
	"github.com/quocbang/api-server-basic/database/services/users"
)

func (dm *DM) Users() database.Users {
	return users.NewService(dm.db, dm.redis)
}

func (dm *DM) Tasks() database.Tasks {
	return tasks.NewService(dm.db)
}
