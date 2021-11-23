package v1

import (
	"github.com/google/uuid"
)

type TaskStatusDetail struct {
	ID    uuid.UUID  `json:"id"`
	Name  string     `json:"name"`
	Tasks []TaskInfo `json:"tasks"`
}
