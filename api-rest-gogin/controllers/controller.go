package controllers

import (
	"api-rest-gogin/database"
	"api-rest-gogin/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosAlunos godoc
// @Summary Recupera e exibe alunos.
// @Description Rota que recupera e exibe todos os alunos cadastrados no db.
// @Tags Alunos
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) { //caputra a informação do endpoint e retorna usando uma mensagem em JSON
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Api diz:": "Eai " + nome + ", tudo certo!?",
	})
}

// CriaNovoAluno godoc
// @Summary Cria um novo aluno.
// @Description Rota que cria um novo aluno a partir de parâmetros do tipo objeto json e salva no db.
// @Tags Alunos
// @Produce json
// @Param aluno body models.AlunoPostReq true "Modelo de Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno // criando uma variavel do tipo aluno para armazenar o conteudo a ser registrado do db
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoID godoc
// @Summary Recupera e exibe aluno a partir do ID.
// @Description Rota que recupera e exibe  um aluno a partir do ID passado (URL/path).
// @Tags Alunos
// @Produce json
// @Param ID path int true "ID do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/{ID} [get]
func BuscaAlunoID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": " Aluno não encotrado..."})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoCPF godoc
// @Summary Recupera e exibe aluno a partir do CPF.
// @Description Rota que recupera e exibe  um aluno a partir do CPF passado (URL/path).
// @Tags Alunos
// @Produce json
// @Param CPF path int true "CPF do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/cpf/{CPF} [get]
func BuscaAlunoCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": " Aluno não encotrado..."})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

// BuscaAlunoRg godoc
// @Summary Recupera e exibe aluno a partir do RG.
// @Description Rota que recupera e exibe  um aluno a partir do RG passado (URL/path).
// @Tags Alunos
// @Produce json
// @Param RG path int true "RG do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/rg/{RG} [get]
func BuscaAlunoRg(c *gin.Context) {
	var aluno models.Aluno
	rg := c.Param("rg")
	database.DB.Where(&models.Aluno{RG: rg}).First((&aluno))

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": " Aluno não encontrado..."})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
// @Summary Deleta o aluno do db.
// @Description Rota que deleta um aluno do db a partir de seu ID passado através da URL.
// @Tags Alunos
// @Param ID path int true "ID do Aluno"
// @Success 200 "Sucesso!"
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/{ID} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso!"})
}

// EditaAluno godoc
// @Summary Edita os dados do aluno.
// @Description Rota que edita os dados de um aluno a partir de seu ID passado através da URL.
// @Tags Alunos
// @Param ID path int true "ID do Aluno"
// @Param aluno body models.AlunoPostReq true "Modelo de Aluno"
// @Success 200 {object} models.Aluno "Sucesso!"
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/{ID} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&models.Aluno{}).Where("id = ?", id).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
	//fmt.Println("Isto é aluno:", aluno)
}

// ExibePagIndex godoc
// @Summary Pagina principal de alunos.
// @Description Rota que recupera a pagina principal com dados dos alunos cadastrados no sistemab.
// @Tags Alunos
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /index [get]
func ExibePagIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}
func Limbo(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func ExibeDocs(c *gin.Context) {
	c.HTML(http.StatusOK, "docs.html", nil)
}

type Any struct {
	Value interface{}
}

func ExibeJson(c *gin.Context) {
	// Abra o arquivo JSON
	jsonFile, err := ioutil.ReadFile("./docs/swagger.json")
	if err != nil {
		panic(err)
	}
	c.JSON(200, string(jsonFile))
}
