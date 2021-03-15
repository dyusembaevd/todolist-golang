package main

import (
	"log"

	"github.com/mohito22/todo-app-golang"
	"github.com/mohito22/todo-app-golang/pkg/handler"
	"github.com/mohito22/todo-app-golang/pkg/repository"
	"github.com/mohito22/todo-app-golang/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server: %s", err)
	}
}
