package routes

import (
	controller "github.com/bipuldan/MeAndYou/auth-service/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incominRoutes *gin.Engine) {
	incominRoutes.POST("/user/signup", controller.Signup())
	incominRoutes.POST("/user/login", controller.Login())
}
