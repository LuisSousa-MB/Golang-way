package routes

import (
	"net/http"
	"projects/mybank/controllers"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/authentic", controllers.Autenticar)
}
