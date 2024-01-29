package services

import (
	"fmt"
	"go-crud-api/pkg/data/users/data"
	interfaces "go-crud-api/pkg/interfaces/users"
)

type UserService struct {
	Repository interfaces.IUserRepository
}

func (service *UserService) FindById(id uint) (data.UsersData, error) {

	var result data.UsersData

	user, err := service.Repository.FindById(id)
	if err != nil {
		fmt.Printf("Error Finding user by ID %d: %s\n", id, err)

		return data.UsersData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}

func (service *UserService) FindByUsername(username string) (data.UsersData, error) {

	var result data.UsersData

	user, err := service.Repository.FindByUsername(username)
	if err != nil {
		fmt.Printf("Error Finding user by username %s: %s\n", username, err)

		return data.UsersData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}

func (service *UserService) Create(userData data.UsersData) (data.UsersData, error) {

	var result data.UsersData

	user, err := service.Repository.Create(userData)
	if err != nil {
		fmt.Printf("Error When Create User : %s\n", err)

		return data.UsersData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}

func (service *UserService) Update(userData data.UsersData) (data.UsersData, error) {

	var result data.UsersData

	user, err := service.Repository.Update(userData)
	if err != nil {
		fmt.Printf("Error When Create User : %s\n", err)

		return data.UsersData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}
