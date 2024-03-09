package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

var (
	ErrFileNotFound = errors.New("file not found")
	ErrParseProduct = errors.New("error parsing product")
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func loadProducts() ([]Product, error) {
	data, err := os.ReadFile("products.json")
	if err != nil {
		return nil, ErrFileNotFound
	}
	var products []Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, ErrParseProduct
	}
	return products, nil
}

func getProductById(products []Product, id int) (Product, bool) {
	for _, product := range products {
		if product.Id == id {
			return product, true
		}
	}
	return Product{}, false
}

func getProductsWithPriceGreaterThan(products []Product, price float64) []Product {
	var productsWithPriceGreaterThan []Product
	for _, product := range products {
		if product.Price > price {
			productsWithPriceGreaterThan = append(productsWithPriceGreaterThan, product)
		}
	}
	return productsWithPriceGreaterThan
}

func getNextId(products []Product) int {
	var maxId int
	for _, product := range products {
		if product.Id > maxId {
			maxId = product.Id
		}
	}
	return maxId + 1
}

func isProductValid(product Product) bool {
	if product.Name == "" {
		return false
	}
	if product.Quantity <= 0 {
		return false
	}
	if product.CodeValue == "" {
		return false
	}
	if product.Price <= 0 {
		return false
	}
	if product.Expiration == "" {
		return false
	}
	return true
}

func existsCodeValue(products []Product, codeValue string) bool {
	for _, product := range products {
		if product.CodeValue == codeValue {
			return true
		}
	}
	return false
}

func isValidDate(date string) bool {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false
	}
	return true
}

func main() {
	products, err := loadProducts()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	})

	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product, ok := getProductById(products, id)
		if !ok {
			w.Header().Add("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("product not found"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	})

	router.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
		priceGtParam := r.URL.Query().Get("priceGt")

		priceGt, err := strconv.ParseFloat(priceGtParam, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		products := getProductsWithPriceGreaterThan(products, priceGt)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)

	})

	router.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !isProductValid(product) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if existsCodeValue(products, product.CodeValue) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !isValidDate(product.Expiration) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product.Id = getNextId(products)
		products = append(products, product)
		w.WriteHeader(http.StatusCreated)

	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		println("Error starting server: ", err.Error())
		return
	}
}
