package models

import (
	"time"
)

// Category represents a product category
type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null;index" json:"name"`
	Color     string    `gorm:"size:7;default:#6B7280" json:"color"` // HEX color
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Products []Product `gorm:"foreignKey:CategoryID" json:"-"`
}

// TableName specifies the table name for Category model
func (Category) TableName() string {
	return "categories"
}
