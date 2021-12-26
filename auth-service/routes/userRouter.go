package routes

import (
	controller "github.com/bipuldan/MeAndYou/auth-service/controllers"
	"github.com/bipuldan/MeAndYou/auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incominRoutes *gin.Engine) {
	incominRoutes.Use(middleware.Authenticate())
	incominRoutes.GET("/users", controller.GetUsers())
	incominRoutes.GET("/user:user_id", controller.GetUser())
}
