package routes

import (
	"github.com/cristopherecruz/go-gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequests() {

	r := gin.Default()

	r.GET("/alunos", controllers.ExibirAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.ExibirAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PUT("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAunoCPF)

	err := r.Run()
	if err != nil {
		log.Panic(err.Error())
	}

}
