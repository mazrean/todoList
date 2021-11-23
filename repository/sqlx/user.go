package sqlx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
)

type User struct {
	db *DB
}

func NewUser(db *DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) CreateUser(ctx context.Context, user *domain.User) error {
	db, err := u.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.Exec(
		"INSERT INTO users (id, name, hashed_password) VALUES (?, ?, ?)",
		uuid.UUID(user.GetID()),
		string(user.GetName()),
		[]byte(user.GetHashedPassword()),
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
