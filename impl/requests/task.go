package requests

import "time"

type GetTaskRequest struct {
	ID int32 `json:"id"`
}

func (gsr GetTaskRequest) IsUse() bool {
	return gsr != GetTaskRequest{}
}

type GetTaskReply struct {
	Tasks []TaskDetail
}

type TaskDetail struct {
	ID             int32
	Status         string
	ProcessPercent int16
	Content        string
	EndTime        time.Time
	Updated        int64
	Created        int64
}

type CreateTaskRequest struct {
	Tasks []CreateTaskDetail `validate:"required,dive" json:"tasks"`
}

type CreateTaskDetail struct {
	ID             int32  `json:"id" validate:"required,gte=1"`
	Status         string `json:"status"`
	ProcessPercent int16  `json:"process_percent"`
	Content        string `validate:"required" json:"content"`
	// units is hour.
	EndTime time.Duration `validate:"required,gte=0" json:"end_time"`
}

type CreateTaskReply struct {
	RowsAffected RowsAffected
}

type UpdateTaskRequest struct {
	ID             int32  `json:"id" validate:"required,gte=1"`
	Status         string `json:"status"`
	ProcessPercent int16  `json:"process_percent"`
}

type UpdateTaskReply struct {
	RowsAffected RowsAffected
}

type DeleteTaskRequest struct {
	IDs []int32 `json:"ids" validate:"required,dive,gte=1"`
}

type DeleteTaskReply struct {
	RowsAffected RowsAffected
}
