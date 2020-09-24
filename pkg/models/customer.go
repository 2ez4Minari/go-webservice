package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model

	CustomerName string
	CustomerAge int
	CustomerAddress string
	CustomerEmail string
	IsActive bool
}
