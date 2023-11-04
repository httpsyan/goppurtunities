package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	remote bool
	Link string
	Salary int64
}