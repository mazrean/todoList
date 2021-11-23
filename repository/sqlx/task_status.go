package sqlx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type TaskStatus struct {
	db *DB
}

func NewTaskStatus(db *DB) *TaskStatus {
	return &TaskStatus{
		db: db,
	}
}

func (ts *TaskStatus) CreateTaskStatus(ctx context.Context, dashboardID values.DashboardID, taskStatus *domain.TaskStatus) error {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"INSERT INTO task_status (id, dashboard_id, name) VALUES (?, ?, ?)",
		uuid.UUID(taskStatus.GetID()),
		uuid.UUID(dashboardID),
		string(taskStatus.GetName()),
	)
	if err != nil {
		return fmt.Errorf("failed to insert task status: %w", err)
	}

	return nil
}

func (ts *TaskStatus) UpdateTaskStatus(ctx context.Context, taskStatus *domain.TaskStatus) error {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"UPDATE task_status SET name = ? WHERE id = ?",
		string(taskStatus.GetName()),
		uuid.UUID(taskStatus.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}

	return nil
}
