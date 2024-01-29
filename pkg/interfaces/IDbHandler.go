package interfaces

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IDbHandler struct {
	DB       *gorm.DB
	Validate *validator.Validate
}
