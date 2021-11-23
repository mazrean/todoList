package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
)

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

func (t *Task) CreateTask(ctx context.Context, taskStatusID values.TaskStatusID, name values.TaskName, description values.TaskDescription) (*domain.Task, error) {
	var task *domain.Task
	err := t.db.Transaction(ctx, nil, func(ctx context.Context) error {
		_, err := t.taskStatusRepository.GetTaskStatus(ctx, taskStatusID, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return errors.New("task status not found")
		}
		if err != nil {
			return fmt.Errorf("failed to get task status: %w", err)
		}

		task = domain.NewTask(
			values.NewTaskID(),
			name,
			description,
			time.Now(),
		)

		err = t.taskRepository.CreateTask(ctx, taskStatusID, task)
		if err != nil {
			return fmt.Errorf("failed to create task: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return task, nil
}

func (t *Task) UpdateTask(ctx context.Context, taskID values.TaskID, name values.TaskName, description values.TaskDescription) (*domain.Task, error) {
	var task *domain.Task
	err := t.db.Transaction(ctx, nil, func(ctx context.Context) error {
		task, err := t.taskRepository.GetTask(ctx, taskID, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		if err != nil {
			return fmt.Errorf("failed to get task: %w", err)
		}

		task.SetName(name)
		task.SetDescription(description)

		err = t.taskRepository.UpdateTask(ctx, task)
		if err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return task, nil
}

func (t *Task) DeleteTask(ctx context.Context, taskID values.TaskID) error {
	err := t.db.Transaction(ctx, nil, func(ctx context.Context) error {
		_, err := t.taskRepository.GetTask(ctx, taskID, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		if err != nil {
			return fmt.Errorf("failed to get task: %w", err)
		}

		err = t.taskRepository.DeleteTask(ctx, taskID)
		if err != nil {
			return fmt.Errorf("failed to delete task: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed in transaction: %w", err)
	}

	return nil
}
