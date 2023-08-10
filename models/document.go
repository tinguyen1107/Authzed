package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	ParentId *uint  `gorm:"uniqueIndex:udx_name"`
	Content  string // Will be encoded with base64
	Name     string `gorm:"uniqueIndex:udx_name"`
	// Composite unique index for ParentId and Name
}
