package main

import (
	"study-app-api/controller"
	"study-app-api/database"
	"study-app-api/repository"
	"study-app-api/router"
	"study-app-api/usecase"
	"study-app-api/validator"
)

func main() {
	db := database.NewDB()
	database.Migrate(db)
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
