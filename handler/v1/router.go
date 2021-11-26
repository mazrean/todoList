package v1

import (
	"github.com/gin-contrib/static"
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

	r.Use(static.Serve("/", static.LocalFile("/static", false)))

	api := r.Group("/api")
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
			dashboard.POST("", a.middleware.LoginAuth(), a.dashboardHandler.PostDashboard)
			dashboard.PATCH("/:dashboardID", a.middleware.DashboardUpdateAuth(), a.dashboardHandler.PatchDashboard)
			dashboard.DELETE("/:dashboardID", a.middleware.DashboardUpdateAuth(), a.dashboardHandler.DeleteDashboard)
			dashboard.GET("/:dashboardID", a.middleware.DashboardUpdateAuth(), a.dashboardHandler.GetDashboardInfo)
			dashboard.POST("/:dashboardID/status", a.middleware.DashboardUpdateAuth(), a.taskStatusHandler.PostTaskStatus)
		}

		task := api.Group("/tasks")
		{
			task.PATCH("/:taskID", a.middleware.TaskUpdateAuth(), a.taskHandler.PatchTask)
			task.DELETE("/:taskID", a.middleware.TaskUpdateAuth(), a.taskHandler.DeleteTask)
			task.PATCH("/:taskID/move", a.middleware.TaskUpdateAuth(), a.taskHandler.PatchMoveTask)

			status := task.Group("/status")
			{
				status.PATCH("/:taskStatusID", a.middleware.TaskStatusUpdateAuth(), a.taskStatusHandler.PatchTaskStatus)
				status.DELETE("/:taskStatusID", a.middleware.TaskStatusUpdateAuth(), a.taskStatusHandler.DeleteTaskStatus, a.middleware.TaskStatusUpdateAuth())
				status.POST("/:taskStatusID/tasks", a.middleware.TaskStatusUpdateAuth(), a.taskHandler.PostTask)
			}
		}
	}

	r.Run(a.addr)
}
