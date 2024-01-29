package main

import (
	"go-crud-api/pkg/common/db"
	injector "go-crud-api/pkg/injector"
	interfaces "go-crud-api/pkg/interfaces"
	auth_routers "go-crud-api/pkg/routers/auth"
	cat_routers "go-crud-api/pkg/routers/cats"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// add env variables as needed
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)
	v := validator.New()

	injectDbHandler := &interfaces.IDbHandler{DB: dbHandler, Validate: v}

	playerController := injector.InjectCatController(injectDbHandler)
	cat_routers.InitCatRouter(router, playerController)

	authController := injector.InjectAuthController(injectDbHandler)
	auth_routers.InitAuthRouter(router, authController)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)
}
