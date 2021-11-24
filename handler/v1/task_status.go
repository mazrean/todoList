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

type TaskStatus struct {
	context           *Context
	session           *Session
	taskStatusService service.TaskStatus
}

func NewTaskStatus(
	context *Context,
	session *Session,
	taskStatusService service.TaskStatus,
) *TaskStatus {
	return &TaskStatus{
		context:           context,
		session:           session,
		taskStatusService: taskStatusService,
	}
}

type TaskStatusInfo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" binding:"required"`
}

type TaskStatusDetail struct {
	ID    uuid.UUID  `json:"id"`
	Name  string     `json:"name"`
	Tasks []TaskInfo `json:"tasks"`
}

func (ts *TaskStatus) PostTaskStatus(c *gin.Context) {
	var taskStatusInfo TaskStatusInfo
	err := c.BindJSON(&taskStatusInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
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

	taskStatus, err := ts.taskStatusService.AddTaskStatus(
		c.Request.Context(),
		values.NewDashboardIDFromUUID(dashboardID),
		values.NewTaskStatusName(taskStatusInfo.Name),
	)
	if errors.Is(err, service.ErrNoDashboard) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no dashboard",
		})
		return
	}
	if err != nil {
		log.Printf("failed to add task status: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to add task status",
		})
		return
	}

	c.JSON(http.StatusCreated, TaskStatusInfo{
		ID:   uuid.UUID(taskStatus.GetID()),
		Name: string(taskStatus.GetName()),
	})
}

func (ts *TaskStatus) PatchTaskStatus(c *gin.Context) {
	var taskStatusInfo TaskStatusInfo
	err := c.BindJSON(&taskStatusInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
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

	taskStatus, err := ts.taskStatusService.UpdateTaskStatus(
		c.Request.Context(),
		values.NewTaskStatusIDFromUUID(taskStatusID),
		values.NewTaskStatusName(taskStatusInfo.Name),
	)
	if errors.Is(err, service.ErrNoTaskStatus) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no task status",
		})
		return
	}
	if err != nil {
		log.Printf("failed to update task status: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update task status",
		})
		return
	}

	c.JSON(http.StatusOK, TaskStatusInfo{
		ID:   uuid.UUID(taskStatus.GetID()),
		Name: string(taskStatus.GetName()),
	})
}

func (ts *TaskStatus) DeleteTaskStatus(c *gin.Context) {
	strTaskStatusID := c.Param("taskStatusID")
	taskStatusID, err := uuid.Parse(strTaskStatusID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task status id",
		})
		return
	}

	err = ts.taskStatusService.DeleteTaskStatus(
		c.Request.Context(),
		values.NewTaskStatusIDFromUUID(taskStatusID),
	)
	if errors.Is(err, service.ErrNoTaskStatus) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no task status",
		})
		return
	}
	if err != nil {
		log.Printf("failed to delete task status: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete task status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
