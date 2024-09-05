package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"socialmood/api/controllers"
	"socialmood/internal/config"
	"socialmood/internal/db/postgres"
	"socialmood/internal/repositories"
	"socialmood/internal/server"
	userUseCases "socialmood/internal/usecases/user-usecases"
	"socialmood/pkg/logger"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	err := godotenv.Load("../../.env")
	conf := config.NewConfig()

	logg := logger.New(conf.Logger.Level)
	db, err := postgres.NewPostgresqlRepository(context.Background())
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}
	logg.Info("Database init")

	// Repositories
	userRepository := repositories.NewUserRepository(db.DB)

	// UseCases
	userCases := userUseCases.NewUserUseCases(userRepository, &conf.JWT)

	// Controllers
	userController := controllers.NewUserController(userCases)
	authController := controllers.NewAuthController(userCases)

	go func() {
		err = server.NewServer(conf, userController, authController)
		if err != nil {
			log.Fatalf("Server init error" + err.Error())
		}
	}()

	<-stopChan
}
