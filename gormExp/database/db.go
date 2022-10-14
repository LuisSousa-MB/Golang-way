package database

import (
	"gormExp/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	log.Println("Conectando")
	conexao := "host=172.19.0.2 user=root password=root dbname=gormExp port=5432 sslmode=disable" // host capturado com: docker-compose exec postgress sh(esse é o serviço que vc deseja caputrar a porta), depois: hostname -i para exibir o ip
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: ", err, "\n...")
	}
	log.Println("Connected")

	//var curso models.Curso
	//DB.First(&curso, 1) // Buscando pelo id

	//DB.Model(&curso).Update("Nome", "Go: 2.0  Gormado") //Atualizando update sql

	//DB.Delete(&curso, 4) // Excluindo pelo id

	DB.Create(&models.Curso{Nome: "Go: Fundamentos de uma aplicação web", Area: "Programacao", Professor: "Guilherme Lima"})

}
