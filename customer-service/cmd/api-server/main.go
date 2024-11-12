package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/config"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/handler"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/middleware"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/repository"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/router"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/service"
)

func main() {
	config.LoadEnv()

	mongoClient := config.NewMongoClient()
	customerRepository := repository.NewMongoCustomerRepository(mongoClient)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)
	customerRouter := router.NewCustomerRouter(customerHandler)

	InitializeServer(customerRouter)
}

func InitializeServer(customerRouter *router.CustomerRouter) {
	server := gin.Default()

	server.Use(middleware.ErrorHandler)

	customerRouter.SetUp(server)

	err := server.Run(":" + config.EnvProps.ServerPort)

	if err != nil {
		log.Fatal(err)
	}
}
