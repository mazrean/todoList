package sqlx

type TaskStatus struct {
	db *DB
}

func NewTaskStatus(db *DB) *TaskStatus {
	return &TaskStatus{
		db: db,
	}
}
