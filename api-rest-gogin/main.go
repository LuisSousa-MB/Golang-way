package main

import (
	"api-rest-gogin/database"
	"api-rest-gogin/routes"
)

func main() {
	database.ConectaComDB()
	routes.HandleRequests()
}
