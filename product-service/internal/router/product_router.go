package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/handler"
)

type ProductRouter struct {
	productHandler *handler.ProductHandler
}

func NewCustomerRouter(customerHandler *handler.ProductHandler) *ProductRouter {
	return &ProductRouter{customerHandler}
}

func (router *ProductRouter) SetUp(server *gin.Engine) {
	server.POST("/api/v1/products", router.productHandler.CreateProduct)
	server.GET("/api/v1/products", router.productHandler.GetAllProducts)
	server.GET("/api/v1/products/ids", router.productHandler.GetProductsByIds)
	server.GET("/api/v1/products/:id", router.productHandler.GetProductById)
	server.PUT("/api/v1/products/:id", router.productHandler.UpdateProductById)
	server.DELETE("/api/v1/products/:id", router.productHandler.DeleteProductById)
}
