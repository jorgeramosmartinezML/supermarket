package validator

import (
	"fmt"
	"go-web/supermarket/internal/service/model"
	"regexp"
)

type productValidatorImpl struct {
	regexCodeValue *regexp.Regexp
}

func NewProductVaidatorImpl(regexCodeValue string) ProductValidator {
	productValidator := new(productValidatorImpl)

	regexCodeValueDefault := `^[A-Z]{3}-[0-9]{3}$`
	if regexCodeValue != "" {
		regexCodeValueDefault = regexCodeValue
	}
	productValidator.regexCodeValue = regexp.MustCompile(regexCodeValueDefault)
	return productValidator
}

func (validator *productValidatorImpl) Validate(p *model.AddProductInput) (err error) {
	if p.Name == "" {
		err = fmt.Errorf("%w: name is empty", ErrValidatorProductFieldRequired)
		return
	}
	if p.CodeValue == "" {
		err = fmt.Errorf("%w: code value is empty", ErrValidatorProductFieldRequired)
		return
	}

	if p.Quantity < 0 {
		err = fmt.Errorf("%w: quantity can't be negative", ErrValidatorProductFieldInvalid)
		return
	}
	if !validator.regexCodeValue.MatchString(p.CodeValue) {
		err = fmt.Errorf("%w: code value format is invalid", ErrValidatorProductFieldInvalid)
		return
	}
	/*if p.Expiration.Before(p.Expiration) {
		err = fmt.Errorf("%w: expiration date can't be before created date", ErrValidatorProductFieldInvalid)
		return
	}*/
	if p.Price < 0 {
		err = fmt.Errorf("%w: price can't be negative", ErrValidatorProductFieldInvalid)
		return
	}

	return
}
