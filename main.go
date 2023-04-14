package main

import (
	"mygram-final-project/database"
	"mygram-final-project/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
