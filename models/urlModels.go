package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	permLink  string
	shortLink string
}
