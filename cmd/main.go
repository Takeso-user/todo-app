package main

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/handler"
	"github.com/Takeso-user/todo-app/pkg/repository"
	"github.com/Takeso-user/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
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
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runing http server: %s", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
