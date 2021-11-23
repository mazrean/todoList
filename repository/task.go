package repository

//go:generate mockgen -source=$GOFILE -destination=mock/${GOFILE} -package=mock

import (
	"context"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type Task interface {
	CreateTask(ctx context.Context, taskStatusID values.TaskStatusID, task *domain.Task) error
	UpdateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, taskID values.TaskID) error
	GetTask(ctx context.Context, taskID values.TaskID, lockType LockType) (*domain.Task, error)
	GetTasks(ctx context.Context, dashboardID values.DashboardID) (map[values.TaskStatusID][]*domain.Task, error)
	GetTaskOwner(ctx context.Context, id values.TaskID) (*domain.User, error)
}
