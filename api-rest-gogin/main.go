package main

import (
	"api-rest-gogin/database"
	"api-rest-gogin/routes"

	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Exemplo para teste do Swagger
// @version 3.0
// @description  Api de teste de modelagem para documentação de API´s no Swagger.
// @version 0.0.1
// @termsOfService https://policies.google.com/terms?hl=pt-BR&fg=1
// @contact.name   Suport Devs
// @contact.url    http://www.exemplo.com
// @contact.email  contato@example.com
// @license.name  Licença: GPLv3
// @license.url   https://www.gnu.org/licenses/gpl-3.0.html
// @host      localhost:8080
// @BasePath  /index
// @securityDefinitions.basic  BasicAuth

func main() {
	database.ConectaComDB()
	routes.HandleRequests()

}
