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

type ProductReader interface {
	GetByIds(ids []string) ([]model.Product, error)
}

type HttpProductReader struct {
	httpClient *http.Client
	baseUrl    *url.URL
}

func NewHttpProductReader() ProductReader {
	parsedUrl, err := url.Parse(config.Props.ProductServiceBaseUrl)

	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}

	return &HttpProductReader{&http.Client{}, parsedUrl}
}

func (productReader *HttpProductReader) GetByIds(ids []string) ([]model.Product, error) {
	finalUrl := productReader.baseUrl.JoinPath("ids")

	queryParams := url.Values{}

	for _, id := range ids {
		queryParams.Add("ids", id)
	}

	finalUrl.RawQuery = queryParams.Encode()

	response, err := productReader.httpClient.Get(finalUrl.String())

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	products := make([]model.Product, 0)

	if err := json.NewDecoder(response.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}
