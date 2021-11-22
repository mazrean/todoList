package repository

import (
	"context"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type TaskStatus interface {
	CreateTaskStatus(ctx context.Context, taskStatus *domain.TaskStatus) error
	UpdateTaskStatus(ctx context.Context, taskStatus *domain.TaskStatus) error
	DeleteTaskStatus(ctx context.Context, taskStatus *domain.TaskStatus) error
	GetTaskStatus(ctx context.Context, taskStatusID values.TaskStatusID, lockType LockType) (*domain.TaskStatus, error)
	GetTaskStatusList(ctx context.Context, dashboardID values.DashboardID) ([]*domain.TaskStatus, error)
}
