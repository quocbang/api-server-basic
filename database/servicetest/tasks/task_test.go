package tasks

import (
	"github.com/labstack/echo"

	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
)

func (s *Suite) TestCreateTasks() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// task already exists.
	{
		// arrange
		s.ClearData()
		taskExistedRequest := requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:      1,
					Status:  "pending",
					Content: "build web server with high performance",
					EndTime: 48, // 4 days
				},
				{
					ID:             1, // ID already exists
					Status:         "preparing",
					ProcessPercent: 5,
					Content:        "build web server with high performance",
					EndTime:        48,
				},
			},
		}

		// act
		_, err := s.DM.Tasks().CreateTasks(ctx, taskExistedRequest)

		// assert
		assertion.Error(err)
		expected := localErr.Error{
			Code:   localErr.Code_ERR_DATA_EXISTED,
			Detail: "Key (id)=(1) already exists.",
		}
		assertion.Equal(expected, err)
	}

	// create task successfully.
	{
		// arrange
		s.ClearData()
		req := requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:      1,
					Status:  "pending",
					Content: "create great server",
					EndTime: 10,
				},
				{
					ID:             2,
					Status:         "doing",
					ProcessPercent: 20,
					Content:        "build suite test",
					EndTime:        30,
				},
			},
		}

		// act
		reply, err := s.DM.Tasks().CreateTasks(ctx, req)

		// assert
		expected := requests.CreateTaskReply{
			RowsAffected: int64(2),
		}
		assertion.NoError(err)
		assertion.Equal(expected, reply)
	}
}

func (s *Suite) TestGetTasks() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// get task with id.
	{
		// arrange
		s.ClearData()
		createTask := requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:             1,
					Status:         "doing",
					ProcessPercent: 10,
					Content:        "get task with id",
					EndTime:        1,
				},
			},
		}
		_, err := s.DM.Tasks().CreateTasks(ctx, createTask)
		assertion.NoError(err)

		req := requests.GetTaskRequest{
			ID: 1,
		}

		// act
		reply, err := s.DM.Tasks().GetTasks(ctx, req)

		// assetion
		assertion.NoError(err)
		assertion.NotNil(reply)
		assertion.Equal(1, len(reply.Tasks))
		if task := reply.Tasks; len(task) == 1 {
			assertion.Equal(int32(1), task[0].ID)
			assertion.Equal("doing", task[0].Status)
			assertion.Equal(int16(10), task[0].ProcessPercent)
			assertion.Equal("get task with id", task[0].Content)
		}
	}

	// get task without id
	{
		// arrange
		s.ClearData()
		createTask := requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:             1,
					Status:         "doing",
					ProcessPercent: 10,
					Content:        "get task without id",
					EndTime:        1,
				},
				{
					ID:             2,
					Status:         "complete",
					ProcessPercent: 10,
					Content:        "get task without id",
					EndTime:        1,
				},
			},
		}
		_, err := s.DM.Tasks().CreateTasks(ctx, createTask)
		assertion.NoError(err)

		// act
		reply, err := s.DM.Tasks().GetTasks(ctx, requests.GetTaskRequest{})

		// assetion
		assertion.NoError(err)
		assertion.NotNil(reply)
		assertion.Equal(2, len(reply.Tasks))
	}

	// get task but not stored any task in database.
	{
		// arrange
		s.ClearData()

		// act
		_, err := s.DM.Tasks().GetTasks(ctx, requests.GetTaskRequest{ID: 1})

		// assetion
		expected := localErr.Error{Detail: localErr.ErrDataNotFound.Error()}
		assertion.Error(err)
		assertion.Equal(expected, err)
	}
}

func (s *Suite) TestUpdateTask() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// update successfully.
	{
		// arrange
		s.ClearData()
		createTask := requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:             1,
					Status:         "doing",
					Content:        "step update task case",
					ProcessPercent: 10,
					EndTime:        10,
				},
			},
		}

		_, err := s.DM.Tasks().CreateTasks(ctx, createTask)
		assertion.NoError(err)

		updateTask := requests.UpdateTaskRequest{
			ID:             1,
			Status:         "complete",
			ProcessPercent: 100,
		}

		// act
		reply, err := s.DM.Tasks().UpdateTask(ctx, updateTask)

		// assert
		expected := requests.UpdateTaskReply{
			RowsAffected: 1,
		}
		assertion.NoError(err)
		assertion.Equal(expected, reply)
	}
}

func (s *Suite) TestDeleteTasks() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// delete task successfully.
	{
		// arrange
		s.ClearData()
		createTask := requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:             1,
					Status:         "preparing",
					ProcessPercent: 0,
					Content:        "first task",
					EndTime:        10,
				},
				{
					ID:             2,
					Status:         "preparing",
					ProcessPercent: 0,
					Content:        "seccond task",
					EndTime:        10,
				},
			},
		}

		s.DM.Tasks().CreateTasks(ctx, createTask)

		req := requests.DeleteTaskRequest{
			IDs: []int32{1, 2},
		}

		// act
		reply, err := s.DM.Tasks().DeleteTasks(ctx, req)

		// assert
		expected := requests.DeleteTaskReply{
			RowsAffected: 2,
		}
		assertion.NoError(err)
		assertion.Equal(expected, reply)
	}
}
