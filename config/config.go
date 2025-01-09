package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	godotenv.Load()
	var connectionString string = fmt.Sprintf("%s%s@tcp(%s:%s)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Conectado a BBDD")
	return db, nil
}
