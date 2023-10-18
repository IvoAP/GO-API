package main

import (
	"API-RESTFUL/database"
	"API-RESTFUL/routes"
	"fmt"
)

func main() {

	
	
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciado Servidor")
	routes.HandleRequest()
}
