package repository

//go:generate mockgen -source=$GOFILE -destination=mock/${GOFILE} -package=mock

import (
	"context"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id values.UserID) error
	GetUser(ctx context.Context, userID values.UserID, lockType LockType) (*domain.User, error)
	GetUserByName(ctx context.Context, name values.UserName) (*domain.User, error)
}
