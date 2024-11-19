package service

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/uuid"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/helper"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/model"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/repository"
)

type OrderProcessor interface {
	Process(order *model.Order)
}

type OrderProcessorImpl struct {
	lockManager    repository.LockManager
	errorLogger    repository.ErrorLogger
	customerReader repository.CustomerReader
	productReader  repository.ProductReader
	orderWriter    repository.OrderWriter
}

func NewOrderProcessorImpl(lockManager repository.LockManager, errorLogger repository.ErrorLogger, customerReader repository.CustomerReader, productReader repository.ProductReader, orderWriter repository.OrderWriter) OrderProcessor {
	return &OrderProcessorImpl{lockManager, errorLogger, customerReader, productReader, orderWriter}
}

func (orderProcessor *OrderProcessorImpl) Process(order *model.Order) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lockKey := helper.GenerateOrderLockKey(order.OrderId)
	errKey := helper.GenerateOrderErrorKey(order.OrderId)
	clientId := uuid.NewString()

	isAcquired, err := orderProcessor.lockManager.AcquireLock(ctx, lockKey, clientId, 5*time.Second)

	if err != nil {
		orderProcessor.errorLogger.LogError(ctx, errKey, err)
		return
	}

	if !isAcquired {
		orderProcessor.errorLogger.LogError(ctx, errKey, errors.New("the lock could not be acquired"))
		return
	}

	customerChan := make(chan *model.Customer)
	productsChan := make(chan []model.Product)
	errChan := make(chan error)

	var dataWg sync.WaitGroup

	dataWg.Add(2)

	go func() {
		defer dataWg.Done()

		var customer *model.Customer

		getCustomerById := func() error {
			var err error
			customer, err = orderProcessor.customerReader.GetById(order.CustomerId)
			return err
		}

		expBackOff := backoff.NewExponentialBackOff()
		expBackOff.MaxElapsedTime = 5 * time.Second

		err := backoff.Retry(getCustomerById, expBackOff)

		if err != nil {
			errChan <- err
			return
		}

		customerChan <- customer
	}()

	go func() {
		defer dataWg.Done()

		var products []model.Product

		getProductsByIds := func() error {
			var err error
			products, err = orderProcessor.productReader.GetByIds(order.ProductIds)
			return err
		}

		expBackOff := backoff.NewExponentialBackOff()
		expBackOff.MaxElapsedTime = 5 * time.Second

		err := backoff.Retry(getProductsByIds, expBackOff)

		if err != nil {
			errChan <- err
			return
		}

		productsChan <- products
	}()

	go func() {
		dataWg.Wait()
		close(customerChan)
		close(productsChan)
		close(errChan)
	}()

	var customer *model.Customer
	var products []model.Product

	for i := 0; i < 2; i++ {
		select {
		case customer = <-customerChan:
			log.Printf("Customer data: %+v\n", customer)
		case products = <-productsChan:
			log.Printf("Products data: %+v\n", products)
		case err := <-errChan:
			orderProcessor.errorLogger.LogError(ctx, errKey, err)
			return
		}
	}

	if !customer.IsActive {
		orderProcessor.errorLogger.LogError(ctx, errKey, errors.New("inactive customer"))
		return
	}

	if len(order.ProductIds) != len(products) {
		orderProcessor.errorLogger.LogError(ctx, errKey, errors.New("invalid products"))
		return
	}

	order.Products = products
	order.CustomerName = customer.Name

	err = orderProcessor.orderWriter.Save(order)

	if err != nil {
		orderProcessor.errorLogger.LogError(ctx, errKey, err)
		return
	}

	isReleased, err := orderProcessor.lockManager.ReleaseLock(ctx, lockKey, clientId)

	if err != nil {
		orderProcessor.errorLogger.LogError(ctx, errKey, err)
		return
	}

	if !isReleased {
		orderProcessor.errorLogger.LogError(ctx, errKey, errors.New("the lock could not be released"))
	}

	log.Printf("Order processed successfully: %+v\n", order)
}
