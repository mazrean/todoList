package v1

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/service"
)

type Task struct {
	context     *Context
	session     *Session
	taskService service.Task
}

func NewTask(
	context *Context,
	session *Session,
	taskService service.Task,
) *Task {
	return &Task{
		context:     context,
		session:     session,
		taskService: taskService,
	}
}

type TaskInfo struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (t *Task) PostTask(c *gin.Context) {
	var taskInfo TaskInfo
	err := c.BindJSON(&taskInfo)
	if err != nil {
		c.AbortWithStatus(400)
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

	task, err := t.taskService.CreateTask(
		c.Request.Context(),
		values.NewTaskStatusIDFromUUID(taskStatusID),
		values.NewTaskName(taskInfo.Name),
		values.NewTaskDescription(taskInfo.Description),
	)
	if errors.Is(err, service.ErrNoTaskStatus) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no task status",
		})
		return
	}
	if err != nil {
		log.Printf("failed to create task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, TaskInfo{
		ID:          uuid.UUID(task.GetID()),
		Name:        string(task.GetName()),
		Description: string(task.GetDescription()),
		CreatedAt:   task.GetCreatedAt(),
	})
}

func (t *Task) PatchTask(c *gin.Context) {
	var taskInfo TaskInfo
	err := c.BindJSON(&taskInfo)
	if err != nil {
		c.AbortWithStatus(400)
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

	task, err := t.taskService.UpdateTask(
		c.Request.Context(),
		values.NewTaskIDFromUUID(taskID),
		values.NewTaskName(taskInfo.Name),
		values.NewTaskDescription(taskInfo.Description),
	)
	if errors.Is(err, service.ErrNoTask) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no task",
		})
		return
	}
	if err != nil {
		log.Printf("failed to update task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, TaskInfo{
		ID:          uuid.UUID(task.GetID()),
		Name:        string(task.GetName()),
		Description: string(task.GetDescription()),
		CreatedAt:   task.GetCreatedAt(),
	})
}

func (t *Task) DeleteTask(c *gin.Context) {
	strTaskID := c.Param("taskID")
	taskID, err := uuid.Parse(strTaskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	err = t.taskService.DeleteTask(
		c.Request.Context(),
		values.NewTaskIDFromUUID(taskID),
	)
	if errors.Is(err, service.ErrNoTask) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no task",
		})
		return
	}
	if err != nil {
		log.Printf("failed to delete task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted",
	})
}

type MoveTaskRequest struct {
	DestTaskStatusID uuid.UUID `json:"dest" binding:"required"`
}

func (t *Task) PatchMoveTask(c *gin.Context) {
	strTaskID := c.Param("taskID")
	taskID, err := uuid.Parse(strTaskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	var moveTaskRequest MoveTaskRequest
	err = c.BindJSON(&moveTaskRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	_, err = t.taskService.MoveTask(
		c.Request.Context(),
		values.NewTaskIDFromUUID(taskID),
		values.NewTaskStatusIDFromUUID(moveTaskRequest.DestTaskStatusID),
	)
	if errors.Is(err, service.ErrNoTask) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no task",
		})
		return
	}
	if errors.Is(err, service.ErrNoTaskStatus) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no task status",
		})
		return
	}
	if errors.Is(err, service.ErrCantMoveToOtherDashboard) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't move to other dashboard",
		})
		return
	}
	if err != nil {
		log.Printf("failed to move task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to move task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task moved",
	})
}
