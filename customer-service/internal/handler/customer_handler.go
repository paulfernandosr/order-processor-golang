package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/model"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (handler *CustomerHandler) CreateCustomer(context *gin.Context) {
	var customer model.Customer
	err := context.ShouldBindJSON(&customer)

	if err != nil {
		context.Error(model.NewBadRequestError(err.Error()))
		return
	}

	err = handler.service.CreateCustomer(context.Request.Context(), &customer)

	if err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *CustomerHandler) GetAllCustomers(context *gin.Context) {
	customers, err := handler.service.GetAllCustomers(context.Request.Context())

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, customers)
}

func (handler *CustomerHandler) GetCustomersByIds(context *gin.Context) {
	ids := context.QueryArray("ids")

	if len(ids) == 0 {
		context.Error(model.NewBadRequestError("Invalid request"))
		return
	}

	customers, err := handler.service.GetCustomersByIds(context.Request.Context(), ids)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, customers)
}

func (handler *CustomerHandler) GetCustomerById(context *gin.Context) {
	id := context.Param("id")

	customer, err := handler.service.GetCustomerById(context.Request.Context(), id)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func (handler *CustomerHandler) UpdateCustomerById(context *gin.Context) {
	id := context.Param("id")

	var customer model.Customer
	err := context.ShouldBindJSON(&customer)

	if err != nil {
		context.Error(model.NewBadRequestError(err.Error()))
		return
	}

	err = handler.service.UpdateCustomerById(context.Request.Context(), id, &customer)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func (handler *CustomerHandler) DeleteCustomerById(context *gin.Context) {
	id := context.Param("id")

	err := handler.service.DeleteCustomerById(context.Request.Context(), id)

	if err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}
