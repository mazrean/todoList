package v1

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/service"
)

type Middleware struct {
	context           *Context
	session           *Session
	dashboardService  service.Dashboard
	taskStatusService service.TaskStatus
	taskService       service.Task
}

func NewMiddleware(
	context *Context,
	session *Session,
	dashboardService service.Dashboard,
	taskStatusService service.TaskStatus,
	taskService service.Task,
) *Middleware {
	return &Middleware{
		context:           context,
		session:           session,
		dashboardService:  dashboardService,
		taskStatusService: taskStatusService,
		taskService:       taskService,
	}
}

func (m *Middleware) LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := m.session.getSession(c)
		_, err := m.session.getUser(session)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		m.context.setSession(c, session)

		c.Next()
	}
}

func (m *Middleware) DashboardUpdateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := m.session.getSession(c)
		user, err := m.session.getUser(session)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		strDashboardID := c.Param("dashboardID")
		dashboardID, err := uuid.Parse(strDashboardID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid dashboard id",
			})
			return
		}

		err = m.dashboardService.DashboardUpdateAuth(
			c.Request.Context(),
			user,
			values.NewDashboardIDFromUUID(dashboardID),
		)
		if errors.Is(err, service.ErrNoDashboard) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no dashboard",
			})
			return
		}
		if errors.Is(err, service.ErrNotDashboardOwner) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "not dashboard owner",
			})
			return
		}
		if err != nil {
			log.Printf("failed to update dashboard auth: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update dashboard auth",
			})
			return
		}

		m.context.setSession(c, session)

		c.Next()
	}
}

func (m *Middleware) TaskStatusUpdateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := m.session.getSession(c)
		user, err := m.session.getUser(session)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		strTaskStatusID := c.Param("taskStatusID")
		taskStatusID, err := uuid.Parse(strTaskStatusID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid task status id",
			})
			return
		}

		taskStatus, err := m.taskStatusService.TaskStatusUpdateAuth(
			c.Request.Context(),
			user,
			values.NewTaskStatusIDFromUUID(taskStatusID),
		)
		if errors.Is(err, service.ErrNoTaskStatus) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no task status",
			})
			return
		}
		if errors.Is(err, service.ErrNotDashboardOwner) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "not dashboard owner",
			})
			return
		}
		if err != nil {
			log.Printf("failed to update task status auth: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update dashboard auth",
			})
			return
		}

		m.context.setSession(c, session)
		m.context.setTaskStatus(c, taskStatus)

		c.Next()
	}
}

func (m *Middleware) TaskUpdateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := m.session.getSession(c)
		user, err := m.session.getUser(session)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		strTaskID := c.Param("taskID")
		taskID, err := uuid.Parse(strTaskID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid task id",
			})
			return
		}

		err = m.taskService.TaskUpdateAuth(
			c.Request.Context(),
			user,
			values.NewTaskIDFromUUID(taskID),
		)
		if errors.Is(err, service.ErrNoTask) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no task",
			})
			return
		}
		if errors.Is(err, service.ErrNotDashboardOwner) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "not dashboard owner",
			})
			return
		}
		if err != nil {
			log.Printf("failed to update task status auth: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update dashboard auth",
			})
			return
		}

		m.context.setSession(c, session)

		c.Next()
	}
}
