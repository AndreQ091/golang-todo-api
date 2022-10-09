package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	todo "github.com/AndreQ091/golang-todo"
	"github.com/AndreQ091/golang-todo/internal/handler"
	"github.com/AndreQ091/golang-todo/internal/repository"
	"github.com/AndreQ091/golang-todo/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("read config error %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("read env variables error %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("database connect error %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	go func() {

		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("server error %s", err.Error())
		}
	}()

	logrus.Print("App is running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGILL)
	<-quit

	logrus.Print("App is exiting...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("shutdown error $s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("db close error $s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
