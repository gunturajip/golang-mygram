package main

import (
	"golang-mygram/database"
	"golang-mygram/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
