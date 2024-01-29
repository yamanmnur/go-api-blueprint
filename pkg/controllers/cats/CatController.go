package controllers

import (
	"fmt"
	base_responses "go-crud-api/pkg/common/responses"
	interfaces "go-crud-api/pkg/interfaces/cats"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	interfaces.ICatService
}

func (controller *CatController) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	cat_id, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cat, err := controller.GetCatsById(cat_id)

	if err != nil {
		response := base_responses.BasicResponse{
			MetaData: base_responses.MetaData{
				Status:  "error",
				Message: err.Error(),
				Code:    "500",
			},
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := base_responses.GenericResponse{
		MetaData: base_responses.MetaData{
			Status:  "success",
			Message: "Success To Get Data Cat",
			Code:    "200",
		},
		Data: cat,
	}

	ctx.JSON(http.StatusOK, response)
}
