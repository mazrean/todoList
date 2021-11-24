package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mazrean/todoList/domain"
)

type Context struct{}

func NewContext() *Context {
	return &Context{}
}

const (
	taskStatusContextKey = "taskStatus"
)

func (ctx *Context) setTaskStatus(c *gin.Context, taskStatus *domain.TaskStatus) {
	c.Set(taskStatusContextKey, taskStatus)
}

func (ctx *Context) getTaskStatus(c *gin.Context) (*domain.TaskStatus, bool) {
	iTaskStatus, ok := c.Get(taskStatusContextKey)
	if !ok {
		return nil, false
	}

	taskStatus, ok := iTaskStatus.(*domain.TaskStatus)
	if !ok {
		return nil, false
	}

	return taskStatus, true
}
