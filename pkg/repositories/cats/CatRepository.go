package repositories

import (
	dbhandler_interface "go-crud-api/pkg/interfaces"
	interfaces "go-crud-api/pkg/interfaces/cats"
	models "go-crud-api/pkg/models/cats"
	"strconv"

	"go-crud-api/pkg/data/cats/data"

	"github.com/afex/hystrix-go/hystrix"
)

type CatRepositoryWithCircuitBreaker struct {
	CatRepository interfaces.ICatRepository
}

func (repository *CatRepositoryWithCircuitBreaker) FindById(id int) (models.Cats, error) {

	output := make(chan models.Cats, 1)
	hystrix.ConfigureCommand("get_cats_by_id", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_cats_by_id", func() error {

		cat, _ := repository.CatRepository.FindById(id)

		output <- cat
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Cats{}, err
	}
}

func (repository *CatRepositoryWithCircuitBreaker) FindByCatId(id int) (models.Cats, error) {

	output := make(chan models.Cats, 1)
	hystrix.ConfigureCommand("get_cats_by_id", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_cats_by_id", func() error {

		cat, _ := repository.CatRepository.FindByCatId(id)

		output <- cat
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Cats{}, err
	}
}

func (repository *CatRepositoryWithCircuitBreaker) Create(catData data.CatsData) (models.Cats, error) {

	output := make(chan models.Cats, 1)
	hystrix.ConfigureCommand("create_cat", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_cat", func() error {

		cat, _ := repository.CatRepository.Create(catData)

		output <- cat
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.Cats{}, err
	}
}

type CatRepository struct {
	*dbhandler_interface.IDbHandler
}

func (repository *CatRepository) FindById(id int) (models.Cats, error) {

	var cat models.Cats

	repository.DB.Raw("SELECT * FROM cats WHERE id = ?", strconv.Itoa(id)).Scan(&cat)

	return cat, nil
}

func (repository *CatRepository) FindByCatId(id int) (models.Cats, error) {

	var cat models.Cats

	repository.DB.Raw("SELECT * FROM cats WHERE id = ?", strconv.Itoa(id)).Scan(&cat)

	return cat, nil
}

func (repository *CatRepository) Create(catData data.CatsData) (models.Cats, error) {

	var cat models.Cats

	cat.Age = catData.Age
	cat.Gender = catData.Gender
	cat.Name = catData.Name
	cat.Race = catData.Race

	if result := repository.DB.Create(&cat); result.Error != nil {
		return models.Cats{}, result.Error
	}

	return cat, nil
}
