package sqlx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type Dashboard struct {
	db *DB
}

func NewDashboard(db *DB) *Dashboard {
	return &Dashboard{
		db: db,
	}
}

func (d *Dashboard) CreateDashboard(ctx context.Context, userID values.UserID, dashboard *domain.Dashboard) error {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"INSERT INTO dashboard (id, user_id, name, description, created_at) VALUES (?, ?, ?, ?, ?)",
		uuid.UUID(dashboard.GetID()),
		uuid.UUID(userID),
		string(dashboard.GetName()),
		string(dashboard.GetDescription()),
		dashboard.GetCreatedAt(),
	)
	if err != nil {
		return fmt.Errorf("failed to create dashboard: %w", err)
	}

	return nil
}
