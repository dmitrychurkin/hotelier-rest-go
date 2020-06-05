package model

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name string
	Age  uint8
}
