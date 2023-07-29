package main

import (
	"github.com/joaoscorissa/go-api-rest/database"
	"github.com/joaoscorissa/go-api-rest/models"
	"github.com/joaoscorissa/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "João", Historia: "João..."},
		{Id: 2, Nome: "Maria", Historia: "Maria..."},
	}
	database.ConnectDB()
	routes.HandleRequest()
}
