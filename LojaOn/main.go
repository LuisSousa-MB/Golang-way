package main

import (
	"LojaOn/routes"
	"net/http"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
