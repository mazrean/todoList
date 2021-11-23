package sqlx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
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

	_, err = db.ExecContext(
		ctx,
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

func (u *User) UpdateUser(ctx context.Context, user *domain.User) error {
	db, err := u.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"UPDATE users SET name = ?, hashed_password = ? WHERE id = ?",
		string(user.GetName()),
		[]byte(user.GetHashedPassword()),
		uuid.UUID(user.GetID()),
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, id values.UserID) error {
	db, err := u.db.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?",
		uuid.UUID(id),
	)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
