package services

import (
	"fmt"
	"stoktakip/internal/database"
	"stoktakip/internal/models"
	"time"
)

// ProductDTO is the data transfer object for products
type ProductDTO struct {
	ID            uint      `json:"id"`
	Code          string    `json:"code"`
	Name          string    `json:"name"`
	CategoryID    uint      `json:"category_id"`
	Unit          string    `json:"unit"`
	CriticalLimit int       `json:"critical_limit"`
	Price         float64   `json:"price"`
	CurrentStock  int       `json:"current_stock"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ProductService handles product-related operations
type ProductService struct {
	dbManager *database.ConnectionManager
}

// NewProductService creates a new product service
func NewProductService(dbManager *database.ConnectionManager) *ProductService {
	return &ProductService{
		dbManager: dbManager,
	}
}

// Helper function to convert model to DTO
func (s *ProductService) toDTO(product *models.Product) ProductDTO {
	return ProductDTO{
		ID:            product.ID,
		Code:          product.Code,
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Unit:          product.Unit,
		CriticalLimit: product.CriticalLimit,
		Price:         product.Price,
		CurrentStock:  product.CurrentStock,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

// GetAll returns all products as DTOs
func (s *ProductService) GetAll() ([]ProductDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var products []models.Product
	if err := db.Order("name ASC").Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	dtos := make([]ProductDTO, len(products))
	for i, product := range products {
		dtos[i] = s.toDTO(&product)
	}

	return dtos, nil
}

// GetByID returns a product by ID as DTO
func (s *ProductService) GetByID(id uint) (*ProductDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	dto := s.toDTO(&product)
	return &dto, nil
}

// Create creates a new product from DTO
func (s *ProductService) Create(dto ProductDTO) (*ProductDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	// Validate
	if dto.Code == "" || dto.Name == "" {
		return nil, fmt.Errorf("code and name are required")
	}

	// Check if product with same code already exists
	var existing models.Product
	if err := db.Where("code = ?", dto.Code).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("product with code '%s' already exists", dto.Code)
	}

	product := &models.Product{
		Code:          dto.Code,
		Name:          dto.Name,
		CategoryID:    dto.CategoryID,
		Unit:          dto.Unit,
		CriticalLimit: dto.CriticalLimit,
		Price:         dto.Price,
		CurrentStock:  0, // Initial stock is 0
	}

	if err := db.Create(product).Error; err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	resultDTO := s.toDTO(product)
	return &resultDTO, nil
}

// Update updates a product from DTO
func (s *ProductService) Update(id uint, dto ProductDTO) (*ProductDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	// Validate
	if dto.Code == "" || dto.Name == "" {
		return nil, fmt.Errorf("code and name are required")
	}

	// Check if another product with same code exists
	var existing models.Product
	if err := db.Where("code = ? AND id != ?", dto.Code, id).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("product with code '%s' already exists", dto.Code)
	}

	// Update fields (but not current_stock, that's managed by movements)
	product.Code = dto.Code
	product.Name = dto.Name
	product.CategoryID = dto.CategoryID
	product.Unit = dto.Unit
	product.CriticalLimit = dto.CriticalLimit
	product.Price = dto.Price

	if err := db.Save(&product).Error; err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	resultDTO := s.toDTO(&product)
	return &resultDTO, nil
}

// Delete deletes a product by ID
func (s *ProductService) Delete(id uint) error {
	db := s.dbManager.GetDB()
	if db == nil {
		return fmt.Errorf("no database connection")
	}

	// Check if product has movements
	var movementCount int64
	if err := db.Model(&models.StockMovement{}).Where("product_id = ?", id).Count(&movementCount).Error; err != nil {
		return fmt.Errorf("failed to check movements: %w", err)
	}

	if movementCount > 0 {
		return fmt.Errorf("cannot delete product with %d movements", movementCount)
	}

	// Delete product
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// GetLowStock returns products with low stock
func (s *ProductService) GetLowStock() ([]ProductDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var products []models.Product
	if err := db.Where("current_stock <= critical_limit AND current_stock > 0").Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch low stock products: %w", err)
	}

	dtos := make([]ProductDTO, len(products))
	for i, product := range products {
		dtos[i] = s.toDTO(&product)
	}

	return dtos, nil
}
