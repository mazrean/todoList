package sqlx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type Task struct {
	db *DB
}

func NewTask(db *DB) *Task {
	return &Task{
		db: db,
	}
}

func (t *Task) CreateTask(ctx context.Context, taskStatusID values.TaskStatusID, task *domain.Task) error {
	db, err := t.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		`INSERT INTO tasks (id, task_status_id, name, description, created_at) VALUES (?, ?, ?, ?, ?)`,
		uuid.UUID(task.GetID()),
		uuid.UUID(taskStatusID),
		string(task.GetName()),
		string(task.GetDescription()),
		task.GetCreatedAt(),
	)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (t *Task) UpdateTask(ctx context.Context, task *domain.Task) error {
	db, err := t.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		`UPDATE tasks SET name = ?, description = ? WHERE id = ?`,
		string(task.GetName()),
		string(task.GetDescription()),
		uuid.UUID(task.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}
