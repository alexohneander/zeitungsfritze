package model

import "gorm.io/gorm"

type Forward struct {
	gorm.Model
	Domain string `gorm:"uniqueIndex;not null"`
	Target string
	Port   string
}
