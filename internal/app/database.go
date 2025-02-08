package app

import (
	"fmt"
	"log"
	"short_url/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dbHost := config.GetEnv(config.EnvDatabaseHost, "localhost")
	dbUsername := config.GetEnv(config.EnvDatabseUsername, "root")
	dbPassword := config.GetEnv(config.EnvDatabsePassword, "root")
	dbPort := config.GetEnv(config.EnvDatabasePort, "5432")
	dbName := config.GetEnv(config.EnvDatabaseName, "short_url")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect to database", err)
		panic(err)
	}

	log.Println("Connected to PSQL!")

	return db
}
