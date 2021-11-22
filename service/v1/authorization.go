package service

import "github.com/mazrean/todoList/repository"

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
