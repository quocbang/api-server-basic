package tasks

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo"
	"gorm.io/gorm"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/orm/models"
	"github.com/quocbang/api-server-basic/database/utils/gorm/postgres"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
)

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) database.Tasks {
	return &service{
		db: db,
	}
}

// GetTasks is get one or all tasks stored in database.
func (s *service) GetTasks(ctx echo.Context, req requests.GetTaskRequest) (requests.GetTaskReply, error) {
	// get task with id.
	if req.IsUse() {
		tasks := models.Tasks{}
		reply := s.db.Where(`id = ?`, req.ID).Take(&tasks)
		if err := reply.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return requests.GetTaskReply{}, localErr.Error{
					Detail: localErr.ErrDataNotFound.Error(),
				}
			}
			return requests.GetTaskReply{}, err
		}

		return requests.GetTaskReply{
			Tasks: []requests.TaskDetail{
				{
					ID:             tasks.ID,
					Status:         tasks.Status,
					ProcessPercent: tasks.ProcessPercent,
					Content:        tasks.Content,
					EndTime:        time.Unix(tasks.EndTime, 0),
					Created:        tasks.Created,
					Updated:        tasks.Updated,
				},
			},
		}, nil
	}

	// get tasks without id.
	tasks := []models.Tasks{}
	reply := s.db.Find(&tasks)
	if err := reply.Error; err != nil {
		return requests.GetTaskReply{}, err
	}

	result := requests.GetTaskReply{Tasks: make([]requests.TaskDetail, len(tasks))}
	for idx, task := range tasks {
		result.Tasks[idx] = requests.TaskDetail{
			ID:             task.ID,
			Status:         task.Status,
			ProcessPercent: task.ProcessPercent,
			Content:        task.Content,
			EndTime:        time.Unix(task.EndTime, 0),
			Updated:        task.Updated,
			Created:        task.Created,
		}
	}

	return requests.GetTaskReply{
		Tasks: result.Tasks,
	}, nil
}

// CreateTasks is create one or multi-tasks into database.
func (s *service) CreateTasks(ctx echo.Context, req requests.CreateTaskRequest) (requests.CreateTaskReply, error) {
	tasks := make([]models.Tasks, len(req.Tasks))
	for idx, task := range req.Tasks {
		tasks[idx] = models.Tasks{
			ID:             task.ID,
			Status:         task.Status,
			ProcessPercent: task.ProcessPercent,
			Content:        task.Content,
			EndTime:        time.Now().Add(task.EndTime * time.Hour).Unix(), // task time * time.Hour
		}
	}

	reply := s.db.Create(&tasks)
	if err := reply.Error; err != nil {
		if postgres.ErrorIs(err, postgres.UniqueViolation) {
			if e, ok := err.(*pgconn.PgError); ok {
				return requests.CreateTaskReply{}, localErr.Error{
					Code:   localErr.Code_ERR_DATA_EXISTED,
					Detail: e.Detail,
				}
			}
		}
		return requests.CreateTaskReply{}, err
	}

	return requests.CreateTaskReply{RowsAffected: reply.RowsAffected}, nil
}

// UpdateTask is update stage of task.
func (s *service) UpdateTask(ctx echo.Context, req requests.UpdateTaskRequest) (requests.UpdateTaskReply, error) {
	reply := s.db.Model(models.Tasks{}).Where("id=?", req.ID).Updates(models.Tasks{Status: req.Status, ProcessPercent: req.ProcessPercent})
	if err := reply.Error; err != nil {
		return requests.UpdateTaskReply{}, err
	}
	return requests.UpdateTaskReply{RowsAffected: reply.RowsAffected}, nil
}

// DeleteTasks is delete with support delete multi-tasks.
func (s *service) DeleteTasks(ctx echo.Context, req requests.DeleteTaskRequest) (requests.DeleteTaskReply, error) {
	reply := s.db.Where(`id in ?`, req.IDs).Delete(&models.Tasks{})
	if err := reply.Error; err != nil {
		return requests.DeleteTaskReply{}, err
	}
	return requests.DeleteTaskReply{RowsAffected: reply.RowsAffected}, nil
}
