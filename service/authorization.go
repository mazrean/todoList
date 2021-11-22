package service

import (
	"context"
	"errors"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserOrPassword = errors.New("invalid user or password")
)

type Authorization interface {
	Signup(ctx context.Context, name values.UserName, password values.UserPassword) (*domain.User, error)
	Login(ctx context.Context, name values.UserName, password values.UserPassword) (*domain.User, error)
}
