package model

type AddProductInput struct {
	Name        string `json:"name" validate:"required,gte=1"`
	Quantity    int    `json:"quantity" validate:"required,min=0"`
	CodeValue   string `json:"code_value" validate:"required,gte=1,regexp=^[A-Z]{3}-[0-9]{3}$"`
	IsPublished bool
	Expiration  string `json:"expiration" validate:"required, gte=1"`
	//Expiration string  `json:"expiration" validate:"required,regexp=^(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[0-2])/\\d{4}$"`
	Price float64 `json:"price" validate:"required,min=0"`
	//Expiration string  `json:"expiration" validate:"required,regexp=^(0[1-9]|[12][0-9]|3[01])[- /.](0[1-9]|1[012])[- /.](19|20)\d\d$"`
}

type ProductOutput struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
