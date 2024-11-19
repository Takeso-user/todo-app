package main

import (
	"context"
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/handler"
	"github.com/Takeso-user/todo-app/pkg/repository"
	"github.com/Takeso-user/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.SetFormatter(new(logrus.JSONFormatter))
		logrus.Fatalf("error initializing configs: %s", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}
	connectPostgres, err := repository.NewPostgresDB(repository.Config{
		Hostname: viper.GetString("db.hostname"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err)
	}

	repos := repository.NewRepository(connectPostgres)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)

	go func() {
		if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while runing http server: %s", err)
		}
	}()
	logrus.Printf("todo-app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Printf("todo-app shutting down")
	err = srv.Stop(context.Background())
	if err != nil {
		logrus.Error("error occured on server shutting down: %s", err)
		return
	}
	err = connectPostgres.Close()
	if err != nil {
		logrus.Errorf("error occured on db connection close: %s", err)
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
