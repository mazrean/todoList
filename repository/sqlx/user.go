package sqlx

type User struct {
	db *DB
}

func NewUser(db *DB) *User {
	return &User{
		db: db,
	}
}
