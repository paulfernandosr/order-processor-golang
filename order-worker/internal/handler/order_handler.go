package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/model"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/repository"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/service"
)

type OrderHandler struct {
	orderProducer  repository.OrderProducer
	orderConsumer  repository.OrderConsumer
	orderProcessor service.OrderProcessor
}

func NewOrderHandler(orderProducer repository.OrderProducer, orderConsumer repository.OrderConsumer, orderProcessor service.OrderProcessor) *OrderHandler {
	return &OrderHandler{orderProducer, orderConsumer, orderProcessor}
}

func (orderHandler *OrderHandler) ProcessOrdersAsync() {
	orderChan := make(chan *model.Order)

	go orderHandler.orderConsumer.Consume(orderChan)

	go func() {
		for order := range orderChan {
			go orderHandler.orderProcessor.Process(order)
		}
	}()
}

func (orderHandler *OrderHandler) SendOrder(context *gin.Context) {
	var order model.Order

	err := context.ShouldBindJSON(&order)

	if err != nil {
		context.Error(model.NewBadRequestError(err.Error()))
		return
	}

	orderHandler.orderProducer.Produce(&order)

	context.Status(http.StatusNoContent)
}
