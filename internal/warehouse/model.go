package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Location  string         `json:"location" gorm:"size:255"`
	Manager   string         `json:"manager" gorm:"size:100"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Warehouse{})
}
