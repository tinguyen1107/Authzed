package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	OwnerId    *uint
	Name       string     `gorm:"uniqueIndex:udx_name"`
	ParentId   *uint      `gorm:"uniqueIndex:udx_name"`
	SubFolders []Folder   `gorm:"foreignkey:ParentId"`
	Documents  []Document `gorm:"foreignkey:ParentId"`
}
