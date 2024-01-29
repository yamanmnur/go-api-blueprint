package interfaces

import (
	"go-crud-api/pkg/data/cats/data"
	models "go-crud-api/pkg/models/cats"
)

type ICatRepository interface {
	FindById(id int) (models.Cats, error)
	FindByCatId(id int) (models.Cats, error)
	// DeleteById(id int) error
	Create(catData data.CatsData) (models.Cats, error)
	// GetCats() (data.CatsData, error)
}
