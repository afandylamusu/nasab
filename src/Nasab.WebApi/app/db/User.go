package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User contains information for current User
type User struct {
	gorm.Model

	Username string `gorm:"not null"`

	Password string `gorm:"not null"`

	TZ *time.Location
}
