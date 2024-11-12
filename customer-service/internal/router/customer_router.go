package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/handler"
)

type CustomerRouter struct {
	customerHandler *handler.CustomerHandler
}

func NewCustomerRouter(customerHandler *handler.CustomerHandler) *CustomerRouter {
	return &CustomerRouter{customerHandler}
}

func (router *CustomerRouter) SetUp(server *gin.Engine) {
	server.POST("/api/v1/customers", router.customerHandler.CreateCustomer)
	server.GET("/api/v1/customers", router.customerHandler.GetAllCustomers)
	server.GET("/api/v1/customers/ids", router.customerHandler.GetCustomersByIds)
	server.GET("/api/v1/customers/:id", router.customerHandler.GetCustomerById)
	server.PUT("/api/v1/customers/:id", router.customerHandler.UpdateCustomerById)
	server.DELETE("/api/v1/customers/:id", router.customerHandler.DeleteCustomerById)
}
