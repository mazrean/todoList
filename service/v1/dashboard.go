package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
	"github.com/mazrean/todoList/service"
)

type Dashboard struct {
	db                   repository.DB
	dashboardRepository  repository.Dashboard
	taskStatusRepository repository.TaskStatus
}

func NewDashboard(
	db repository.DB,
	dashboardRepository repository.Dashboard,
	taskStatusRepository repository.TaskStatus,
) *Dashboard {
	return &Dashboard{
		db:                   db,
		dashboardRepository:  dashboardRepository,
		taskStatusRepository: taskStatusRepository,
	}
}

var (
	defaultTaskStatus = []values.TaskStatusName{
		values.TaskStatusName("todo"),
		values.TaskStatusName("in progress"),
		values.TaskStatusName("done"),
	}
)

func (d *Dashboard) CreateDashboard(ctx context.Context, user *domain.User, name values.DashboardName, description values.DashboardDescription) (*domain.Dashboard, error) {
	var dashboard *domain.Dashboard
	err := d.db.Transaction(ctx, nil, func(context.Context) error {
		dashboard = domain.NewDashboard(
			values.NewDashboardID(),
			name,
			description,
			time.Now(),
		)

		err := d.dashboardRepository.CreateDashboard(ctx, user.GetID(), dashboard)
		if err != nil {
			return fmt.Errorf("failed to create dashboard: %w", err)
		}

		// TODO: bulk insertした方がパフォーマンスの観点で良い
		for _, statusName := range defaultTaskStatus {
			taskStatus := domain.NewTaskStatus(
				values.NewTaskStatusID(),
				statusName,
			)

			err := d.taskStatusRepository.CreateTaskStatus(ctx, dashboard.GetID(), taskStatus)
			if err != nil {
				return fmt.Errorf("failed to create task status: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return dashboard, nil
}

func (d *Dashboard) UpdateDashboard(ctx context.Context, id values.DashboardID, name values.DashboardName, description values.DashboardDescription) (*domain.Dashboard, error) {
	var dashboard *domain.Dashboard
	err := d.db.Transaction(ctx, nil, func(context.Context) error {
		var err error
		dashboard, err = d.dashboardRepository.GetDashboard(ctx, id, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return service.ErrNoDashboard
		}
		if err != nil {
			return fmt.Errorf("failed to get dashboard: %w", err)
		}

		dashboard.SetName(name)
		dashboard.SetDescription(description)

		err = d.dashboardRepository.UpdateDashboard(ctx, dashboard)
		if err != nil {
			return fmt.Errorf("failed to update dashboard: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return dashboard, nil
}

func (d *Dashboard) DeleteDashboard(ctx context.Context, id values.DashboardID) error {
	err := d.db.Transaction(ctx, nil, func(context.Context) error {
		_, err := d.dashboardRepository.GetDashboard(ctx, id, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return service.ErrNoDashboard
		}
		if err != nil {
			return fmt.Errorf("failed to get dashboard: %w", err)
		}

		err = d.dashboardRepository.DeleteDashboard(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to delete dashboard: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed in transaction: %w", err)
	}

	return nil
}

func (d *Dashboard) GetMyDashboards(ctx context.Context, user *domain.User) ([]*domain.Dashboard, error) {
	dashboards, err := d.dashboardRepository.GetMyDashboards(ctx, user.GetID())
	if err != nil {
		return nil, fmt.Errorf("failed to get my dashboards: %w", err)
	}

	return dashboards, nil
}

func (d *Dashboard) DashboardUpdateAuth(ctx context.Context, user *domain.User, id values.DashboardID) error {
	err := d.db.Transaction(ctx, nil, func(context.Context) error {
		_, err := d.dashboardRepository.GetDashboard(ctx, id, repository.LockTypeRecord)
		if errors.Is(err, repository.ErrRecordNotFound) {
			return service.ErrNoDashboard
		}
		if err != nil {
			return fmt.Errorf("failed to get dashboard: %w", err)
		}

		owner, err := d.dashboardRepository.GetDashboardOwner(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to update dashboard auth: %w", err)
		}

		if owner.GetID() != user.GetID() {
			return service.ErrNotDashboardOwner
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed in transaction: %w", err)
	}

	return nil
}
