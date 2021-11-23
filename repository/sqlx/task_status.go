package sqlx

import (
	"context"

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
}
