package main

import (
	"dietcalc/cmd/DietCalc/api"
	"dietcalc/internal/config"
)

func main() {
	config.Cfg = config.MustLoad()
	api.Run()
}

// TODO:
// 1) Product must be created only by user (jwt)
// 2) Add page + limit in getall
// 3) Make CRUD for meal
// 4) Make OpenAPI specification
