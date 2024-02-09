package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "COPPED"
)

func GetConnection() *gorm.DB {
	log.Println("TESET")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	log.Printf("psql: %s", psqlInfo)
	db, err := gorm.Open(postgres.Open(psqlInfo))

	if err != nil {
		log.Printf("error: %v", err)
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")
	return db
}
