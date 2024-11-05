package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/config"
	"github.com/paulfernandosr/order-processor-golang/customer-service/handlers"
	"github.com/paulfernandosr/order-processor-golang/customer-service/repository"
	"github.com/paulfernandosr/order-processor-golang/customer-service/routes"
)

func main() {
	config.LoadEnvironment()

	mongoClient := config.NewMongoClient()
	customerRepository := repository.NewMongoCustomerRepository(mongoClient, "customerdb", "customers")
	customerHandler := handlers.NewCustomerHandler(customerRepository)
	customerRouter := routes.NewCustomerRouter(customerHandler)

	InitializeServer(customerRouter)
}

func InitializeServer(customerRouter *routes.CustomerRouter) {
	server := gin.Default()

	customerRouter.RegisterRoutes(server)

	err := server.Run(":" + os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
