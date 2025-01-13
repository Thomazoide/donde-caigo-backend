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

var db *gorm.DB

func ConnectDB() error {
	godotenv.Load()
	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	var err error
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	log.Println("Conectado a BBDD")
	return nil
}

func GetInstance() *gorm.DB {
	return db
}
