package router

import (
	"mygram-final-project/controllers"
	"mygram-final-project/middlewares"

	"github.com/gin-gonic/gin"

	_ "mygram-final-project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description This is a sample service for managing books
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		// Create
		commentRouter.POST("/", controllers.CreateComment)
		// GetAll
		commentRouter.GET("/", controllers.GetComments)
		// Get
		commentRouter.GET("/:commentId", controllers.GetComment)
		// Update
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		// Delete
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.GET("/:photoId", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	socialMediaRouter := r.Group("/socialMedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedias)
		socialMediaRouter.GET("/:socialMediaRouterId", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaRouterId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaRouterId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
