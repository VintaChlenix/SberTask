package main

import (
	"SberTask/internal/app"
	"github.com/go-chi/chi"
)

func newRouter(a *app.App) chi.Router {
	r := chi.NewRouter()

	r.Post("/create_task", a.CreateTaskHandler)
	r.Get("/get_task/{task_id}", a.ReadTaskHandler)
	r.Post("/update_task/{task_id}", a.UpdateTaskHandler)
	r.Delete("/delete_task/{task_id}", a.DeleteTaskHandler)
	r.Get("/get_tasks", a.GetTasksHandler)

	return r
}
