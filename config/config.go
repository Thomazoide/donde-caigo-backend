package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	godotenv.Load()
	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db.AutoMigrate(&models.Post{}, &models.User{})
	log.Println("Conectado a BBDD")
	return db, nil
}
