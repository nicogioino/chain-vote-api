package repositories

import (
	 "chain-vote-api/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDataBase() {

	url := os.Getenv("DATABASE_URL")

	dbUrl := "postgres://test:test@" + url + ":5432/chain-vote-db"

	DbConnection, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ")
		DB = DbConnection
	}

	err = DB.AutoMigrate(&models.User{}, &models.Election{})

	if err != nil {
		fmt.Println("Cannot migrate database ")
		log.Fatal("migration error:", err)
	}
}
