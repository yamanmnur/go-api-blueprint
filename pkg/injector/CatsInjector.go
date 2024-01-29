package injector

import (
	controllers "go-crud-api/pkg/controllers/cats"
	dbhandler_interface "go-crud-api/pkg/interfaces"
	repositories "go-crud-api/pkg/repositories/cats"
	services "go-crud-api/pkg/services/cats"
)

type IServiceContainer interface {
	InjectCatController() controllers.CatController
}

type kernel struct{}

func InjectCatController(dbHandler *dbhandler_interface.IDbHandler) controllers.CatController {

	catRepository := &repositories.CatRepository{IDbHandler: dbHandler}
	// catService := &services.CatService{repository: &repositories.CatRepositoryWithCircuitBreaker{CatRepository: catRepository}}
	catService := &services.CatService{Repository: &repositories.CatRepositoryWithCircuitBreaker{CatRepository: catRepository}}
	catController := controllers.CatController{ICatService: catService}

	return catController
}
