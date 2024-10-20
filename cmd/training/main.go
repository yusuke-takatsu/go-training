package main

import (
	"github.com/joho/godotenv"
	"github.com/yusuke-takatsu/go-training/config/database"
	"log"
)

func main() {
	loadEnv()
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
