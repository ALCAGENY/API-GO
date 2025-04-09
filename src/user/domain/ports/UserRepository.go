package ports

import "api-go/src/user/domain/entities"

type IUserRepository interface {
	Create(user entities.User) (entities.User, error)
	GetByID(id int64) (entities.User, error)
	Delete(id int64) (bool, error)
	GetByEmail(email string) (entities.User, error)
}