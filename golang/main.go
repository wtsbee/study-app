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
	taskRepository := repository.NewTaskRepository(db)
	taskDetailRepository := repository.NewTaskDetailRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskDetailRepository)
	taskDetailUsecase := usecase.NewTaskDetailUsecase(taskDetailRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	taskDetailController := controller.NewTaskDetailController(taskDetailUsecase)
	e := router.NewRouter(userController, taskController, taskDetailController)
	e.Logger.Fatal(e.Start(":8080"))
}
