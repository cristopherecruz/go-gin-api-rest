package controllers

import (
	"github.com/cristopherecruz/go-gin-api-rest/database"
	"github.com/cristopherecruz/go-gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExibirAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.Db.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {

	nome := c.Params.ByName("nome")

	c.JSON(http.StatusOK, gin.H{"API diz:": "E ai " + nome + ", tudo beleza?"})
}

func CriarAluno(c *gin.Context) {

	var aluno models.Aluno

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.Db.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func ExibirAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.Db.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, "Aluno não encontrado!")
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.Db.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, "Aluno não encontrado!")
		return
	}

	database.Db.Delete(&aluno, id)

	c.JSON(http.StatusNoContent, "Aluno deletado com sucesso!")
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.Db.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, "Aluno não encontrado!")
		return
	}

	database.Db.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

func BuscarAunoCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	database.Db.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, "Aluno não encontrado!")
		return
	}

	c.JSON(http.StatusOK, aluno)
}
