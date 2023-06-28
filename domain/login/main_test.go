package login

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	cwd, err := os.Getwd()
	log.Println(cwd, err)

	err = godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//Run tests
	code := m.Run()

	log.Println("Do stuff AFTER the tests!")
	os.Exit(code)
}
