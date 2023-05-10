package main

import (
	"github.com/cristopherecruz/go-gin-api-rest/database"
	"github.com/cristopherecruz/go-gin-api-rest/routes"
)

func main() {
	database.ConectarBancoDados()

	routes.HandleRequests()
}
