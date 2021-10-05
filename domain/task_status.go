package domain

import "github.com/mazrean/todoList/domain/values"

type TaskStatus struct {
	id   values.TaskStatusID
	name values.TaskStatusName
}

func NewTaskStatus(id values.TaskStatusID, name values.TaskStatusName) *TaskStatus {
	return &TaskStatus{
		id: id,
		name: name,
	}
}
