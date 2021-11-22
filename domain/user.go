package domain

import "github.com/mazrean/todoList/domain/values"

type User struct {
	id             values.UserID
	name           values.UserName
	hashedPassword values.UserHashedPassword
}

func NewUser(
	id values.UserID,
	name values.UserName,
	hashedPassword values.UserHashedPassword,
) *User {
	return &User{
		id:             id,
		name:           name,
		hashedPassword: hashedPassword,
	}
}
