package presenter

import (
	"go-web/supermarket/internal/entity"
	"go-web/supermarket/internal/service/model"
)

type ProductPresenter interface {
	Output(productEntity entity.Product) model.ProductOutput
}
