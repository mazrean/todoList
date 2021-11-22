package service

import (
	"context"
	"errors"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

var (
	ErrNoTaskStatus = errors.New("no task status")
)

type TaskStatus interface {
	AddTaskStatus(ctx context.Context, dashboardID values.DashboardID, name values.TaskStatusName) (*domain.TaskStatus, error)
	UpdateTaskStatus(ctx context.Context, id values.TaskStatusID, name values.TaskStatusName) (*domain.TaskStatus, error)
	DeleteTaskStatus(ctx context.Context, id values.TaskStatusID) error
	TaskStatusUpdateAuth(ctx context.Context, user *domain.User, id values.TaskStatusID) (*domain.TaskStatus, error)
}
