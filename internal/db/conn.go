package db

import (
	"fmt"
	"log"
	"os"
	entitites "websocket-service/internal/entities"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to system environment variables")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = DB.AutoMigrate(&entitites.Event{})
	if err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}

	fmt.Println("Banco de dados migrado com sucesso!")
}
