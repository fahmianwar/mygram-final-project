package main

import (
	"mygram-final-project/database"
	"mygram-final-project/router"
	"os"
)

// @title MyGram API
// @version 1.0
// @description This is a sample service for managing books
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host https://mygram-final-project-production.up.railway.app/
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
