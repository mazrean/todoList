package sqlx

type Dashboard struct {
	db *DB
}

func NewDashboard(db *DB) *Dashboard {
	return &Dashboard{
		db: db,
	}
}
