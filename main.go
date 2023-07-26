package main

import (
	"github.com/ashishbhoi/go-crud/database"
	"github.com/ashishbhoi/go-crud/routes"
)

func main() {
	database.InitialMigration()
	routes.InitializeRouter(":9000")
}
