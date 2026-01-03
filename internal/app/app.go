package app

import (
	"context"
	"fmt"
	"log"
	"stoktakip/internal/config"
	"stoktakip/internal/database"
	"stoktakip/internal/services"
	"stoktakip/internal/utils"
)

// App struct
type App struct {
	ctx             context.Context
	pathManager     *utils.PathManager
	configManager   *config.Manager
	dbManager       *database.ConnectionManager
	databaseService *services.DatabaseService
	productService  *services.ProductService
	categoryService *services.CategoryService
	movementService *services.MovementService
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	// Initialize path manager
	pathManager, err := utils.NewPathManager()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize path manager: %w", err)
	}

	// Ensure Data folder exists
	if err := pathManager.EnsureDataFolder(); err != nil {
		return nil, fmt.Errorf("failed to create Data folder: %w", err)
	}

	// Initialize config manager
	configManager := config.NewManager(pathManager)

	// Load configuration
	if _, err := configManager.Load(); err != nil {
		log.Printf("Warning: Failed to load config: %v", err)
	}

	// Initialize database manager
	dbManager := database.GetConnectionManager()

	// Initialize services immediately
	databaseService := services.NewDatabaseService(dbManager, pathManager, configManager)
	productService := services.NewProductService(dbManager)
	categoryService := services.NewCategoryService(dbManager)
	movementService := services.NewMovementService(dbManager)

	app := &App{
		pathManager:     pathManager,
		configManager:   configManager,
		dbManager:       dbManager,
		databaseService: databaseService,
		productService:  productService,
		categoryService: categoryService,
		movementService: movementService,
	}

	return app, nil
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	log.Println("Application started")

	// Try to connect to last used database
	lastDB := a.configManager.GetLastDatabase()
	if lastDB != "" && a.pathManager.FileExists(lastDB) {
		if err := a.dbManager.Connect(lastDB); err != nil {
			log.Printf("Warning: Failed to connect to last database: %v", err)
		} else {
			log.Printf("Reconnected to last database: %s", lastDB)
		}
	}
}

// Shutdown is called when the app is closing
func (a *App) Shutdown(ctx context.Context) {
	log.Println("Application shutting down")

	// Close database connection
	if err := a.dbManager.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	// Save configuration
	if err := a.configManager.Save(); err != nil {
		log.Printf("Error saving config: %v", err)
	}
}

// GetDatabaseService returns the database service (for Wails binding)
func (a *App) GetDatabaseService() *services.DatabaseService {
	return a.databaseService
}

// GetProductService returns the product service (for Wails binding)
func (a *App) GetProductService() *services.ProductService {
	return a.productService
}

// GetCategoryService returns the category service (for Wails binding)
func (a *App) GetCategoryService() *services.CategoryService {
	return a.categoryService
}

// GetMovementService returns the movement service (for Wails binding)
func (a *App) GetMovementService() *services.MovementService {
	return a.movementService
}

// Database service methods - exported for Wails

// ListDatabases returns all database files
func (a *App) ListDatabases() ([]services.DatabaseInfo, error) {
	return a.databaseService.ListDatabases()
}

// CreateDatabase creates a new database
func (a *App) CreateDatabase(name string) error {
	return a.databaseService.CreateDatabase(name)
}

// SwitchDatabase switches to a different database
func (a *App) SwitchDatabase(path string) error {
	return a.databaseService.SwitchDatabase(path)
}

// GetCurrentDatabase returns the currently connected database info
func (a *App) GetCurrentDatabase() (*services.DatabaseInfo, error) {
	return a.databaseService.GetCurrentDatabase()
}

// IsConnected checks if database is connected
func (a *App) IsConnected() bool {
	return a.databaseService.IsConnected()
}

// Category service methods - exported for Wails

// GetAllCategories returns all categories
func (a *App) GetAllCategories() ([]services.CategoryDTO, error) {
	return a.categoryService.GetAll()
}

// GetCategoryByID returns a category by ID
func (a *App) GetCategoryByID(id uint) (*services.CategoryDTO, error) {
	return a.categoryService.GetByID(id)
}

// CreateCategory creates a new category
func (a *App) CreateCategory(dto services.CategoryDTO) (*services.CategoryDTO, error) {
	return a.categoryService.Create(dto)
}

// UpdateCategory updates an existing category
func (a *App) UpdateCategory(id uint, dto services.CategoryDTO) (*services.CategoryDTO, error) {
	return a.categoryService.Update(id, dto)
}

// DeleteCategory deletes a category
func (a *App) DeleteCategory(id uint) error {
	return a.categoryService.Delete(id)
}

// Product service methods - exported for Wails

// GetAllProducts returns all products
func (a *App) GetAllProducts() ([]services.ProductDTO, error) {
	return a.productService.GetAll()
}

// GetProductByID returns a product by ID
func (a *App) GetProductByID(id uint) (*services.ProductDTO, error) {
	return a.productService.GetByID(id)
}

// CreateProduct creates a new product
func (a *App) CreateProduct(dto services.ProductDTO) (*services.ProductDTO, error) {
	return a.productService.Create(dto)
}

// UpdateProduct updates an existing product
func (a *App) UpdateProduct(id uint, dto services.ProductDTO) (*services.ProductDTO, error) {
	return a.productService.Update(id, dto)
}

// DeleteProduct deletes a product
func (a *App) DeleteProduct(id uint) error {
	return a.productService.Delete(id)
}

// GetLowStockProducts returns products with low stock
func (a *App) GetLowStockProducts() ([]services.ProductDTO, error) {
	return a.productService.GetLowStock()
}

// Movement service methods - exported for Wails

// GetAllMovements returns all movements
func (a *App) GetAllMovements() ([]services.MovementDTO, error) {
	return a.movementService.GetAll()
}

// GetMovementByID returns a movement by ID
func (a *App) GetMovementByID(id uint) (*services.MovementDTO, error) {
	return a.movementService.GetByID(id)
}

// CreateMovement creates a new movement
func (a *App) CreateMovement(dto services.MovementDTO) (*services.MovementDTO, error) {
	return a.movementService.Create(dto)
}

// DeleteMovement deletes a movement
func (a *App) DeleteMovement(id uint) error {
	return a.movementService.Delete(id)
}

// GetMovementsByProduct returns movements for a specific product
func (a *App) GetMovementsByProduct(productID uint) ([]services.MovementDTO, error) {
	return a.movementService.GetByProduct(productID)
}

// GetMovementStats returns movement statistics
func (a *App) GetMovementStats() (*services.MovementStats, error) {
	return a.movementService.GetStats()
}
