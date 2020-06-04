package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name string
	Age  sql.NullInt64
}