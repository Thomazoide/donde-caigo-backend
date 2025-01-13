package middleware

import (
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
	saltrounds, err := strconv.Atoi("SALT")
	if err != nil {
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
	err := bcrypt.CompareHashAndPassword([]byte(passwordAndPepper), []byte(hashedPassword))
	return err == nil
}