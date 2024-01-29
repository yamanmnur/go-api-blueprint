package services

import (
	"errors"
	"fmt"
	auth_data "go-crud-api/pkg/data/auth/data"
	"go-crud-api/pkg/data/auth/requests"
	user_data "go-crud-api/pkg/data/users/data"
	interfaces "go-crud-api/pkg/interfaces/users"
	"go-crud-api/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService interfaces.IUserService
}

func (service *AuthService) Login(request requests.LoginRequest) (auth_data.AuthData, error) {

	var result auth_data.AuthData

	user, err := service.UserService.FindByUsername(request.Username)

	if err != nil {
		fmt.Printf("Error Finding user by ID %s: %s\n", request.Username, err)

		return auth_data.AuthData{}, err
	}

	err = utils.VerifyPassword(request.Password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return auth_data.AuthData{}, err
	}

	fmt.Printf("USER ID FROM TABLE %d", user.Id)
	token, err := utils.GenerateToken(user.Id)

	result.User = user
	result.Token = token

	return result, nil
}

func (service *AuthService) Register(request requests.RegisterRequest) (auth_data.AuthData, error) {

	var result auth_data.AuthData

	user, _ := service.UserService.FindByUsername(request.Username)

	if user != (user_data.UsersData{}) {
		return auth_data.AuthData{}, errors.New("Username Already Exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	var newUser user_data.UsersData

	newUser.Name = request.Name
	newUser.Username = request.Username
	newUser.Password = string(hashedPassword)

	resultUser, err := service.UserService.Create(newUser)
	if err != nil {
		return auth_data.AuthData{}, err
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return auth_data.AuthData{}, err
	}

	result.User = resultUser
	result.Token = token

	return result, nil
}

func (service *AuthService) Profile(id uint) (user_data.UsersData, error) {

	user, err := service.UserService.FindById(id)

	if err != nil {
		return user_data.UsersData{}, errors.New("User Not Found")
	}

	return user, nil
}
