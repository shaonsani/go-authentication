package router

import (
 "go-authentication/controller"

 "github.com/gin-gonic/gin"
 swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
   _ "go-authentication/docs"
)

func AuthRouter(authController *controller.AuthController) *gin.Engine {
 service := gin.Default()
 router := service.Group("/auth")


 router.POST("/register", authController.Register)
 router.POST("/login", authController.Login)
 router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

 return service
}