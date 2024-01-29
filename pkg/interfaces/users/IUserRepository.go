package interfaces

import (
	"go-crud-api/pkg/data/users/data"
	models "go-crud-api/pkg/models/users"
)

type IUserRepository interface {
	FindById(id uint) (models.Users, error)
	FindByUsername(username string) (models.Users, error)
	Create(catData data.UsersData) (models.Users, error)
	Update(catData data.UsersData) (models.Users, error)
}
