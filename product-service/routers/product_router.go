package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/handlers"
)

type ProductRouter struct {
	productHandler *handlers.ProductHandler
}

func NewProductRouter(productHandler *handlers.ProductHandler) *ProductRouter {
	return &ProductRouter{productHandler}
}

func (router *ProductRouter) RegisterRoutes(server *gin.Engine) {
	server.GET("/products", router.productHandler.CreateNewProduct)
	server.GET("/products/:id", router.productHandler.CreateNewProduct)
	server.POST("/products", router.productHandler.CreateNewProduct)
}
