package main

import (
	"API-RESTFUL/database"
	"API-RESTFUL/models"
	"API-RESTFUL/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Histoia2"},
		{Id: 2, Nome: "Nome2", Historia: "Histoia1"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciado Servidor")
	routes.HandleRequest()
}
