package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type DashboardTable struct {
	ID          uuid.UUID    `db:"id"`
	UserID      uuid.UUID    `db:"user_id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	CreatedAt   time.Time    `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

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

func (d *Dashboard) UpdateDashboard(ctx context.Context, dashboard *domain.Dashboard) error {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"UPDATE dashboard SET name = ?, description = ? WHERE id = ? && deleted_at IS NULL",
		string(dashboard.GetName()),
		string(dashboard.GetDescription()),
		uuid.UUID(dashboard.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update dashboard: %w", err)
	}

	return nil
}

func (d *Dashboard) DeleteDashboard(ctx context.Context, id values.DashboardID) error {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"UPDATE dashboard SET deleted_at = ? WHERE id = ?",
		uuid.UUID(id),
	)
	if err != nil {
		return fmt.Errorf("failed to delete dashboard: %w", err)
	}

	return nil
}

func (d *Dashboard) GetMyDashboards(ctx context.Context, userID values.UserID) ([]*domain.Dashboard, error) {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	dashboardTables := []DashboardTable{}
	err = db.SelectContext(
		ctx,
		&dashboardTables,
		"SELECT id, user_id, name, description, created_at FROM dashboard WHERE user_id = ? AND deleted_at IS NULL",
		uuid.UUID(userID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard: %w", err)
	}

	dashboards := make([]*domain.Dashboard, 0, len(dashboardTables))
	for _, dashboardTable := range dashboardTables {
		dashboard := domain.NewDashboard(
			values.DashboardID(dashboardTable.ID),
			values.DashboardName(dashboardTable.Name),
			values.DashboardDescription(dashboardTable.Description),
			dashboardTable.CreatedAt,
		)
		dashboards = append(dashboards, dashboard)
	}

	return dashboards, nil
}
