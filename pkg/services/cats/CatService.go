package services

import (
	"fmt"
	data "go-crud-api/pkg/data/cats/data"
	interfaces "go-crud-api/pkg/interfaces/cats"
)

type CatService struct {
	Repository interfaces.ICatRepository
}

func (service *CatService) GetCatsById(id int) (data.CatsData, error) {

	var result data.CatsData

	cat, err := service.Repository.FindByCatId(id)
	if err != nil {
		fmt.Printf("Error Finding cat by ID %d: %s\n", id, err)

		return data.CatsData{}, err
	}

	result.Id = cat.ID
	result.Age = cat.Age
	result.Gender = cat.Gender
	result.Name = cat.Name
	result.Race = cat.Name

	newData := &result
	newData.Id = 0
	service.Repository.Create(result)

	return result, nil
}
