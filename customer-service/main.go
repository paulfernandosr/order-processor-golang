package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/config"
	"github.com/paulfernandosr/order-processor-golang/customer-service/handlers"
	"github.com/paulfernandosr/order-processor-golang/customer-service/repository"
	"github.com/paulfernandosr/order-processor-golang/customer-service/routes"
)

func main() {
	mongoClient := config.NewMongoClient()
	customerRepository := repository.NewCustomerMongoRepository(mongoClient, "customerdb", "customers")
	customerHandler := handlers.NewCustomerHandler(customerRepository)
	customerRouter := routes.NewCustomerRouter(customerHandler)

	server := gin.Default()
	customerRouter.RegisterRoutes(server)

	server.Run(":8080")
}
