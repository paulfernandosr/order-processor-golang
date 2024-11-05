package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/models"
)

type ProductRepository interface {
	FindAllProducts(context *gin.Context) ([]models.Product, error)
	FindProductById(context *gin.Context, id string) (*models.Product, error)
	CreateNewProduct(context *gin.Context, product *models.Product) error
}
