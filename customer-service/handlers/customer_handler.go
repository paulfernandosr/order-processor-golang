package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/models"
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

func (handler *CustomerHandler) GetCustomerById(context *gin.Context) {
	customerId := context.Param("id")

	customer, err := handler.repository.FindCustomerById(context, customerId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error geeting customer"})
		return
	}

	context.JSON(http.StatusOK, customer)
}

func (handler *CustomerHandler) CreateNewCustomer(context *gin.Context) {
	var customer models.Customer
	err := context.ShouldBindJSON(&customer)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.repository.CreateNewCustomer(context, &customer)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, customer)
}
