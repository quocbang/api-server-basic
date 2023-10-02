package tasks

import (
	"github.com/labstack/echo"

	"github.com/quocbang/api-server-basic/impl/service"
	"github.com/quocbang/api-server-basic/middleware/authorization"
)

func RegisterTasksAPI(e *echo.Echo, serivePath string, taskService service.TaskService) {
	// ListTasks are get all tasks.
	//
	// DO:
	//  - list task with request ID or,
	//  - list all tasks if request ID not found.
	e.GET(serivePath, taskService.ListTasks, authorization.Authorization)
	// CreateTasks is create new task.
	//
	// DO:
	//  - create tasks.
	//
	// Required:
	//	- id
	//  - status
	//  - processpercent
	//  - content
	//  - endtime
	e.POST(serivePath, taskService.CreateTasks, authorization.Authorization)
	// UpdateTask is update stage of task.
	//
	// DO:
	//	- update stage of task
	//
	// Required:
	//  - id
	//  - status
	//  - process_percent
	e.PATCH(serivePath, taskService.UpdateTask, authorization.Authorization)
	// DeleteTasks are delete given tasks ids
	//
	// DO:
	//	- delete tasks
	//
	// Required:
	//  - ids
	e.DELETE(serivePath, taskService.DeleteTasks, authorization.Authorization)
}
