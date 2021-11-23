package v1

import "github.com/mazrean/todoList/repository"

type Task struct {
	db                   repository.DB
	dashboardRepository  repository.Dashboard
	taskStatusRepository repository.TaskStatus
	taskRepository       repository.Task
}

func NewTask(
	db repository.DB,
	dashboardRepository repository.Dashboard,
	taskStatusRepository repository.TaskStatus,
	taskRepository repository.Task,
) *Task {
	return &Task{
		db:                   db,
		dashboardRepository:  dashboardRepository,
		taskStatusRepository: taskStatusRepository,
		taskRepository:       taskRepository,
	}
}
