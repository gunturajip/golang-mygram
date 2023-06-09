package database

import (
	"fmt"
	"golang-mygram/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// host     = "localhost"
	host = os.Getenv("PGHOST")
	// user     = "postgres"
	user = os.Getenv("PGUSER")
	// password = "postgres"
	password = os.Getenv("PGPASSWORD")
	// port     = "5432"
	port = os.Getenv("PGPORT")
	// dbname   = "db-go-sql"
	dbname = os.Getenv("PGDATABASE")
	db     *gorm.DB
	err    error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("successfully connected to database🔥")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Socialmedia{}, models.Comment{})
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if sqlDB != nil {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal("error closing to database :", err)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
