package main

import (
	"dietcalc/cmd/DietCalc/api"
	"dietcalc/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("No .env file were found in root of project")
	}
}

func main() {
	config.Cfg = config.MustLoad()
	api.Run()
}

// TODO:
// 1) Product must be created only by user (jwt)
// 2) Add page + limit in getall
// 3) Make CRUD for meal
// 4) Make OpenAPI specification
