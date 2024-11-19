package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/config"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/handler"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/middleware"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/repository"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/router"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/service"
)

func main() {
	config.LoadEnvironment()

	mongoClient := config.NewMongoClient()
	redisClient := config.NewRedisClient()
	kafkaProducer := config.NewKafkaProducer()
	kafkaConsumer := config.NewKafkaConsumer()

	orderWriter := repository.NewMongoOrderWriter(mongoClient)
	lockerManager := repository.NewRedisLockManager(redisClient)
	errorLogger := repository.NewRedisErrorLogger(redisClient)
	orderProducer := repository.NewKafkaOrderProducer(kafkaProducer)
	orderConsumer := repository.NewKafkaOrderConsumer(kafkaConsumer)

	customerReader := repository.NewHttpCustomerReader()
	productReader := repository.NewHttpProductReader()

	orderProcessor := service.NewOrderProcessorImpl(lockerManager, errorLogger, customerReader, productReader, orderWriter)
	orderHandler := handler.NewOrderHandler(orderProducer, orderConsumer, orderProcessor)
	orderRouter := router.NewOrderRouter(orderHandler)

	orderHandler.ProcessOrdersAsync()

	InitializeServer(orderRouter)
}

func InitializeServer(orderRouter *router.OrderRouter) {
	server := gin.Default()

	server.Use(middleware.ErrorHandler)

	orderRouter.SetUp(server)

	err := server.Run(":" + config.Props.ServerPort)

	if err != nil {
		log.Fatal(err)
	}
}
