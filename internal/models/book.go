package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model        // Id = auto inc
	Title      string `json:"Title"`
	PageCount  int    `json:"PageCount"`
}
