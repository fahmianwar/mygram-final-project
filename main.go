package main

import (
	"mygram-final-project/database"
	"mygram-final-project/router"
	"os"
)

// @title MyGram API
// @version 1.0
// @description This is a sample service for managing MyGram API
// @termOfService http://swagger.io/terms/
// @contact.name Github Repo
// @contact.url https://github.com/fahmianwar/mygram-final-project
// @host localhost
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	database.StartDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
