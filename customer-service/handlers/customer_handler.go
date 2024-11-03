package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/repository"
)

type CustomerHandler struct {
	repository repository.CustomerRepository
}

func NewCustomerHandler(repository repository.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{repository}
}

func (handler *CustomerHandler) GetAllCustomers(context *gin.Context) {
	customers, err := handler.repository.FindAllCustomers(context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error geeting customers"})
		return
	}

	context.JSON(http.StatusOK, customers)
}
