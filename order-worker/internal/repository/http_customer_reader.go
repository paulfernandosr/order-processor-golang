package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/config"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/model"
)

type CustomerReader interface {
	GetById(id string) (*model.Customer, error)
}

type HttpCustomerReader struct {
	httpClient *http.Client
	baseUrl    *url.URL
}

func NewHttpCustomerReader() CustomerReader {
	parsedUrl, err := url.Parse(config.Props.CustomerServiceBaseUrl)

	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}

	return &HttpCustomerReader{&http.Client{}, parsedUrl}
}

func (customerReader *HttpCustomerReader) GetById(id string) (*model.Customer, error) {
	finalUrl := customerReader.baseUrl.JoinPath(id)

	response, err := customerReader.httpClient.Get(finalUrl.String())

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var customer model.Customer

	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
