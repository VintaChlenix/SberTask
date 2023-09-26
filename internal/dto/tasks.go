package dto

import (
	"SberTask/internal/model"
)

type CreateTaskRequest struct {
	Task model.Task `json:"task"`
}

type CreateTaskResponse struct {
	TaskID int `json:"task_id"`
}

type GetTaskResponse struct {
	Task model.Task `json:"task"`
}

type UpdateTaskRequest struct {
	Task model.Task `json:"task"`
}

type GetTasksRequest struct {
	Pagination model.Pagination `json:"pagination"`
}

type GetTasksResponse struct {
	Tasks []model.Task `json:"tasks"`
}
