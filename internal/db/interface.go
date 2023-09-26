package db

import (
	"SberTask/internal/model"
	"context"
)

type DB interface {
	CreateTask(ctx context.Context, task model.Task) (int, error)
	SelectTask(ctx context.Context, taskID int) (*model.Task, error)
	UpdateTask(ctx context.Context, taskID int, task model.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	SelectTasks(ctx context.Context, pagination model.Pagination) ([]model.Task, error)
}
