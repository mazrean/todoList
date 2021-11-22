package v1

import "github.com/mazrean/todoList/repository"

type Dashboard struct {
	db                  repository.DB
	dashboardRepository repository.Dashboard
}

func NewDashboard(db repository.DB, dashboardRepository repository.Dashboard) *Dashboard {
	return &Dashboard{
		db:                  db,
		dashboardRepository: dashboardRepository,
	}
}
