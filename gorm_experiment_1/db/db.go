package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Note struct {
	gorm.Model
	Title       string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:text"`
}

var DB *gorm.DB

func InitPostgres() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	err = DB.DB().Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")

	// Migrate
	DB.AutoMigrate(&Note{})

}

// Close closes the database connection
func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Fatalf("failed to close database connection: %v", err)
		}
	}
}
