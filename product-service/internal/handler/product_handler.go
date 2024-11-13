package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/model"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (handler *ProductHandler) CreateProduct(context *gin.Context) {
	var product model.Product
	err := context.ShouldBindJSON(&product)

	if err != nil {
		context.Error(model.NewBadRequestError(err.Error()))
		return
	}

	err = handler.service.CreateProduct(context.Request.Context(), &product)

	if err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *ProductHandler) GetAllProducts(context *gin.Context) {
	products, err := handler.service.GetAllProducts(context.Request.Context())

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) GetProductsByIds(context *gin.Context) {
	ids := context.QueryArray("ids")

	if len(ids) == 0 {
		context.Error(model.NewBadRequestError("Invalid request"))
		return
	}

	products, err := handler.service.GetProductsByIds(context.Request.Context(), ids)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) GetProductById(context *gin.Context) {
	id := context.Param("id")

	product, err := handler.service.GetProductById(context.Request.Context(), id)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) UpdateProductById(context *gin.Context) {
	id := context.Param("id")

	var product model.Product
	err := context.ShouldBindJSON(&product)

	if err != nil {
		context.Error(model.NewBadRequestError(err.Error()))
		return
	}

	err = handler.service.UpdateProductById(context.Request.Context(), id, &product)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) DeleteProductById(context *gin.Context) {
	id := context.Param("id")

	err := handler.service.DeleteProductById(context.Request.Context(), id)

	if err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}
