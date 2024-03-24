package implement

import (
	"go-web/supermarket/internal/entity"
	"go-web/supermarket/internal/service/model"
	"go-web/supermarket/internal/service/presenter"
)

type productPresenterImplement struct {
}

func NewCustomerPresenter() presenter.ProductPresenter {
	productPresenter := new(productPresenterImplement)
	return productPresenter
}

func (presenter *productPresenterImplement) Output(productEntity entity.Product) model.ProductOutput {
	return model.ProductOutput{
		Id:          productEntity.Id,
		Name:        productEntity.Name,
		Quantity:    productEntity.Quantity,
		CodeValue:   productEntity.CodeValue,
		IsPublished: productEntity.IsPublished,
		Expiration:  "",
		Price:       productEntity.Price,
	}
}
