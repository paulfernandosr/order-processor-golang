package service

import (
	"context"

	"github.com/paulfernandosr/order-processor-golang/product-service/internal/model"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/repository"
)

type ProductService interface {
	CreateProduct(context.Context, *model.Product) error
	GetAllProducts(context.Context) ([]model.Product, error)
	GetProductsByIds(context.Context, []string) ([]model.Product, error)
	GetProductById(context.Context, string) (*model.Product, error)
	UpdateProductById(context.Context, string, *model.Product) error
	DeleteProductById(context.Context, string) error
}

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{repository}
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, product *model.Product) error {
	return service.repository.InsertOneProduct(ctx, product)
}

func (service *ProductServiceImpl) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	return service.repository.FindAllProducts(ctx)
}

func (service *ProductServiceImpl) GetProductsByIds(ctx context.Context, productIds []string) ([]model.Product, error) {
	return service.repository.FindProductsByIds(ctx, productIds)
}

func (service *ProductServiceImpl) GetProductById(ctx context.Context, productId string) (*model.Product, error) {
	return service.repository.FindOneProductById(ctx, productId)
}

func (service *ProductServiceImpl) UpdateProductById(ctx context.Context, productId string, product *model.Product) error {
	return service.repository.UpdateOneProductById(ctx, productId, product)
}

func (service *ProductServiceImpl) DeleteProductById(ctx context.Context, productId string) error {
	return service.repository.DeleteOneProductById(ctx, productId)
}
