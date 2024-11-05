package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/config"
	"github.com/paulfernandosr/order-processor-golang/product-service/handlers"
	"github.com/paulfernandosr/order-processor-golang/product-service/repositories"
	"github.com/paulfernandosr/order-processor-golang/product-service/routers"
)

func main() {
	config.LoadEnviroment()

	mongoClient := config.NewMongoClient()
	productRepository := repositories.NewMongoProductRepository(mongoClient, "order-processor", "products")
	productHandler := handlers.NewProductHandler(productRepository)
	productRouter := routers.NewProductRouter(productHandler)

	InitializeServer(productRouter)
}

func InitializeServer(productRouter *routers.ProductRouter) {
	server := gin.Default()

	productRouter.RegisterRoutes(server)

	err := server.Run(":" + config.EnvironmentProps.ServerPort)

	if err != nil {
		log.Fatal(err)
	}
}
