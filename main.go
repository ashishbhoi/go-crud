package main

import (
	"github.com/ashishbhoi/go-crud/initializers"
	"github.com/ashishbhoi/go-crud/models"
	"github.com/ashishbhoi/go-crud/routes"
	"log"
)

func init() {
	initializers.InitializeEnv()
}

func main() {
	models.InitializeDatabase()
	r := routes.InitializeRoutes()
	log.Fatal(
		r.Run(":9000"),
	)
}
