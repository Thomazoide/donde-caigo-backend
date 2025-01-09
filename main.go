package main

import (
	"fmt"
	"net/http"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("servidor en http://localhost:5000", db.Stats())
	http.ListenAndServe(":5000", nil)
}
