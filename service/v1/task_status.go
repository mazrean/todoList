package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
	"github.com/mazrean/todoList/service"
)

type TaskStatus struct {
	db                   repository.DB
	dashboardRepository  repository.Dashboard
	taskStatusRepository repository.TaskStatus
}

func NewTaskStatus(db repository.DB, dashboardRepository repository.Dashboard, taskStatusRepository repository.TaskStatus) TaskStatus {
	return TaskStatus{
		db:                   db,
		dashboardRepository:  dashboardRepository,
		taskStatusRepository: taskStatusRepository,
	}
}

func (ts *TaskStatus) AddTaskStatus(ctx context.Context, dashboardID values.DashboardID, name values.TaskStatusName) (*domain.TaskStatus, error) {
	var taskStatus *domain.TaskStatus
	err := ts.db.Transaction(ctx, nil, func(ctx context.Context) error {
		_, err := ts.dashboardRepository.GetDashboard(ctx, dashboardID, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return service.ErrNoDashboard
		}
		if err != nil {
			return fmt.Errorf("failed to get dashboard: %w", err)
		}

		taskStatus = domain.NewTaskStatus(
			values.NewTaskStatusID(),
			name,
		)

		err = ts.taskStatusRepository.CreateTaskStatus(ctx, dashboardID, taskStatus)
		if err != nil {
			return fmt.Errorf("failed to create task status: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return taskStatus, err
}

func (ts *TaskStatus) UpdateTaskStatus(ctx context.Context, id values.TaskStatusID, name values.TaskStatusName) (*domain.TaskStatus, error) {
	var taskStatus *domain.TaskStatus
	err := ts.db.Transaction(ctx, nil, func(ctx context.Context) error {
		_, err := ts.taskStatusRepository.GetTaskStatus(ctx, id, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return service.ErrNoTaskStatus
		}
		if err != nil {
			return fmt.Errorf("failed to get task status: %w", err)
		}

		taskStatus = domain.NewTaskStatus(
			id,
			name,
		)

		err = ts.taskStatusRepository.UpdateTaskStatus(ctx, taskStatus)
		if err != nil {
			return fmt.Errorf("failed to update task status: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return taskStatus, err
}

func (ts *TaskStatus) DeleteTaskStatus(ctx context.Context, id values.TaskStatusID) error {
	err := ts.db.Transaction(ctx, nil, func(ctx context.Context) error {
		_, err := ts.taskStatusRepository.GetTaskStatus(ctx, id, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return service.ErrNoTaskStatus
		}
		if err != nil {
			return fmt.Errorf("failed to get task status: %w", err)
		}

		err = ts.taskStatusRepository.DeleteTaskStatus(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to delete task status: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed in transaction: %w", err)
	}

	return err
}
