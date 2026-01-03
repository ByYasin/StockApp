package utils

import (
	"os"
	"path/filepath"
)

// PathManager manages all file paths for the application
type PathManager struct {
	executablePath string
	rootPath       string
}

// NewPathManager creates a new PathManager instance
func NewPathManager() (*PathManager, error) {
	execPath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	// Get the directory of the executable
	rootPath := filepath.Dir(execPath)

	return &PathManager{
		executablePath: execPath,
		rootPath:       rootPath,
	}, nil
}

// GetRootPath returns the root directory of the application
func (pm *PathManager) GetRootPath() string {
	return pm.rootPath
}

// GetDataFolder returns the path to the Data folder
func (pm *PathManager) GetDataFolder() string {
	return filepath.Join(pm.rootPath, "Data")
}

// GetConfigPath returns the path to the config.json file
func (pm *PathManager) GetConfigPath() string {
	return filepath.Join(pm.rootPath, "config.json")
}

// GetDatabasePath returns the full path to a database file
func (pm *PathManager) GetDatabasePath(filename string) string {
	return filepath.Join(pm.GetDataFolder(), filename)
}

// EnsureDataFolder creates the Data folder if it doesn't exist
func (pm *PathManager) EnsureDataFolder() error {
	dataFolder := pm.GetDataFolder()
	return os.MkdirAll(dataFolder, 0755)
}

// ListDatabaseFiles returns a list of all .db files in the Data folder
func (pm *PathManager) ListDatabaseFiles() ([]string, error) {
	dataFolder := pm.GetDataFolder()

	// Ensure folder exists
	if err := pm.EnsureDataFolder(); err != nil {
		return nil, err
	}

	// Read directory
	files, err := os.ReadDir(dataFolder)
	if err != nil {
		return nil, err
	}

	// Filter .db files
	var dbFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".db" {
			dbFiles = append(dbFiles, file.Name())
		}
	}

	return dbFiles, nil
}

// FileExists checks if a file exists
func (pm *PathManager) FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetFileSize returns the size of a file in bytes
func (pm *PathManager) GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
