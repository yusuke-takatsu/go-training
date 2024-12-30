package main

import (
	"github.com/joho/godotenv"
	"github.com/yusuke-takatsu/go-training/config/database"
	"github.com/yusuke-takatsu/go-training/infra/user/repository"
	"github.com/yusuke-takatsu/go-training/interface/user/handler"
	"github.com/yusuke-takatsu/go-training/interface/user/router"
	"github.com/yusuke-takatsu/go-training/middleware"
	"github.com/yusuke-takatsu/go-training/service/user/usecase"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv()
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewHandler(userService)

	r := router.NewRouter(userHandler)
	r.Use(middleware.Logging)

	log.Println("Starting server port is ", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), r); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
