package service

import (
	"errors"
	"go-web/supermarket/internal/service/model"
)

var (
	ErrServiceProductAlreadyExists = errors.New("service: product already exists")
)

type ProductService interface {
	AddProduct(input model.AddProductInput) (model.ProductOutput, error)
}
