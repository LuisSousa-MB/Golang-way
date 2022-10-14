package main

import (
	"gormExp/database"
	"gormExp/models"
	"log"
)

func main() {
	log.Println("iniciando conexão...")
	database.ConectaComDB()
	database.DB.AutoMigrate(&models.Curso{}) //Criar tabela a partir da struct
	//database.DB.Create(&models.Curso{Id: 4, Nome: "Go: Fundamentos de uma aplicacao web", Area: "Programacao", Professor: "Guilherme Lima"})

}
