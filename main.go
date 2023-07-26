package main

import (
	"github.com/ashishbhoi/go-crud/initializers"
	"github.com/ashishbhoi/go-crud/models"
	"github.com/ashishbhoi/go-crud/routes"
)

func init() {
	initializers.InitializeEnv()
}

func main() {
	models.InitializeDatabase()
	routes.InitializeRoutes(":9000")
}
