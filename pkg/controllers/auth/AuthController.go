package controllers

import (
	"fmt"
	base_responses "go-crud-api/pkg/common/responses"
	"go-crud-api/pkg/data/auth/requests"
	interfaces "go-crud-api/pkg/interfaces/auth"
	"go-crud-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service interfaces.IAuthService
}

func (controller *AuthController) Login(ctx *gin.Context) {

	var requestBody requests.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authData, err := controller.Service.Login(requestBody)

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
			Message: "Success To Login",
			Code:    "200",
		},
		Data: authData,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *AuthController) Register(ctx *gin.Context) {

	var requestBody requests.RegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authData, err := controller.Service.Register(requestBody)

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
			Message: "Success To Register",
			Code:    "200",
		},
		Data: authData,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *AuthController) Profile(ctx *gin.Context) {
	user_id, err := utils.ExtractTokenID(ctx)
	fmt.Printf("ID USER %d", user_id)
	userData, err := controller.Service.Profile(user_id)

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
			Message: "Success To Register",
			Code:    "200",
		},
		Data: userData,
	}

	ctx.JSON(http.StatusOK, response)
}
