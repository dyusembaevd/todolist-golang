package main

import (
	_ "github.com/lib/pq"
	"github.com/mohito22/todo-app-golang"
	"github.com/mohito22/todo-app-golang/pkg/handler"
	"github.com/mohito22/todo-app-golang/pkg/repository"
	"github.com/mohito22/todo-app-golang/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing config: %s", err)
	}

	db, err := repository.NewPostgressDB(repository.Config{
		Host:     "localhost",
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error while connecting to DB: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
