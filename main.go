package main

import (
	"mygram-final-project/database"
	"mygram-final-project/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
