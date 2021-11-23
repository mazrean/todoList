package v1

import "github.com/mazrean/todoList/service"

type Dashboard struct {
	context          *Context
	session          *Session
	dashboardService service.Dashboard
}

func NewDashboard(
	context *Context,
	session *Session,
	dashboardService service.Dashboard,
) *Dashboard {
	return &Dashboard{
		context:          context,
		session:          session,
		dashboardService: dashboardService,
	}
}
