package handler

import (
	"go-web/supermarket/internal/handler/response"
	"go-web/supermarket/internal/service"
	"go-web/supermarket/internal/service/model"
	"go-web/supermarket/platform/web/request"
	"net/http"
)

type HandlerProduct struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *HandlerProduct {
	handlerProduct := new(HandlerProduct)
	handlerProduct.service = service
	return handlerProduct
}

func (handler *HandlerProduct) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := new(response.ResponsePayload)

		var reqBody model.AddProductInput
		err := request.JSON(r, &reqBody)
		if err != nil {
			res.SetData(nil).SetMessage("Invalid request body").SetStatus(http.StatusBadRequest).SetSuccess(false).ToJson(w)
			return
		}

		input := model.AddProductInput{
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		handler.service.AddProduct(input)

	}
}
