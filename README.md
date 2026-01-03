# Stock Tracking System

A modern, portable stock tracking desktop application built with Wails (Go + Vue.js). Manage multiple databases from USB drives, track inventory, and monitor stock movements with an intuitive dark/light theme interface.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Wails](https://img.shields.io/badge/Wails-v2-673ab8?logo=wails)
![Vue](https://img.shields.io/badge/Vue-3-4FC08D?logo=vue.js)

## Features

### Core Functionality
- **Multi-Database Support**: Create and switch between multiple SQLite databases
- **Portable Design**: Store databases on USB drives with relative path support
- **Product Management**: Full CRUD operations with filtering and search
- **Category Management**: Color-coded categories with product counts
- **Stock Movements**: Track IN/OUT operations with automatic stock calculation
- **Dashboard Statistics**: Real-time overview of products, categories, and movements
- **Pagination**: Browse large datasets (25 items per page) efficiently

### User Interface
- **Dark/Light Theme**: Toggle between themes with persistent settings
- **Modern Design**: Clean, responsive interface built with Tailwind CSS
- **Database Switcher**: Quick database change with app reload
- **Navigation**: Intuitive back buttons and menu system
- **Color Coding**: Visual category identification
- **Real-time Updates**: Automatic data refresh across views

### Technical Features
- **Zero Dependencies**: Single executable, no external database required
- **Type-Safe API**: Go backend with Wails bindings to Vue.js frontend
- **State Management**: Pinia stores for reactive data handling
- **Routing**: Vue Router for seamless navigation
- **ORM**: GORM for elegant database operations
- **Configuration**: JSON-based settings (theme, last database)

## Prerequisites

- Go 1.21 or higher
- Node.js 16+ and npm
- Wails CLI v2
- Windows (primary target, but can be adapted for Linux/macOS)

## Quick Start

### 1. Install Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 2. Clone and Setup

```bash
git clone <repository-url>
cd StokTakipDesktop
```

### 3. Install Dependencies

```bash
# Backend dependencies
go mod download

# Frontend dependencies
cd frontend
npm install
cd ..
```

### 4. Run Development Mode

```bash
wails dev
```

The application will open with hot-reload enabled for both frontend and backend.

### 5. Build Production

```bash
wails build
```

Find the executable in `build/bin/`.

## Project Structure

```
StokTakipDesktop/
├── frontend/                 # Vue.js Frontend
│   ├── src/
│   │   ├── views/           # Page components
│   │   │   ├── Dashboard.vue        # Main dashboard
│   │   │   ├── DatabaseSelector.vue # DB selection screen
│   │   │   ├── Products.vue         # Product management
│   │   │   ├── Categories.vue       # Category management
│   │   │   └── Movements.vue        # Stock movements
│   │   ├── stores/          # Pinia state stores
│   │   │   ├── database.js
│   │   │   ├── products.js
│   │   │   ├── categories.js
│   │   │   ├── movements.js
│   │   │   └── theme.js
│   │   ├── router/          # Vue Router config
│   │   └── App.vue          # Root component
│   ├── package.json
│   └── tailwind.config.js
│
├── internal/                 # Go Backend
│   ├── app/                 # Application core
│   │   └── app.go           # Main app struct & Wails bindings
│   ├── models/              # Database models (GORM)
│   │   ├── product.go
│   │   ├── category.go
│   │   └── movement.go
│   ├── database/            # Database connection manager
│   │   └── connection.go    # Singleton connection handler
│   ├── services/            # Business logic
│   │   ├── database_service.go
│   │   ├── product_service.go
│   │   ├── category_service.go
│   │   └── movement_service.go
│   ├── config/              # Configuration management
│   │   └── manager.go       # JSON config handler
│   └── utils/               # Utility functions
│       └── path_manager.go  # Portable path handling
│
├── main.go                  # Application entry point
├── wails.json              # Wails configuration
├── go.mod                  # Go module definition
└── README.md               # This file
```

## Configuration

The application stores configuration in `config.json` (created on first run):

```json
{
  "last_database": "database1.db",
  "theme": "dark",
  "language": "en"
}
```

### Configuration Options

- **last_database**: Auto-connect to this database on startup
- **theme**: UI theme (`light` or `dark`)
- **language**: Interface language (currently `en` or `tr`)

## Database Schema

### Products Table
```sql
CREATE TABLE products (
  id            INTEGER PRIMARY KEY AUTOINCREMENT,
  code          TEXT NOT NULL UNIQUE,
  name          TEXT NOT NULL,
  category_id   INTEGER NOT NULL,
  unit          TEXT NOT NULL,
  current_stock REAL DEFAULT 0,
  critical_limit REAL DEFAULT 0,
  price         REAL DEFAULT 0,
  created_at    DATETIME,
  updated_at    DATETIME,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

### Categories Table
```sql
CREATE TABLE categories (
  id          INTEGER PRIMARY KEY AUTOINCREMENT,
  name        TEXT NOT NULL UNIQUE,
  description TEXT,
  color       TEXT DEFAULT '#3B82F6',
  created_at  DATETIME,
  updated_at  DATETIME
);
```

### Movements Table
```sql
CREATE TABLE movements (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  product_id INTEGER NOT NULL,
  type       TEXT NOT NULL CHECK(type IN ('IN', 'OUT')),
  quantity   REAL NOT NULL,
  note       TEXT,
  created_at DATETIME,
  FOREIGN KEY (product_id) REFERENCES products(id)
);
```

## Usage Guide

### First Launch

1. **Database Selection**: Choose an existing database or create a new one
2. **Create Categories**: Set up product categories with colors
3. **Add Products**: Define your inventory with codes, names, and units
4. **Track Movements**: Record stock IN/OUT operations

### Managing Databases

**Switch Database:**
1. Click "Change Database" button in the header
2. Application reloads automatically
3. Select a different database or create a new one

**Create New Database:**
1. Click "Create New Database" in the selector
2. Enter a unique name (e.g., "warehouse_2024")
3. Database file is created in the `Data/` folder

**Portable Usage:**
- Copy the entire application folder to a USB drive
- All databases in `Data/` folder will be accessible
- Configuration is preserved across different computers

### Working with Data

**Products:**
- Filter by category or stock status
- Search by name or code
- View current stock levels with color indicators
- 25 products per page with pagination

**Categories:**
- Color-code for visual identification
- Track product count per category
- Edit or delete (if no products assigned)
- Grid view with 25 items per page

**Stock Movements:**
- IN: Add stock (purchases, returns)
- OUT: Remove stock (sales, usage)
- Automatic stock calculation
- Optional notes for each movement

## Themes

Toggle between Light and Dark themes using the button in the header.

**Light Theme:**
- White backgrounds
- Dark text for readability
- Subtle shadows

**Dark Theme:**
- Dark gray backgrounds
- Light text
- Blue accents

Theme preference is saved in `config.json` and persists across sessions.

## Data Security

- **Local Storage**: All data stored locally in SQLite files
- **No Cloud**: No internet connection required
- **Backup Friendly**: Simple `.db` files easy to backup
- **Portable**: Take your data anywhere on USB drive

## Troubleshooting

### Database Not Opening

1. Check file permissions in `Data/` folder
2. Ensure database file is not corrupted
3. Try creating a new database

### Application Won't Start

1. Verify Go and Node.js are installed
2. Run `go mod tidy` to update dependencies
3. Clear frontend cache: `cd frontend && npm clean-install`

### Build Fails

1. Update Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
2. Check system requirements (Windows build tools, etc.)
3. Run `wails doctor` to diagnose issues

## Distribution

### Windows

Build with embedded assets:

```bash
wails build -clean
```

Output: `build/bin/StokTakipDesktop.exe`

**Portable Package:**
1. Copy the `.exe` to a folder
2. Create `Data/` subfolder (optional, auto-created on first run)
3. Distribute the folder

### Other Platforms

Linux and macOS builds require platform-specific flags:

```bash
# Linux
GOOS=linux GOARCH=amd64 wails build

# macOS
GOOS=darwin GOARCH=amd64 wails build
```

## Roadmap

- Export/Import functionality (Excel, CSV)
- Barcode scanning support
- Reports and analytics
- Multi-language support
- Cloud backup integration
- User authentication
- Purchase order management

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Wails](https://wails.io/) - Fantastic Go + Web framework
- [Vue.js](https://vuejs.org/) - Progressive JavaScript framework
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [GORM](https://gorm.io/) - Developer-friendly ORM for Go
- [Pinia](https://pinia.vuejs.org/) - Vue.js state management

## Contact

For questions or support, please open an issue on GitHub.

---

Built with Wails, Go, and Vue.js
