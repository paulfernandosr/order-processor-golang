package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/handlers"
)

type CustomerRouter struct {
	customerHandler *handlers.CustomerHandler
}

func NewCustomerRouter(customerHandler *handlers.CustomerHandler) *CustomerRouter {
	return &CustomerRouter{customerHandler}
}

func (router *CustomerRouter) RegisterRoutes(server *gin.Engine) {
	server.GET("/customers", router.customerHandler.GetAllCustomers)
	server.GET("/customers/:id", router.customerHandler.GetCustomerById)
	server.POST("/customers", router.customerHandler.CreateNewCustomer)
}
