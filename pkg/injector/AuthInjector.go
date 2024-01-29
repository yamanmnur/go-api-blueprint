package injector

import (
	controllers "go-crud-api/pkg/controllers/auth"
	dbhandler_interface "go-crud-api/pkg/interfaces"
	repositories "go-crud-api/pkg/repositories/users"
	auth_services "go-crud-api/pkg/services/auth"
	user_services "go-crud-api/pkg/services/users"
)

func InjectAuthController(dbHandler *dbhandler_interface.IDbHandler) controllers.AuthController {

	userRepository := &repositories.UserRepository{IDbHandler: dbHandler}
	userService := &user_services.UserService{Repository: &repositories.UserRepositoryWithCircuitBreaker{UserRepository: userRepository}}
	authService := &auth_services.AuthService{UserService: userService}
	authController := controllers.AuthController{Service: authService}

	return authController
}
