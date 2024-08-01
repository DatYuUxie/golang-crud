package main

import (
	"CRUDProject/cfg"
	"CRUDProject/controllers"
	"CRUDProject/models"
	"net/http"
)

func main() {
	// Database connection
	cfg.ConnectDB()
	models.Migrate()

	r := controllers.SetupRouter()

	http.ListenAndServe(":8080", r)
}
