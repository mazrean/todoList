package service

import (
	"context"
	"errors"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

var (
	ErrNoDashboard       = errors.New("no dashboard")
	ErrNotDashboardOwner = errors.New("not dashboard owner")
)

type Dashboard interface {
	CreateDashboard(ctx context.Context, name values.DashboardName, description values.DashboardDescription) (*domain.Dashboard, error)
	UpdateDashboard(ctx context.Context, id values.DashboardID, name values.DashboardName, description values.DashboardDescription) (*domain.Dashboard, error)
	DeleteDashboard(ctx context.Context, id values.DashboardID) error
	GetMyDashboards(ctx context.Context, user *domain.User) ([]*domain.Dashboard, error)
	GetDashboardInfo(ctx context.Context, id values.DashboardID) (*DashboardInfo, error)
	DashboardUpdateAuth(ctx context.Context, user *domain.User, id values.DashboardID) error
}

type DashboardInfo struct {
	*domain.Dashboard
	TaskStatus []*TaskStatusInfo
}

type TaskStatusInfo struct {
	*domain.TaskStatus
	Tasks []*domain.Task
}
