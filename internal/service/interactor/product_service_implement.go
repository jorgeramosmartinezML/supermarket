package interactor

import (
	"go-web/supermarket/internal/entity"
	"go-web/supermarket/internal/repository"
	"go-web/supermarket/internal/service"
	"go-web/supermarket/internal/service/model"
	"go-web/supermarket/internal/service/presenter"
)

type productInteractor struct {
	productRepository repository.ProductRepository
	productPresenter  presenter.ProductPresenter
}

func NewProductInteractor(productRepository repository.ProductRepository, productPresenter presenter.ProductPresenter) service.ProductService {
	productInteractor := new(productInteractor)

	productInteractor.productRepository = productRepository
	productInteractor.productPresenter = productPresenter
	return productInteractor
}

func (interactor *productInteractor) AddProduct(input model.AddProductInput) (model.ProductOutput, error) {
	productEntity := entity.Product{Name: input.Name, Quantity: input.Quantity, CodeValue: input.CodeValue, IsPublished: input.IsPublished, Price: input.Price}

	existsProduct, err := interactor.productRepository.FindByCodeValue(productEntity)
	if existsProduct {
		if err == nil {
			return model.ProductOutput{}, service.ErrServiceProductAlreadyExists
		}
		return model.ProductOutput{}, err
	}
	newProductEntity, err := interactor.productRepository.Save(productEntity)
	output := interactor.productPresenter.Output(newProductEntity)
	return output, err
}
