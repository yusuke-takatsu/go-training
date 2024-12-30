package main

import (
	"github.com/joho/godotenv"
	"github.com/yusuke-takatsu/go-training/config/database"
	"github.com/yusuke-takatsu/go-training/infra/user/repository"
	"github.com/yusuke-takatsu/go-training/interface/user/handler"
	"github.com/yusuke-takatsu/go-training/interface/user/router"
	"github.com/yusuke-takatsu/go-training/middleware"
	"github.com/yusuke-takatsu/go-training/service/user/usecase"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	loadEnv()
	if os.Getenv("APP_ENV") == "production" {
		return
	}

	f, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(io.MultiWriter(os.Stdout, f))
}

func main() {
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
