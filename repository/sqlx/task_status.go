package sqlx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
)

type TaskStatusTable struct {
	ID          uuid.UUID `db:"id"`
	DashboardID uuid.UUID `db:"dashboard_id"`
	Name        string    `db:"name"`
}

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

func (ts *TaskStatus) DeleteTaskStatus(ctx context.Context, id values.TaskStatusID) error {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"DELETE FROM task_status WHERE id = ?",
		uuid.UUID(id),
	)
	if err != nil {
		return fmt.Errorf("failed to delete task status: %w", err)
	}

	return nil
}

func (ts *TaskStatus) GetTaskStatus(ctx context.Context, taskStatusID values.TaskStatusID, lockType repository.LockType) (*domain.TaskStatus, error) {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	taskStatusTable := TaskStatusTable{}
	err = db.GetContext(
		ctx,
		&taskStatusTable,
		"SELECT id, name FROM task_status WHERE id = ?",
		uuid.UUID(taskStatusID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get task status: %w", err)
	}

	taskStatus := domain.NewTaskStatus(
		values.NewTaskStatusIDFromUUID(taskStatusTable.ID),
		values.NewTaskStatusName(taskStatusTable.Name),
	)

	return taskStatus, nil
}

func (ts *TaskStatus) GetTaskStatusList(ctx context.Context, dashboardID values.DashboardID) ([]*domain.TaskStatus, error) {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	taskStatusTableList := []TaskStatusTable{}
	err = db.SelectContext(
		ctx,
		&taskStatusTableList,
		"SELECT id, name FROM task_status WHERE dashboard_id = ?",
		uuid.UUID(dashboardID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get task status list: %w", err)
	}

	taskStatusList := make([]*domain.TaskStatus, 0, len(taskStatusTableList))
	for _, taskStatusTable := range taskStatusTableList {
		taskStatusList = append(taskStatusList, domain.NewTaskStatus(
			values.NewTaskStatusIDFromUUID(taskStatusTable.ID),
			values.NewTaskStatusName(taskStatusTable.Name),
		))
	}

	return taskStatusList, nil
}
