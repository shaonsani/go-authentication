package router

import (
	"go-authentication/controller"

	_ "go-authentication/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AuthRouter(authController *controller.AuthController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/auth")

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/uploads", "./uploads")

	return service
}
