package localMap

import (
	"go-web/supermarket/internal/entity"
	"go-web/supermarket/internal/repository"
	"time"
)

type productRepositoryMap struct {
	db         map[int]entity.Product
	lastId     int
	formatDate string
}

func NewProductRepository(db map[int]entity.Product, lastId int, formatDate string) repository.ProductRepository {
	productRepository := new(productRepositoryMap)
	productRepository.db = db
	productRepository.lastId = lastId
	defaultFormatDate := time.DateOnly
	if formatDate != "" {
		productRepository.formatDate = defaultFormatDate
	}
	return productRepository
}

func (repository *productRepositoryMap) Save(product entity.Product) (entity.Product, error) {
	repository.lastId++
	product.Id = repository.lastId
	repository.db[repository.lastId] = product

	return product, nil
}

func (repository *productRepositoryMap) FindByCodeValue(product entity.Product) (bool, error) {
	for _, value := range repository.db {
		if value.CodeValue == product.CodeValue {
			return true, nil
		}
	}
	return false, nil
}
