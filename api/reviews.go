package api

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Comment string `gorm:"column:comment"`
	Movie   string `gorm:"column:movie"`
}
