package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
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
