package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

	result, err := db.ExecContext(
		ctx,
		"UPDATE task_status SET name = ? WHERE id = ?",
		string(taskStatus.GetName()),
		uuid.UUID(taskStatus.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return repository.ErrNoRecordUpdated
	}

	return nil
}

func (ts *TaskStatus) DeleteTaskStatus(ctx context.Context, id values.TaskStatusID) error {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result, err := db.ExecContext(
		ctx,
		"UPDATE task_status SET deleted_at = ? WHERE id = ?",
		time.Now(),
		uuid.UUID(id),
	)
	if err != nil {
		return fmt.Errorf("failed to delete task status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return repository.ErrNoRecordDeleted
	}

	return nil
}

func (ts *TaskStatus) GetTaskStatus(ctx context.Context, taskStatusID values.TaskStatusID, lockType repository.LockType) (*domain.TaskStatus, error) {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	query := "SELECT id, name FROM task_status WHERE id = ?"
	switch lockType {
	case repository.LockTypeRecord:
		query += " FOR UPDATE"
	}

	taskStatusTable := TaskStatusTable{}
	err = db.GetContext(
		ctx,
		&taskStatusTable,
		query,
		uuid.UUID(taskStatusID),
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrRecordNotFound
	}
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

func (ts *TaskStatus) GetTaskStatusOwner(ctx context.Context, id values.TaskStatusID) (*domain.User, error) {
	db, err := ts.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	userTable := UsersTable{}
	err = db.GetContext(
		ctx,
		&userTable,
		"SELECT users.id, users.name, users.hashed_password FROM users "+
			"JOIN dashboards ON users.id = dashboards.user_id "+
			"JOIN task_status ON dashboards.id = task_status.dashboard_id "+
			"WHERE task_status.id = ? AND users.deleted_at IS NULL AND dashboards.deleted_at IS NULL AND task_status.deleted_at IS NULL",
		uuid.UUID(id),
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrRecordNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get task status owner: %w", err)
	}

	return domain.NewUser(
		values.NewUserIDFromUUID(userTable.ID),
		values.NewUserName(userTable.Name),
		values.NewUserHashedPassword(userTable.HashedPassword),
	), nil
}
