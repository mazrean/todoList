package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
)

type DashboardsTable struct {
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
		"INSERT INTO dashboards (id, user_id, name, description, created_at) VALUES (?, ?, ?, ?, ?)",
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

	result, err := db.ExecContext(
		ctx,
		"UPDATE dashboards SET name = ?, description = ? WHERE id = ? && deleted_at IS NULL",
		string(dashboard.GetName()),
		string(dashboard.GetDescription()),
		uuid.UUID(dashboard.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update dashboard: %w", err)
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

func (d *Dashboard) DeleteDashboard(ctx context.Context, id values.DashboardID) error {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result, err := db.ExecContext(
		ctx,
		"UPDATE dashboards SET deleted_at = ? WHERE id = ?",
		uuid.UUID(id),
	)
	if err != nil {
		return fmt.Errorf("failed to delete dashboard: %w", err)
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

func (d *Dashboard) GetMyDashboards(ctx context.Context, userID values.UserID) ([]*domain.Dashboard, error) {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	dashboardTables := []DashboardsTable{}
	err = db.SelectContext(
		ctx,
		&dashboardTables,
		"SELECT id, name, description, created_at FROM dashboards WHERE user_id = ? AND deleted_at IS NULL",
		uuid.UUID(userID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard: %w", err)
	}

	dashboards := make([]*domain.Dashboard, 0, len(dashboardTables))
	for _, dashboardTable := range dashboardTables {
		dashboard := domain.NewDashboard(
			values.NewDashboardIDFromUUID(dashboardTable.ID),
			values.NewDashboardName(dashboardTable.Name),
			values.NewDashboardDescription(dashboardTable.Description),
			dashboardTable.CreatedAt,
		)
		dashboards = append(dashboards, dashboard)
	}

	return dashboards, nil
}

func (d *Dashboard) GetDashboard(ctx context.Context, id values.DashboardID, lockType repository.LockType) (*domain.Dashboard, error) {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	query := "SELECT id, name, description, created_at FROM dashboards WHERE id = ? AND deleted_at IS NULL"
	switch lockType {
	case repository.LockTypeRecord:
		query += " FOR UPDATE"
	}

	dashboardTable := DashboardsTable{}
	err = db.GetContext(
		ctx,
		&dashboardTable,
		query,
		uuid.UUID(id),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard: %w", err)
	}

	dashboard := domain.NewDashboard(
		values.NewDashboardIDFromUUID(dashboardTable.ID),
		values.NewDashboardName(dashboardTable.Name),
		values.NewDashboardDescription(dashboardTable.Description),
		dashboardTable.CreatedAt,
	)

	return dashboard, nil
}

func (d *Dashboard) GetDashboardOwner(ctx context.Context, id values.DashboardID) (*domain.User, error) {
	db, err := d.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	userTable := UsersTable{}
	err = db.GetContext(
		ctx,
		&userTable,
		"SELECT users.id, users.name, users.hashed_password FROM users JOIN dashboards ON users.id = dashboards.user_id WHERE dashboards.id = ? AND users.deleted_at IS NULL AND dashboards.deleted_at IS NULL",
		uuid.UUID(id),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user := domain.NewUser(
		values.NewUserIDFromUUID(userTable.ID),
		values.NewUserName(userTable.Name),
		values.NewUserHashedPassword(userTable.HashedPassword),
	)

	return user, nil
}
