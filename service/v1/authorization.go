package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/repository"
	"github.com/mazrean/todoList/service"
)

type Authorization struct {
	db             repository.DB
	userRepository repository.User
}

func NewAuthorization(db repository.DB, userRepository repository.User) *Authorization {
	return &Authorization{
		db:             db,
		userRepository: userRepository,
	}
}

func (a *Authorization) Signup(ctx context.Context, name values.UserName, password values.UserPassword) (*domain.User, error) {
	var user *domain.User
	err := a.db.Transaction(ctx, nil, func(ctx context.Context) error {
		_, err := a.userRepository.GetUserByName(ctx, name)
		if err == nil {
			return service.ErrUserAlreadyExists
		}
		if err != nil && !errors.Is(err, repository.ErrRecordNotFound) {
			return fmt.Errorf("failed to get user: %w", err)
		}

		hashedPassword, err := password.Hash()
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		user = domain.NewUser(
			values.NewUserID(),
			name,
			hashedPassword,
		)

		err = a.userRepository.CreateUser(ctx, user)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return user, nil
}
