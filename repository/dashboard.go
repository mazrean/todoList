package repository

//go:generate mockgen -source=$GOFILE -destination=mock/${GOFILE} -package=mock

import (
	"context"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type Dashboard interface {
	CreateDashboard(ctx context.Context, dashboard *domain.Dashboard) error
	UpdateDashboard(ctx context.Context, dashboard *domain.Dashboard) error
	DeleteDashboard(ctx context.Context, id values.DashboardID) error
	GetMyDashboards(ctx context.Context, userID values.UserID) ([]*domain.Dashboard, error)
	GetDashboard(ctx context.Context, id values.DashboardID, lockType LockType) (*domain.Dashboard, error)
	GetDashboardOwner(ctx context.Context, id values.DashboardID) (*domain.User, error)
}
