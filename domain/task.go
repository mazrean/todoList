package domain

import (
	"time"

	"github.com/mazrean/todoList/domain/values"
)

type Task struct {
	id          values.TaskID
	name        values.TaskName
	description values.TaskDescription
	createdAt   time.Time
}

func NewTask(
	id values.TaskID,
	name values.TaskName,
	description values.TaskDescription,
	createdAt time.Time,
) *Task {
	return &Task{
		id:          id,
		name:        name,
		description: description,
		createdAt:   createdAt,
	}
}
