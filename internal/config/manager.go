package config

import (
	"encoding/json"
	"os"
	"stoktakip/internal/utils"
)

// Config represents the application configuration
type Config struct {
	LastDatabase string `json:"last_database"`
	Theme        string `json:"theme"`
	Language     string `json:"language"`
}

// Manager handles configuration file operations
type Manager struct {
	pathManager *utils.PathManager
	config      *Config
}

// NewManager creates a new config manager
func NewManager(pathManager *utils.PathManager) *Manager {
	return &Manager{
		pathManager: pathManager,
		config:      nil,
	}
}

// Load loads the configuration from file
func (m *Manager) Load() (*Config, error) {
	configPath := m.pathManager.GetConfigPath()

	// If config file doesn't exist, return default config
	if !m.pathManager.FileExists(configPath) {
		m.config = m.getDefaultConfig()
		return m.config, nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Parse JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		// If parsing fails, return default config
		m.config = m.getDefaultConfig()
		return m.config, nil
	}

	m.config = &config
	return m.config, nil
}

// Save saves the configuration to file
func (m *Manager) Save() error {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}

	configPath := m.pathManager.GetConfigPath()

	// Marshal to JSON with indentation
	data, err := json.MarshalIndent(m.config, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(configPath, data, 0644)
}

// Get returns the current configuration
func (m *Manager) Get() *Config {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}
	return m.config
}

// SetLastDatabase updates the last used database
func (m *Manager) SetLastDatabase(dbPath string) error {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}
	m.config.LastDatabase = dbPath
	return m.Save()
}

// GetLastDatabase returns the last used database path
func (m *Manager) GetLastDatabase() string {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}
	return m.config.LastDatabase
}

// SetTheme updates the theme
func (m *Manager) SetTheme(theme string) error {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}
	m.config.Theme = theme
	return m.Save()
}

// GetTheme returns the current theme
func (m *Manager) GetTheme() string {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}
	return m.config.Theme
}

// ClearLastDatabase clears the last database setting
func (m *Manager) ClearLastDatabase() error {
	if m.config == nil {
		m.config = m.getDefaultConfig()
	}
	m.config.LastDatabase = ""
	return m.Save()
}

// getDefaultConfig returns the default configuration
func (m *Manager) getDefaultConfig() *Config {
	return &Config{
		LastDatabase: "",
		Theme:        "light",
		Language:     "tr",
	}
}
