package routers

import (
	"github.com/gin-gonic/gin"

	"go-crud-api/pkg/common/middlewares"
	controllers "go-crud-api/pkg/controllers/auth"
)

func InitAuthRouter(router *gin.Engine, controller controllers.AuthController) {

	routes := router.Group("/auth")
	routes.POST("/login", controller.Login)
	routes.POST("/register", controller.Register)

	protectedRoutes := router.Group("/profile")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	protectedRoutes.GET("/", controller.Profile)
}
