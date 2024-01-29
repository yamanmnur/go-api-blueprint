package routers

import (
	"github.com/gin-gonic/gin"

	controllers "go-crud-api/pkg/controllers/cats"
)

func InitCatRouter(router *gin.Engine, controller controllers.CatController) {

	routes := router.Group("/cats")
	routes.GET("/:id", controller.GetById)
}
