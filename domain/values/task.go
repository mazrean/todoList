package values

import "github.com/google/uuid"

type (
	TaskID          uuid.UUID
	TaskName        string
	TaskDescription string

	TaskStatusID   uuid.UUID
	TaskStatusName string
)

func NewTaskID() TaskID {
	return TaskID(uuid.New())
}

func NewTaskIDFromUUID(id uuid.UUID) TaskID {
	return TaskID(id)
}

func NewTaskName(name string) TaskName {
	return TaskName(name)
}

func NewTaskDescription(description string) TaskDescription {
	return TaskDescription(description)
}

func NewTaskStatusID() TaskStatusID {
	return TaskStatusID(uuid.New())
}

func NewTaskStatusIDFromUUID(id uuid.UUID) TaskStatusID {
	return TaskStatusID(id)
}

func NewTaskStatusName(name string) TaskStatusName {
	return TaskStatusName(name)
}
