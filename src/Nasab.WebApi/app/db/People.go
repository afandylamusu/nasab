package db

import (
	"github.com/jinzhu/gorm"
)

// People contains information for current People connection
type People struct {
	gorm.Model

	Name string
}
