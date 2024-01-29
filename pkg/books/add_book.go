package books

import (
	"fmt"
	"go-crud-api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CrudApiErrResponse struct {
	Message string `json:"message"`
}

type StoreLocation struct {
	Name string `json:"name" validate:"required"`
	City string `json:"city" validate:"required"`
}

type AddBookRequestBody struct {
	Title          string          `json:"title" validate:"required"`
	Author         string          `json:"author" validate:"required"`
	Description    string          `json:"description" validate:"required"`
	StoreLocations []StoreLocation `json:"store_locations,ommitempty" validate:"dive"`
}

func (h handler) AddBook(ctx *gin.Context) {
	body := AddBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// validate the request body
	if err := h.validate.Struct(body); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			message := fmt.Sprintf("%s is required ", e.Field())

			ctx.AbortWithStatusJSON(http.StatusBadRequest, CrudApiErrResponse{Message: message})
			return
		}
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
