package domain

import (
	"time"

	"github.com/mazrean/todoList/domain/values"
)

type Dashboard struct {
	id         values.DashboardID
	name       values.DashboardName
	description values.DashboardDescription
	createdAt  time.Time
}

func NewDashboard(
	id 				values.DashboardID,
	name values.DashboardName,
	description values.DashboardDescription,
	createdAt time.Time,
) *Dashboard {
	return &Dashboard{
		id:         id,
		name:       name,
		description: description,
		createdAt:  createdAt,
	}
}

func (d *Dashboard) GetID() values.DashboardID {
	return d.id
}

func (d *Dashboard) GetName() values.DashboardName {
	return d.name
}

func (d *Dashboard) GetDescription() values.DashboardDescription {
	return d.description
}

func (d *Dashboard) GetCreatedAt() time.Time {
	return d.createdAt
}
