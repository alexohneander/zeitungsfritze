package model

import "gorm.io/gorm"

type Forward struct {
	gorm.Model
	Domain string
	Target string
	Port   string
}
