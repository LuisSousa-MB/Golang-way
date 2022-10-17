module api-rest-gogin

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/assert/v2 v2.0.1
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.8.0 // indirect -- go get github.com/stretchr/testify  -- go test -run (NomeDaFunçãoTest) exec o teste unitario
	gopkg.in/validator.v2 v2.0.1
	gorm.io/driver/postgres v1.4.4
	gorm.io/gorm v1.24.0
)
