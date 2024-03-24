package repository

import (
	"errors"
	"go-web/supermarket/internal/entity"
)

var (
	ErrRepositoryProductAlreadyExists = errors.New("repository: product already exists")
)

type ProductRepository interface {
	Save(product entity.Product) (entity.Product, error)
	FindByCodeValue(product entity.Product) (bool, error)
}
