package main

import (
	"api/rest-Go/database"
	"api/rest-Go/routes"
	"fmt"
)

func main() {
	database.ConectaComDB()
	fmt.Println("Iniciando  o servidor Rest com GO")
	routes.HandleRequest()
}
