package main

import (
	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/controller"
)

func main() {
	config.ConnectDB()
	server := controller.NewAPIServer(":8080")
	server.RunServer()
}
