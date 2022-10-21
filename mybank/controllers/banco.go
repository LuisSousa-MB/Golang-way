package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

/* func Autenticar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cpf := r.FormValue("CPF")
		pass := r.FormValue("typePasswordX")

		db := db()

		passCheck, err := db.Query("select senha from cliente where cpf=$1", cpf)

		if err != nil {
			panic(err.Error())
		}
		var passx int

		err = passCheck.Scan(&passx)

		passConv, err := strconv.Atoi(pass)

		if err != nil {
			panic(err.Error())
		}

		if passx == passConv {
			fmt.Println("Deu certo!")
			http.Redirect(w, r, "/sucess", 301)

		} else {
			http.Redirect(w, r, "/fail", 301)

		}

	} else {
		http.Redirect(w, r, "/fail", 301)

	}

}
*/
