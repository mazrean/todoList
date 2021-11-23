package sqlx

type Task struct {
	db *DB
}

func NewTask(db *DB) *Task {
	return &Task{
		db: db,
	}
}
