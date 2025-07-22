package main

import (
	"log"
	"subscription-service/config"
	_ "subscription-service/docs"
	"subscription-service/handler"
	"subscription-service/repository"
	"subscription-service/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.InitDB()
	r := gin.Default()

	subRepo := repository.NewSubscriptionRepo(db)
	subService := service.NewSubscriptionService(subRepo)
	subHandler := handler.NewSubscriptionHandler(subService)

	subHandler.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
