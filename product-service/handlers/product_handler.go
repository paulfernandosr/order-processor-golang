package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/models"
	"github.com/paulfernandosr/order-processor-golang/product-service/repositories"
)

type ProductHandler struct {
	repository repositories.ProductRepository
}

func NewProductHandler(repository repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{repository}
}

func (handler *ProductHandler) CreateNewProduct(context *gin.Context) {
	var product models.Product
	err := context.ShouldBindJSON(&product)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.repository.CreateNewProduct(context, &product)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, product)
}
