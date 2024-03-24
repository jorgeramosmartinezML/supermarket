package validator

import (
	"errors"
	"go-web/supermarket/internal/service/model"
)

var (
	ErrValidatorProductFieldRequired = errors.New("validator: product field required")
	ErrValidatorProductFieldInvalid  = errors.New("validator: product field invalid")
)

type ProductValidator interface {
	Validate(product *model.AddProductInput) error
}
