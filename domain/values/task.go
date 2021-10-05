package values

import "github.com/google/uuid"

type (
	TaskID          uuid.UUID
	TaskName        string
	TaskDescription string

	TaskStatusID   uuid.UUID
	TaskStatusName string

	TaskStatusBindID uuid.UUID
)
