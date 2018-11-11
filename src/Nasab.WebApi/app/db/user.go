package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User contains information for current User
type User struct {
	gorm.Model

	Username string `gorm:"not null" json:"username"`

	Password string `gorm:"not null" json:"password"`

	TZ *time.Location
}
