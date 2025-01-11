package main

import (
	"github.com/Thomazoide/donde-caigo-backend/controller"
)

func main() {
	server := controller.NewAPIServer(":4000")
	server.RunServer()
}
