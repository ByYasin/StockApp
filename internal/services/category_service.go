package services

import (
	"fmt"
	"stoktakip/internal/database"
	"stoktakip/internal/models"
	"time"
)

// CategoryDTO is the data transfer object for categories
type CategoryDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryService handles category-related operations
type CategoryService struct {
	dbManager *database.ConnectionManager
}

// NewCategoryService creates a new category service
func NewCategoryService(dbManager *database.ConnectionManager) *CategoryService {
	return &CategoryService{
		dbManager: dbManager,
	}
}

// GetAllCategories returns all categories
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var categories []models.Category
	if err := db.Order("name ASC").Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	return categories, nil
}

// GetCategoryByID returns a category by its ID
func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	return &category, nil
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(name, color string) (*models.Category, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	// Validate
	if name == "" {
		return nil, fmt.Errorf("category name cannot be empty")
	}

	// Check if category with same name already exists
	var existing models.Category
	if err := db.Where("name = ?", name).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("category with name '%s' already exists", name)
	}

	// Set default color if not provided
	if color == "" {
		color = "#6B7280"
	}

	category := &models.Category{
		Name:  name,
		Color: color,
	}

	if err := db.Create(category).Error; err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return category, nil
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(id uint, name, color string) error {
	db := s.dbManager.GetDB()
	if db == nil {
		return fmt.Errorf("no database connection")
	}

	// Find category
	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	// Validate
	if name == "" {
		return fmt.Errorf("category name cannot be empty")
	}

	// Check if another category with same name exists
	var existing models.Category
	if err := db.Where("name = ? AND id != ?", name, id).First(&existing).Error; err == nil {
		return fmt.Errorf("category with name '%s' already exists", name)
	}

	// Update fields
	category.Name = name
	if color != "" {
		category.Color = color
	}

	if err := db.Save(&category).Error; err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}

	return nil
}

// DeleteCategory deletes a category
func (s *CategoryService) DeleteCategory(id uint) error {
	db := s.dbManager.GetDB()
	if db == nil {
		return fmt.Errorf("no database connection")
	}

	// Check if category has products
	var productCount int64
	if err := db.Model(&models.Product{}).Where("category_id = ?", id).Count(&productCount).Error; err != nil {
		return fmt.Errorf("failed to check products: %w", err)
	}

	if productCount > 0 {
		return fmt.Errorf("cannot delete category with %d products. Please reassign or delete the products first", productCount)
	}

	// Delete category
	if err := db.Delete(&models.Category{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

// GetCategoryCount returns the total number of categories
func (s *CategoryService) GetCategoryCount() (int64, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return 0, fmt.Errorf("no database connection")
	}

	var count int64
	if err := db.Model(&models.Category{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count categories: %w", err)
	}

	return count, nil
}

// Helper function to convert model to DTO
func (s *CategoryService) toDTO(category *models.Category) CategoryDTO {
	return CategoryDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Color:       category.Color,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

// GetAll returns all categories as DTOs
func (s *CategoryService) GetAll() ([]CategoryDTO, error) {
	categories, err := s.GetAllCategories()
	if err != nil {
		return nil, err
	}

	dtos := make([]CategoryDTO, len(categories))
	for i, cat := range categories {
		dtos[i] = s.toDTO(&cat)
	}

	return dtos, nil
}

// GetByID returns a category by ID as DTO
func (s *CategoryService) GetByID(id uint) (*CategoryDTO, error) {
	category, err := s.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	dto := s.toDTO(category)
	return &dto, nil
}

// Create creates a new category from DTO
func (s *CategoryService) Create(dto CategoryDTO) (*CategoryDTO, error) {
	category, err := s.CreateCategory(dto.Name, dto.Color)
	if err != nil {
		return nil, err
	}

	// Update description if provided
	if dto.Description != "" {
		db := s.dbManager.GetDB()
		category.Description = dto.Description
		if err := db.Save(category).Error; err != nil {
			return nil, fmt.Errorf("failed to update description: %w", err)
		}
	}

	resultDTO := s.toDTO(category)
	return &resultDTO, nil
}

// Update updates a category from DTO
func (s *CategoryService) Update(id uint, dto CategoryDTO) (*CategoryDTO, error) {
	db := s.dbManager.GetDB()
	if db == nil {
		return nil, fmt.Errorf("no database connection")
	}

	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// Update fields
	category.Name = dto.Name
	category.Description = dto.Description
	category.Color = dto.Color

	if err := db.Save(&category).Error; err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	resultDTO := s.toDTO(&category)
	return &resultDTO, nil
}

// Delete deletes a category by ID
func (s *CategoryService) Delete(id uint) error {
	return s.DeleteCategory(id)
}
