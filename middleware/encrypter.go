package middleware

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Encrypter struct {
	pepper     string
	saltrounds int
}

func NewEncrypter() *Encrypter {
	godotenv.Load()
	saltrounds, err := strconv.Atoi(os.Getenv("SALT"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return nil
	}
	return &Encrypter{
		pepper:     pepper,
		saltrounds: saltrounds,
	}
}

func (e *Encrypter) HashPassword(password string) (string, error) {
	passwordAndPepper := password + e.pepper
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordAndPepper), e.saltrounds)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (e *Encrypter) VerifyPassword(password string, hashedPassword string) bool {
	passwordAndPepper := password + e.pepper
	fmt.Println(passwordAndPepper)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordAndPepper))
	fmt.Println(err)
	return err == nil
}
