package models

import (
	"time"
)

// MovementType represents the type of stock movement
type MovementType string

const (
	MovementTypeIn  MovementType = "IN"
	MovementTypeOut MovementType = "OUT"
)

// StockMovement represents a stock movement (in or out)
type StockMovement struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	ProductID uint         `gorm:"not null;index" json:"product_id"`
	Type      MovementType `gorm:"type:varchar(3);not null;index" json:"type"`
	Quantity  int          `gorm:"not null" json:"quantity"` // Always positive
	Date      time.Time    `gorm:"not null;index" json:"date"`
	Note      string       `gorm:"type:text" json:"note"`
	CreatedAt time.Time    `json:"created_at"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// TableName specifies the table name for StockMovement model
func (StockMovement) TableName() string {
	return "stock_movements"
}

// IsValid checks if the movement type is valid
func (m MovementType) IsValid() bool {
	return m == MovementTypeIn || m == MovementTypeOut
}
