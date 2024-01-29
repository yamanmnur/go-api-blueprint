package interfaces

import (
	"go-crud-api/pkg/data/cats/data"
)

type ICatService interface {
	// GetCats() (data.CatsData, error)
	GetCatsById(id int) (data.CatsData, error)
}
