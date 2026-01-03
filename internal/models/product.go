package models

import (
	"time"
)

// Product represents a product in the inventory
type Product struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Code          string    `gorm:"size:50;uniqueIndex;not null" json:"code"`
	Name          string    `gorm:"size:200;not null;index" json:"name"`
	CategoryID    uint      `gorm:"not null;index" json:"category_id"`
	Unit          string    `gorm:"size:20;not null" json:"unit"` // adet, kg, litre, etc.
	CriticalLimit int       `gorm:"default:0" json:"critical_limit"`
	Price         float64   `gorm:"type:decimal(10,2);default:0" json:"price"`
	CurrentStock  int       `gorm:"default:0" json:"current_stock"` // Computed from movements
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relations
	Category  Category        `gorm:"foreignKey:CategoryID" json:"category"`
	Movements []StockMovement `gorm:"foreignKey:ProductID" json:"-"`
}

// TableName specifies the table name for Product model
func (Product) TableName() string {
	return "products"
}

// ProductWithStock is a computed view that includes current stock information
type ProductWithStock struct {
	Product
	CurrentStock int     `json:"current_stock"`
	StockValue   float64 `json:"stock_value"` // CurrentStock * Price
	IsLowStock   bool    `json:"is_low_stock"`
}
