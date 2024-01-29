package interfaces

import (
	auth_data "go-crud-api/pkg/data/auth/data"
	user_data "go-crud-api/pkg/data/users/data"

	"go-crud-api/pkg/data/auth/requests"
)

type IAuthService interface {
	Login(request requests.LoginRequest) (auth_data.AuthData, error)
	Register(userData requests.RegisterRequest) (auth_data.AuthData, error)
	Profile(id uint) (user_data.UsersData, error)
}
