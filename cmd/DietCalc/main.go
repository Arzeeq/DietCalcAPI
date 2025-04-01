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
	config.Cfg = config.MustLoad()
}

func main() {
	api.Run()
}

// TODO:
// 1) Add JWT Authorization
// 2) Finish CRUD for user
// 3) Make CRUD for product
// 4) Make CRUD for meal
// 5) Make OpenAPI specification
