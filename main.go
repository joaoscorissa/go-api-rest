package main

import (
	"github.com/joaoscorissa/go-api-rest/database"
	"github.com/joaoscorissa/go-api-rest/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequest()
}
