package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {

	conexao := "host=172.18.0.2 user=root password=root dbname=root port=5432 sslmode=disable" // host capturado com: docker-compose exec postgress sh(esse é o serviço que vc deseja caputrar a porta), depois: hostname -i para exibir o ip
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: ", err, "\n...")
	}
}
