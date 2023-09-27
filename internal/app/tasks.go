package app

import (
	"SberTask/internal/dto"
	"SberTask/pkg/handlers"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (a *App) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request dto.CreateTaskRequest
	if err := handlers.UnmarshalJSON(r, &request.Task); err != nil {
		a.log.Errorf("failed to unmarshal request json: %v", err)
		handlers.RenderBadRequest(w, err)
		return
	}

	response, err := a.createTaskHandler(ctx, request)
	if err != nil {
		a.log.Errorf(err.Error())
		handlers.RenderInternalError(w, err)
		return
	}

	handlers.RenderJSON(w, response)
}

func (a *App) createTaskHandler(ctx context.Context, request dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	taskID, err := a.db.CreateTask(ctx, request.Task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	return &dto.CreateTaskResponse{TaskID: taskID}, nil
}

func (a *App) ReadTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rawTaskID := chi.URLParam(r, "task_id")
	taskID, err := strconv.Atoi(rawTaskID)
	if err != nil {
		a.log.Errorf("failed to get task ID: %v", err)
		handlers.RenderBadRequest(w, err)
		return
	}

	response, err := a.readTaskHandler(ctx, taskID)
	if err != nil {
		a.log.Errorf(err.Error())
		handlers.RenderNotFoundError(w, err)
		return
	}

	handlers.RenderJSON(w, response)
}

func (a *App) readTaskHandler(ctx context.Context, taskID int) (*dto.GetTaskResponse, error) {
	task, err := a.db.SelectTask(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to select task by id: %w", err)
	}
	return &dto.GetTaskResponse{Task: *task}, nil
}

func (a *App) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rawTaskID := chi.URLParam(r, "task_id")
	taskID, err := strconv.Atoi(rawTaskID)
	if err != nil {
		a.log.Errorf("failed to get task ID: %v", err)
		handlers.RenderBadRequest(w, err)
		return
	}
	var request dto.UpdateTaskRequest
	if err := handlers.UnmarshalJSON(r, &request.Task); err != nil {
		a.log.Errorf("failed to unmarshal request json: %v", err)
		handlers.RenderBadRequest(w, err)
		return
	}

	if err := a.updateTaskHandler(ctx, taskID, request); err != nil {
		a.log.Errorf(err.Error())
		handlers.RenderNotFoundError(w, err)
		return
	}

	handlers.RenderOK(w)
}

func (a *App) updateTaskHandler(ctx context.Context, taskID int, request dto.UpdateTaskRequest) error {
	if err := a.db.UpdateTask(ctx, taskID, request.Task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (a *App) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rawTaskID := chi.URLParam(r, "task_id")
	taskID, err := strconv.Atoi(rawTaskID)
	if err != nil {
		a.log.Errorf("failed to get task ID: %v", err)
		handlers.RenderBadRequest(w, err)
		return
	}

	if err := a.deleteTaskHandler(ctx, taskID); err != nil {
		a.log.Errorf(err.Error())
		handlers.RenderNotFoundError(w, err)
		return
	}

	handlers.RenderOK(w)
}

func (a *App) deleteTaskHandler(ctx context.Context, taskID int) error {
	if err := a.db.DeleteTask(ctx, taskID); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (a *App) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request dto.GetTasksRequest
	if err := handlers.UnmarshalJSON(r, &request.Pagination); err != nil {
		a.log.Errorf("failed to unmarshal request json: %v", err)
		handlers.RenderBadRequest(w, err)
		return
	}

	response, err := a.getTasksHandler(ctx, request)
	if err != nil {
		a.log.Errorf(err.Error())
		handlers.RenderInternalError(w, err)
		return
	}

	handlers.RenderJSON(w, response)
}

func (a *App) getTasksHandler(ctx context.Context, request dto.GetTasksRequest) (*dto.GetTasksResponse, error) {
	tasks, err := a.db.SelectTasks(ctx, request.Pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to select tasks: %w", err)
	}
	return &dto.GetTasksResponse{Tasks: tasks}, nil
}
