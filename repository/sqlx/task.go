package sqlx

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
)

type TaskTable struct {
	ID           uuid.UUID `db:"id"`
	TaskStatusID uuid.UUID `db:"task_status_id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
}

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
		"INSERT INTO tasks (id, task_status_id, name, description, created_at) VALUES (?, ?, ?, ?, ?)",
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
		"UPDATE tasks SET name = ?, description = ? WHERE id = ?",
		string(task.GetName()),
		string(task.GetDescription()),
		uuid.UUID(task.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (t *Task) UpdateTaskStatus(ctx context.Context, taskID values.TaskID, taskStatusID values.TaskStatusID) error {
	db, err := t.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"UPDATE tasks SET task_status_id = ? WHERE id = ?",
		uuid.UUID(taskStatusID),
		uuid.UUID(taskID),
	)
	if err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}

	return nil
}

func (t *Task) DeleteTask(ctx context.Context, taskID values.TaskID) error {
	db, err := t.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"DELETE FROM tasks WHERE id = ?",
		uuid.UUID(taskID),
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (t *Task) GetTask(ctx context.Context, taskID values.TaskID, lockType repository.LockType) (*domain.Task, error) {
	db, err := t.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	query := "SELECT id, name, description, created_at FROM tasks WHERE id = ?"
	switch lockType {
	case repository.LockTypeRecord:
		query += " FOR UPDATE"
	}

	taskTable := TaskTable{}
	err = db.GetContext(
		ctx,
		&taskTable,
		query,
		uuid.UUID(taskID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return domain.NewTask(
		values.NewTaskIDFromUUID(taskTable.ID),
		values.NewTaskName(taskTable.Name),
		values.NewTaskDescription(taskTable.Description),
		taskTable.CreatedAt,
	), nil
}

func (t *Task) GetTasks(ctx context.Context, dashboardID values.DashboardID) (map[values.TaskStatusID][]*domain.Task, error) {
	db, err := t.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	tasks := []TaskTable{}
	err = db.SelectContext(
		ctx,
		&tasks,
		"SELECT id, task_status_id, name, description, created_at FROM tasks "+
			"JOIN task_status ON task_status.id = tasks.task_status_id "+
			"WHERE task_status.dashboard_id = ?",
		uuid.UUID(dashboardID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	taskMap := make(map[values.TaskStatusID][]*domain.Task)
	for _, taskTable := range tasks {
		task := domain.NewTask(
			values.NewTaskIDFromUUID(taskTable.ID),
			values.NewTaskName(taskTable.Name),
			values.NewTaskDescription(taskTable.Description),
			taskTable.CreatedAt,
		)
		taskMap[values.NewTaskStatusIDFromUUID(taskTable.TaskStatusID)] = append(taskMap[values.NewTaskStatusIDFromUUID(taskTable.TaskStatusID)], task)
	}

	return taskMap, nil
}
