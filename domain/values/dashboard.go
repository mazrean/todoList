package values

import "github.com/google/uuid"

type (
	DashboardID uuid.UUID
	DashboardName string
	DashboardDescription string
)

func NewDashboardID() DashboardID {
	return DashboardID(uuid.New())
}

func NewDashboadIDFromUUID(id uuid.UUID) DashboardID {
	return DashboardID(id)
}

func NewDashboardName(name string) DashboardName {
	return DashboardName(name)
}

func NewDashboardDescription(description string) DashboardDescription {
	return DashboardDescription(description)
}
