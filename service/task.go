package service

import (
	"context"
	"errors"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

var (
	ErrNoTask                   = errors.New("no task")
	ErrCantMoveToOtherDashboard = errors.New("cant move to other dashboard")
)

type Task interface {
	CreateTask(ctx context.Context, taskStatusID values.TaskStatusID, name values.TaskName, description values.TaskDescription) (*domain.Task, error)
	UpdateTask(ctx context.Context, taskID values.TaskID, name values.TaskName, description values.TaskDescription) (*domain.Task, error)
	DeleteTask(ctx context.Context, taskID values.TaskID) error
	MovedTask(ctx context.Context, taskID values.TaskID, taskStatusID values.TaskStatusID) (*domain.Task, error)
}
