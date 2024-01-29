package repositories

import (
	"go-crud-api/pkg/data/users/data"
	dbhandler_interface "go-crud-api/pkg/interfaces"
	interfaces "go-crud-api/pkg/interfaces/users"
	models "go-crud-api/pkg/models/users"

	"github.com/afex/hystrix-go/hystrix"
)

type UserRepositoryWithCircuitBreaker struct {
	UserRepository interfaces.IUserRepository
}

func (repository *UserRepositoryWithCircuitBreaker) FindById(id uint) (models.Users, error) {

	output := make(chan models.Users, 1)
	hystrix.ConfigureCommand("get_users_by_id", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_users_by_id", func() error {

		cat, _ := repository.UserRepository.FindById(id)

		output <- cat
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Users{}, err
	}
}

func (repository *UserRepositoryWithCircuitBreaker) FindByUsername(username string) (models.Users, error) {

	output := make(chan models.Users, 1)
	hystrix.ConfigureCommand("get_users_by_username", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_users_by_username", func() error {

		cat, _ := repository.UserRepository.FindByUsername(username)

		output <- cat
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Users{}, err
	}
}

func (repository *UserRepositoryWithCircuitBreaker) Create(usersData data.UsersData) (models.Users, error) {

	output := make(chan models.Users, 1)
	hystrix.ConfigureCommand("create_user", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_user", func() error {

		user, _ := repository.UserRepository.Create(usersData)

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Users{}, err
	}
}

func (repository *UserRepositoryWithCircuitBreaker) Update(usersData data.UsersData) (models.Users, error) {

	output := make(chan models.Users, 1)
	hystrix.ConfigureCommand("update_user", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("update_user", func() error {

		user, _ := repository.UserRepository.Update(usersData)

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Users{}, err
	}
}

type UserRepository struct {
	*dbhandler_interface.IDbHandler
}

func (repository *UserRepository) FindById(id uint) (models.Users, error) {

	var user models.Users

	repository.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)

	return user, nil
}

func (repository *UserRepository) FindByUsername(username string) (models.Users, error) {

	var user models.Users

	repository.DB.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user)

	return user, nil
}

func (repository *UserRepository) Create(userData data.UsersData) (models.Users, error) {

	var user models.Users

	user.Name = userData.Name
	user.Username = userData.Username
	user.Password = userData.Password

	if result := repository.DB.Create(&user); result.Error != nil {
		return models.Users{}, result.Error
	}

	return user, nil
}

func (repository *UserRepository) Update(userData data.UsersData) (models.Users, error) {

	var user models.Users

	user.Name = userData.Name
	user.Username = userData.Username
	user.Password = userData.Password

	if result := repository.DB.Updates(&user); result.Error != nil {
		return models.Users{}, result.Error
	}

	return user, nil
}
