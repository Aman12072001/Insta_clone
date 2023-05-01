package server

import (
	_ "main/docs"
	"main/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {
	
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	server.engine.POST("/user-signup-phone",handler.UserSignupPhone)
	server.engine.POST("/verify-phone",handler.VerifyOtpHandler)
	

	
}