package services

import (
	"fmt"
	"os"
	"path/filepath"
	"stoktakip/internal/config"
	"stoktakip/internal/database"
	"stoktakip/internal/utils"
	"time"
)

// DatabaseInfo contains information about a database file
type DatabaseInfo struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Size     float64 `json:"size"`      // Size in MB
	Modified string  `json:"modified"`  // Last modified date
	IsActive bool    `json:"is_active"` // Currently connected
}

// DatabaseService handles database-related operations
type DatabaseService struct {
	dbManager     *database.ConnectionManager
	pathManager   *utils.PathManager
	configManager *config.Manager
}

// NewDatabaseService creates a new database service
func NewDatabaseService(
	dbManager *database.ConnectionManager,
	pathManager *utils.PathManager,
	configManager *config.Manager,
) *DatabaseService {
	return &DatabaseService{
		dbManager:     dbManager,
		pathManager:   pathManager,
		configManager: configManager,
	}
}

// ListDatabases returns a list of all database files in the Data folder
func (s *DatabaseService) ListDatabases() ([]DatabaseInfo, error) {
	dbFiles, err := s.pathManager.ListDatabaseFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to list database files: %w", err)
	}

	var databases []DatabaseInfo
	for _, filename := range dbFiles {
		fullPath := s.pathManager.GetDatabasePath(filename)

		// Get file info
		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		// Calculate size in MB
		sizeMB := float64(info.Size()) / (1024 * 1024)

		databases = append(databases, DatabaseInfo{
			Name:     filename,
			Path:     fullPath,
			Size:     sizeMB,
			Modified: info.ModTime().Format("2006-01-02 15:04:05"),
			IsActive: false, // Will be set below
		})
	}

	return databases, nil
}

// CreateDatabase creates a new database file with the given name
func (s *DatabaseService) CreateDatabase(name string) error {
	// Step 1: Validate name
	if name == "" {
		return fmt.Errorf("veritabanı adı boş olamaz")
	}

	// Step 2: Ensure .db extension
	if filepath.Ext(name) != ".db" {
		name = name + ".db"
	}

	// Step 3: Get and log paths
	dataFolder := s.pathManager.GetDataFolder()
	fmt.Printf("Data folder path: %s\n", dataFolder)

	// Step 4: Ensure Data folder exists
	if err := s.pathManager.EnsureDataFolder(); err != nil {
		return fmt.Errorf("Data klasörü oluşturulamadı (%s): %w", dataFolder, err)
	}
	fmt.Printf("Data folder ensured\n")

	// Step 5: Check folder is writable by trying to create a temp file
	testFile := filepath.Join(dataFolder, ".test_write")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return fmt.Errorf("Data klasörüne yazma izni yok (%s): %w", dataFolder, err)
	}
	os.Remove(testFile)
	fmt.Printf("Write permission OK\n")

	// Step 6: Get database path
	dbPath := s.pathManager.GetDatabasePath(name)
	fmt.Printf("Database path: %s\n", dbPath)

	// Step 7: Check if file already exists
	if s.pathManager.FileExists(dbPath) {
		return fmt.Errorf("'%s' adında veritabanı zaten var", name)
	}

	// Step 8: Create and connect to the new database
	fmt.Printf("Connecting to database...\n")
	if err := s.dbManager.Connect(dbPath); err != nil {
		return fmt.Errorf("veritabanı bağlantısı başarısız: %w", err)
	}
	fmt.Printf("Database connected successfully\n")

	// Step 9: Update config with new database (save only filename for portability)
	if err := s.configManager.SetLastDatabase(name); err != nil {
		return fmt.Errorf("yapılandırma kaydedilemedi: %w", err)
	}
	fmt.Printf("Config saved\n")

	return nil
}

// SwitchDatabase switches to a different database
func (s *DatabaseService) SwitchDatabase(path string) error {
	// Check if file exists
	if !s.pathManager.FileExists(path) {
		return fmt.Errorf("database file not found: %s", path)
	}

	// Connect to the database
	if err := s.dbManager.Connect(path); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Update config (save only filename for portability)
	filename := filepath.Base(path)
	if err := s.configManager.SetLastDatabase(filename); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	return nil
}

// GetCurrentDatabase returns information about the currently connected database
func (s *DatabaseService) GetCurrentDatabase() (*DatabaseInfo, error) {
	if !s.dbManager.IsConnected() {
		return nil, fmt.Errorf("no database connected")
	}

	lastDB := s.configManager.GetLastDatabase()
	if lastDB == "" {
		return nil, fmt.Errorf("no database path in config")
	}

	// Get file info
	info, err := os.Stat(lastDB)
	if err != nil {
		return nil, fmt.Errorf("failed to get database info: %w", err)
	}

	sizeMB := float64(info.Size()) / (1024 * 1024)

	return &DatabaseInfo{
		Name:     filepath.Base(lastDB),
		Path:     lastDB,
		Size:     sizeMB,
		Modified: info.ModTime().Format("2006-01-02 15:04:05"),
		IsActive: true,
	}, nil
}

// DeleteDatabase deletes a database file
func (s *DatabaseService) DeleteDatabase(path string) error {
	// Check if it's the currently connected database
	currentDB := s.configManager.GetLastDatabase()
	if currentDB == path {
		return fmt.Errorf("cannot delete currently connected database")
	}

	// Check if file exists
	if !s.pathManager.FileExists(path) {
		return fmt.Errorf("database file not found: %s", path)
	}

	// Delete the file
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to delete database: %w", err)
	}

	return nil
}

// BackupDatabase creates a backup of the current database
func (s *DatabaseService) BackupDatabase() (string, error) {
	if !s.dbManager.IsConnected() {
		return "", fmt.Errorf("no database connected")
	}

	currentDB := s.configManager.GetLastDatabase()
	if currentDB == "" {
		return "", fmt.Errorf("no database path in config")
	}

	// Create backup filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	backupName := fmt.Sprintf("%s_backup_%s.db",
		filepath.Base(currentDB[:len(currentDB)-3]), // Remove .db extension
		timestamp,
	)

	backupPath := s.pathManager.GetDatabasePath(backupName)

	// Read source file
	data, err := os.ReadFile(currentDB)
	if err != nil {
		return "", fmt.Errorf("failed to read database: %w", err)
	}

	// Write backup file
	if err := os.WriteFile(backupPath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write backup: %w", err)
	}

	return backupPath, nil
}

// IsConnected checks if there's an active database connection
func (s *DatabaseService) IsConnected() bool {
	return s.dbManager.IsConnected()
}
