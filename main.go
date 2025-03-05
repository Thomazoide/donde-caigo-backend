package main

import (
	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/controller"
)

// main function
// @title Donde Caigo APP
// @description API para la aplicacion de Donde Caigo
// @version 0.1
// @host localhost:8080
// @BasePath /
// @schemes http
// @Contact.name Thomas Tellerias
// @Contact.email ttellerias01@outlook.com
// @License.name Donde Caigo
// @ExternalDocs.description Basado en OpenAPI 3.0
// @ExternalDocs.url https://swagger.io/resources/open-api/
func main() {
	config.ConnectDB()
	server := controller.NewAPIServer(":8080")
	server.RunServer()
}
