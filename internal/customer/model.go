package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"not null"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	Address     string     `json:"address"`
	SIPA        string     `json:"sipa"`
	ExpiredDate *time.Time `json:"expired_date,omitempty"`
	CreatedAt   int64      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64      `json:"updated_at" gorm:"autoUpdateTime"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Customer{})
}
