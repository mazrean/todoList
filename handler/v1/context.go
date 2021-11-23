package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mazrean/todoList/domain"
)

type Context struct{}

const (
	sessionContextKey    = "session"
	taskStatusContextKey = "taskStatus"
)

func (ctx *Context) setSession(c *gin.Context, session sessions.Session) {
	c.Set(sessionContextKey, session)
}

func (ctx *Context) getSession(c *gin.Context) (sessions.Session, bool) {
	iSession, ok := c.Get(sessionContextKey)
	if !ok {
		return nil, false
	}

	session, ok := iSession.(sessions.Session)
	if !ok {
		return nil, false
	}

	return session, true
}

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
