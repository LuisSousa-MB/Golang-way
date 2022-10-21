package database

import (
	"log"
	"mybank/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	//log.Println("Conectando")
	conexao := "host=172.22.0.2 user=root password=root dbname=root port=5432 sslmode=disable" // host capturado com: docker-compose exec postgress sh(esse é o serviço que vc deseja caputrar a porta), depois: hostname -i para exibir o ip
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: ", err, "\n...")
	}
	//log.Println("Connected")

	DB.AutoMigrate(&models.Titular{})
	DB.AutoMigrate(&models.ContaCorrente{})
	DB.AutoMigrate(&models.ContaPoupanca{})

	//var curso models.Curso
	//DB.First(&curso, 1) // Buscando pelo id

	//DB.Model(&curso).Update("Nome", "Go: 2.0  Gormado") //Atualizando update sql

	//DB.Delete(&curso, 4) // Excluindo pelo id
}
