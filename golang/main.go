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
	taskListRepository := repository.NewTaskListRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskDetailRepository)
	taskDetailUsecase := usecase.NewTaskDetailUsecase(taskDetailRepository)
	taskListUsecase := usecase.NewTaskListUsecase(taskListRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	taskDetailController := controller.NewTaskDetailController(taskDetailUsecase)
	taskListController := controller.NewTaskListController(taskListUsecase)
	e := router.NewRouter(userController, taskController, taskDetailController, taskListController)
	e.Logger.Fatal(e.Start(":8080"))
}
