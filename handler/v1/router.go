package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mazrean/todoList/common"
)

type API struct {
	addr              string
	session           *Session
	middleware        *Middleware
	userHandler       *User
	dashboardHandler  *Dashboard
	taskStatusHandler *TaskStatus
	taskHandler       *Task
}

func NewAPI(
	addr common.Addr,
	session *Session,
	middleware *Middleware,
	userHandler *User,
	dashboardHandler *Dashboard,
	taskStatusHandler *TaskStatus,
	taskHandler *Task,
) *API {
	return &API{
		addr:              string(addr),
		session:           session,
		middleware:        middleware,
		userHandler:       userHandler,
		dashboardHandler:  dashboardHandler,
		taskStatusHandler: taskStatusHandler,
		taskHandler:       taskHandler,
	}
}

func (a *API) Start() {
	r := gin.Default()

	a.session.Use(r)

	api := r.Group("/api/v1")
	{
		api.POST("/signup", a.userHandler.PostSignup)
		api.POST("/login", a.userHandler.PostLogin)

		me := api.Group("/users/me", a.middleware.LoginAuth())
		{
			me.GET("", a.userHandler.GetMe)
			me.PATCH("", a.userHandler.PatchMe)
			me.DELETE("", a.userHandler.DeleteMe)
			me.GET("/dashboards", a.dashboardHandler.GetMyDashboards)
		}

		dashboard := api.Group("/dashboards")
		{
			dashboard.POST("", a.dashboardHandler.PostDashboard, a.middleware.LoginAuth())
			dashboard.PATCH("/:dashboardID", a.dashboardHandler.PatchDashboard, a.middleware.DashboardUpdateAuth())
			dashboard.DELETE("/:dashboardID", a.dashboardHandler.DeleteDashboard, a.middleware.DashboardUpdateAuth())
			dashboard.GET("/:dashboardID", a.dashboardHandler.GetDashboardInfo, a.middleware.DashboardUpdateAuth())
			dashboard.POST("/:dashboardID/status", a.taskStatusHandler.PostTaskStatus, a.middleware.DashboardUpdateAuth())
		}

		task := api.Group("/tasks")
		{
			task.PATCH("/:taskID", a.taskHandler.PatchTask, a.middleware.TaskUpdateAuth())
			task.DELETE("/:taskID", a.taskHandler.DeleteTask, a.middleware.TaskUpdateAuth())
			task.PATCH("/:taskID/move", a.taskHandler.PatchMoveTask, a.middleware.TaskUpdateAuth())

			status := task.Group("/status")
			{
				status.PATCH("/:taskStatusID", a.taskStatusHandler.PatchTaskStatus, a.middleware.TaskStatusUpdateAuth())
				status.DELETE("/:taskStatusID", a.taskStatusHandler.DeleteTaskStatus, a.middleware.TaskStatusUpdateAuth())
				status.POST("/:taskStatusID/tasks", a.taskHandler.PostTask, a.middleware.TaskStatusUpdateAuth())
			}
		}
	}

	r.Run(a.addr)
}
