package main

import (
	"github.com/timurkaev/todo-app-golang"
	"github.com/timurkaev/todo-app-golang/pkg/handler"
	"github.com/timurkaev/todo-app-golang/pkg/repository"
	"github.com/timurkaev/todo-app-golang/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running https server: %s", err.Error())
	}
}
