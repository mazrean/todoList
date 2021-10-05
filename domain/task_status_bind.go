package domain

import (
	"time"

	"github.com/mazrean/todoList/domain/values"
)

type TaskStatusBind struct {
	id        values.TaskStatusBindID
	task      *Task
	status    *TaskStatus
	createdAt time.Time
}

func NewTaskStatusBind(id values.TaskStatusBindID, task *Task, status *TaskStatus, createdAt time.Time) *TaskStatusBind {
	return &TaskStatusBind{
		id:        id,
		task:      task,
		status:    status,
		createdAt: createdAt,
	}
}
