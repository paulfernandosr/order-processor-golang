package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/handler"
)

type OrderRouter struct {
	orderHandler *handler.OrderHandler
}

func NewOrderRouter(orderHandler *handler.OrderHandler) *OrderRouter {
	return &OrderRouter{orderHandler}
}

func (router *OrderRouter) SetUp(server *gin.Engine) {
	server.POST("/api/v1/orders", router.orderHandler.SendOrder)
}
