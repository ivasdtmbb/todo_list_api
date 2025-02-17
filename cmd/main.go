package main

import (
	"github.com/ivasdtmbb/todo_list_project/internal/todo"
	"github.com/ivasdtmbb/todo_list_project/pkg/handler"
	"github.com/ivasdtmbb/todo_list_project/pkg/repository"
	"github.com/ivasdtmbb/todo_list_project/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"os"
)

// @title ToDo App API
// @version 1.0.0
// @description API Server for ToDoList Application

// @host localhost:8000
// @Basepath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization


func main() {
    logrus.SetFormatter(new(logrus.JSONFormatter))

    if err := initConfig(); err != nil {
        logrus.Fatalf("error initializing configs: %s", err.Error())
    }

    if err := godotenv.Load(); err != nil {
        logrus.Fatal("error loading env variables: %s", err.Error())
    }

    db, err := repository.NewPostgressDB(repository.Config{
        Host:       viper.GetString("db.host"),
        Port:       viper.GetString("db.port"),
        Username:   viper.GetString("db.username"),
        DBName:     viper.GetString("db.dbname"),
        SSLMode:    viper.GetString("db.sslmode"),
        Password:   os.Getenv("DB_PASSWORD"),
    })

    if err != nil {
        logrus.Fatalf("failed to initiflize db: %s", err.Error())
    }

    repos := repository.NewRepository(db)
    services := service.NewService(repos)
    handlers := handler.NewHandler(services)

    srv := new(todo.Server)
    // go func() {
    //     if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
    //         logrus.Fatalf("error occured while running http server: %s", err.Error())
    //     }
    // }()
    if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
        logrus.Fatalf("error occured while running http server: %s", err.Error())
    }
}

func initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}
