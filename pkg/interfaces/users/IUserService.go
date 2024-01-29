package interfaces

import "go-crud-api/pkg/data/users/data"

type IUserService interface {
	FindById(id uint) (data.UsersData, error)
	FindByUsername(username string) (data.UsersData, error)
	Create(userData data.UsersData) (data.UsersData, error)
	Update(userData data.UsersData) (data.UsersData, error)
}
