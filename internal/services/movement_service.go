package services

import (
	"fmt"
	"stoktakip/internal/database"
	"stoktakip/internal/models"
	"time"
)

// MovementDTO is the data transfer object for movements
type MovementDTO struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Type      string    `json:"type"` // "IN" or "OUT"
	Quantity  int       `json:"quantity"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

// MovementStats holds statistics about movements
type MovementStats struct {
	TotalIn       int   `json:"total_in"`
	TotalOut      int   `json:"total_out"`
	TodayIn       int   `json:"today_in"`
	TodayOut      int   `json:"today_out"`
	MovementCount int64 `json:"movement_count"`
}

// MovementService handles stock movement operations
type MovementService struct {
	dbManager *database.ConnectionManager
}

// NewMovementService creates a new movement service
func NewMovementService(dbManager *database.ConnectionManager) *MovementService {
	return &MovementService{
		dbManager: dbManager,
	}
}

// Helper function to convert model to DTO
func (s *MovementService) toDTO(movement *models.StockMovement) MovementDTO {
	return MovementDTO{
		ID:        movement.ID,
		ProductID: movement.ProductID,
		Type:      string(movement.Type),
		Quantity:  movement.Quantity,
		Note:      movement.Note,
		CreatedAt: movement.CreatedAt,
	}
}

// GetAll returns all movements as DTOs
func (s *MovementService) GetAll() ([]MovementDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var movements []models.StockMovement
	if err := db.Order("created_at DESC").Find(&movements).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch movements: %w", err)
	}

	dtos := make([]MovementDTO, len(movements))
	for i, movement := range movements {
		dtos[i] = s.toDTO(&movement)
	}

	return dtos, nil
}

// GetByID returns a movement by ID as DTO
func (s *MovementService) GetByID(id uint) (*MovementDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var movement models.StockMovement
	if err := db.First(&movement, id).Error; err != nil {
		return nil, fmt.Errorf("movement not found: %w", err)
	}

	dto := s.toDTO(&movement)
	return &dto, nil
}

// Create creates a new movement from DTO
func (s *MovementService) Create(dto MovementDTO) (*MovementDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	// Validate movement type
	if dto.Type != "IN" && dto.Type != "OUT" {
		return nil, fmt.Errorf("invalid movement type: %s", dto.Type)
	}

	// Check if product exists
	var product models.Product
	if err := db.First(&product, dto.ProductID).Error; err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	// For OUT movements, check if sufficient stock
	if dto.Type == "OUT" && product.CurrentStock < dto.Quantity {
		return nil, fmt.Errorf("insufficient stock: available %d, requested %d", product.CurrentStock, dto.Quantity)
	}

	// Create movement
	movement := &models.StockMovement{
		ProductID: dto.ProductID,
		Type:      models.MovementType(dto.Type),
		Quantity:  dto.Quantity,
		Date:      time.Now(),
		Note:      dto.Note,
	}

	if err := db.Create(movement).Error; err != nil {
		return nil, fmt.Errorf("failed to create movement: %w", err)
	}

	// Update product stock
	if dto.Type == "IN" {
		product.CurrentStock += dto.Quantity
	} else {
		product.CurrentStock -= dto.Quantity
	}

	if err := db.Save(&product).Error; err != nil {
		// Rollback movement if stock update fails
		db.Delete(movement)
		return nil, fmt.Errorf("failed to update product stock: %w", err)
	}

	resultDTO := s.toDTO(movement)
	return &resultDTO, nil
}

// Delete deletes a movement by ID
func (s *MovementService) Delete(id uint) error {
	db := s.dbManager.GetDB()
	if db == nil {
		return fmt.Errorf("no database connection")
	}

	// Get movement first
	var movement models.StockMovement
	if err := db.First(&movement, id).Error; err != nil {
		return fmt.Errorf("movement not found: %w", err)
	}

	// Get product
	var product models.Product
	if err := db.First(&product, movement.ProductID).Error; err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	// Reverse the stock change
	if movement.Type == models.MovementTypeIn {
		product.CurrentStock -= movement.Quantity
	} else {
		product.CurrentStock += movement.Quantity
	}

	// Prevent negative stock
	if product.CurrentStock < 0 {
		return fmt.Errorf("cannot delete movement: would result in negative stock")
	}

	// Update product stock
	if err := db.Save(&product).Error; err != nil {
		return fmt.Errorf("failed to update product stock: %w", err)
	}

	// Delete movement
	if err := db.Delete(&movement).Error; err != nil {
		return fmt.Errorf("failed to delete movement: %w", err)
	}

	return nil
}

// GetByProduct returns movements for a specific product
func (s *MovementService) GetByProduct(productID uint) ([]MovementDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var movements []models.StockMovement
	if err := db.Where("product_id = ?", productID).Order("created_at DESC").Find(&movements).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch movements: %w", err)
	}

	dtos := make([]MovementDTO, len(movements))
	for i, movement := range movements {
		dtos[i] = s.toDTO(&movement)
	}

	return dtos, nil
}

// GetStats returns movement statistics
func (s *MovementService) GetStats() (*MovementStats, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	stats := &MovementStats{}

	// Total movement count
	if err := db.Model(&models.StockMovement{}).Count(&stats.MovementCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count movements: %w", err)
	}

	// Total IN
	var totalIn int64
	if err := db.Model(&models.StockMovement{}).Where("type = ?", "IN").Select("COALESCE(SUM(quantity), 0)").Scan(&totalIn).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate total IN: %w", err)
	}
	stats.TotalIn = int(totalIn)

	// Total OUT
	var totalOut int64
	if err := db.Model(&models.StockMovement{}).Where("type = ?", "OUT").Select("COALESCE(SUM(quantity), 0)").Scan(&totalOut).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate total OUT: %w", err)
	}
	stats.TotalOut = int(totalOut)

	// Today's movements
	today := time.Now().Format("2006-01-02")
	var todayIn int64
	if err := db.Model(&models.StockMovement{}).Where("type = ? AND DATE(date) = ?", "IN", today).Select("COALESCE(SUM(quantity), 0)").Scan(&todayIn).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate today's IN: %w", err)
	}
	stats.TodayIn = int(todayIn)

	var todayOut int64
	if err := db.Model(&models.StockMovement{}).Where("type = ? AND DATE(date) = ?", "OUT", today).Select("COALESCE(SUM(quantity), 0)").Scan(&todayOut).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate today's OUT: %w", err)
	}
	stats.TodayOut = int(todayOut)

	return stats, nil
}
