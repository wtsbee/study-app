package main

import (
	"study-app-api/controller"
	"study-app-api/database"
	"study-app-api/repository"
	"study-app-api/router"
	"study-app-api/usecase"
)

func main() {
	db := database.NewDB()
	database.Migrate(db)
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
