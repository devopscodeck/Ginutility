package main

import (
	"github.com/devopscodeck/Ginutility/database"
	"github.com/devopscodeck/Ginutility/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
