package entity

import "time"

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}
