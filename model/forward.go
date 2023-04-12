package model

import "gorm.io/gorm"

type Forward struct {
	gorm.Model
	Domain string `gorm:"unique"`
	Target string
	Port   string
}
