package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/config"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/handler"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/middleware"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/repository"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/router"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/service"
)

func main() {
	config.LoadEnvironment()

	mongoClient := config.NewMongoClient()
	productRepository := repository.NewMongoProductRepository(mongoClient)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	productRouter := router.NewCustomerRouter(productHandler)

	InitializeServer(productRouter)
}

func InitializeServer(productRouter *router.ProductRouter) {
	server := gin.Default()

	server.Use(middleware.ErrorHandler)

	productRouter.SetUp(server)

	err := server.Run(":" + config.Props.ServerPort)

	if err != nil {
		log.Fatal(err)
	}
}
