package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/models"
)

type CustomerRepository interface {
	FindAllCustomers(context *gin.Context) (*[]models.Customer, error)
	FindCustomerById(context *gin.Context, id string) (*models.Customer, error)
	CreateNewCustomer(context *gin.Context, customer *models.Customer) error
}
