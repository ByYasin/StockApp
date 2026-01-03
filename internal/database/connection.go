package database

import (
	"database/sql"
	"fmt"
	"log"
	"stoktakip/internal/models"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // Pure Go SQLite driver (no CGO needed!)
)

// ConnectionManager manages database connections (Singleton pattern)
type ConnectionManager struct {
	db    *gorm.DB
	mutex sync.RWMutex
}

var (
	instance *ConnectionManager
	once     sync.Once
)

// GetConnectionManager returns the singleton instance of ConnectionManager
func GetConnectionManager() *ConnectionManager {
	once.Do(func() {
		instance = &ConnectionManager{}
	})
	return instance
}

// Connect opens a connection to the specified database file
func (cm *ConnectionManager) Connect(dbPath string) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// Close existing connection if any
	if cm.db != nil {
		sqlDB, err := cm.db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	// Open database with modernc.org/sqlite (pure Go, no CGO required)
	// First open with database/sql to use modernc driver
	sqlDB, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database with modernc sqlite: %w", err)
	}

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Now wrap with GORM using the existing connection
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		sqlDB.Close()
		return fmt.Errorf("failed to initialize GORM: %w", err)
	}

	// Configure connection pool
	poolDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	poolDB.SetMaxOpenConns(10)
	poolDB.SetMaxIdleConns(5)

	// Run migrations
	if err := cm.runMigrations(db); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	cm.db = db
	log.Printf("Successfully connected to database: %s", dbPath)

	return nil
}

// GetDB returns the current database connection
func (cm *ConnectionManager) GetDB() *gorm.DB {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.db
}

// IsConnected checks if there's an active database connection
func (cm *ConnectionManager) IsConnected() bool {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.db != nil
}

// Close closes the current database connection
func (cm *ConnectionManager) Close() error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if cm.db == nil {
		return nil
	}

	sqlDB, err := cm.db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	cm.db = nil
	log.Println("Database connection closed")
	return nil
}

// runMigrations automatically migrates the database schema
func (cm *ConnectionManager) runMigrations(db *gorm.DB) error {
	// Auto migrate all models
	if err := db.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.StockMovement{},
	); err != nil {
		return err
	}

	// Seed default categories if database is empty
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count == 0 {
		if err := cm.seedDefaultCategories(db); err != nil {
			log.Printf("Warning: Failed to seed default categories: %v", err)
		}
	}

	return nil
}

// seedDefaultCategories creates default categories
func (cm *ConnectionManager) seedDefaultCategories(db *gorm.DB) error {
	defaultCategories := []models.Category{
		{Name: "Genel", Color: "#6B7280"},
		{Name: "Elektronik", Color: "#3B82F6"},
		{Name: "Yedek Parça", Color: "#EF4444"},
		{Name: "Malzeme", Color: "#10B981"},
		{Name: "Kimyasal", Color: "#F59E0B"},
		{Name: "El Aletleri", Color: "#8B5CF6"},
	}

	for _, category := range defaultCategories {
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}

	log.Println("Default categories seeded successfully")
	return nil
}
