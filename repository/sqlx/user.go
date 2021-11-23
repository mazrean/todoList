package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
)

type UsersTable struct {
	ID             uuid.UUID    `db:"id"`
	Name           string       `db:"name"`
	HashedPassword []byte       `db:"hashed_password"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}

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
		"UPDATE users SET name = ?, hashed_password = ? WHERE id = ? AND deleted_at IS NULL",
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
		"UPDATE users SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL",
		time.Now(),
		uuid.UUID(id),
	)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (u *User) GetUser(ctx context.Context, userID values.UserID, lockType repository.LockType) (*domain.User, error) {
	db, err := u.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var userTable UsersTable
	err = db.GetContext(
		ctx,
		&userTable,
		"SELECT id, name, hashed_password FROM users WHERE id = ? AND deleted_at IS NULL",
		uuid.UUID(userID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user := domain.NewUser(
		values.NewUserIDFromUUID(userTable.ID),
		values.NewUserName(userTable.Name),
		values.NewUserHashedPassword(userTable.HashedPassword),
	)

	return user, nil
}

func (u *User) GetUserByName(ctx context.Context, name values.UserName) (*domain.User, error) {
	db, err := u.db.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var userTable UsersTable
	err = db.GetContext(
		ctx,
		&userTable,
		"SELECT id, name, hashed_password FROM users WHERE name = ? AND deleted_at IS NULL",
		string(name),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user := domain.NewUser(
		values.NewUserIDFromUUID(userTable.ID),
		values.NewUserName(userTable.Name),
		values.NewUserHashedPassword(userTable.HashedPassword),
	)

	return user, nil
}
