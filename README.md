#Charm Pokedex

A beautiful Terminal User Interface (TUI) Pokemon explorer built with Go, Bubble Tea, and Lipgloss. Browse, search, and view detailed information about Pokemon directly in your terminal!

![Go Version](https://img.shields.io/badge/Go-1.24.4-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)

## ✨ Features

- 🎮 **Interactive Pokemon List** - Browse through 150+ Pokemon with fuzzy search
- 🔍 **Smart Search** - Search by Pokemon name or ID number
- 📊 **Detailed Stats** - View comprehensive Pokemon information including:
  - Types with color-coded badges
  - Base stats with visual progress bars
  - Height, weight, and base experience
  - Abilities (including hidden abilities)
- ⌨️ **Keyboard Navigation** - Fully keyboard-driven interface
- 🎨 **Beautiful UI** - Clean, colorful terminal interface
- ⚡ **Real-time Data** - Fetches data from PokeAPI
- 🚀 **Fast & Responsive** - Optimized for smooth user experience

## 📷 Screenshots

### Main Pokemon List
![dashboard](/asssets/dashboard.png)


### Pokemon Details View
![details](/asssets/detail.png)

### Search Interface
![search](/asssets/search.png)

## 🚀 Installation

### Prerequisites

- Go 1.24.4 or higher
- Terminal with Unicode support
- Internet connection (for PokeAPI)

### Quick Install

```bash
# Clone the repository
git clone https://github.com/fanchann/charm-pokedex.git

# Navigate to project directory
cd charm-pokedex

# Download dependencies
go mod tidy

# Run the application
go run main.go
```

### Build from Source

```bash
# Build binary
go build -o bin/pokedex main.go

# Run the binary
./bin/pokedex
```

### Install Globally

```bash
# Install to $GOPATH/bin
go install github.com/fanchann/charm-pokedex@latest

# Run from anywhere
charm-pokedex
```

## 🎮 Usage

### Keyboard Controls

| Key | Action |
|-----|--------|
| `↑/↓` | Navigate list |
| `Enter` | Select Pokemon / Confirm search |
| `/` | Open search mode |
| `ESC` | Go back / Cancel |
| `b` | Go back (detail view) |
| `r` | Refresh Pokemon list |
| `q` | Quit application |
| `Ctrl+C` | Force quit |

### Navigation Flow

1. **Main List** → Browse Pokemon with arrow keys
2. **Select Pokemon** → Press `Enter` to view details
3. **Search Mode** → Press `/` and type Pokemon name or ID
4. **Go Back** → Press `ESC` or `b` to return to list

### Search Examples

- Search by name: `pikachu`, `charizard`, `mewtwo`
- Search by ID: `25`, `150`, `1`
- Case insensitive: `PIKACHU`, `Charizard`

## 🏗️ Architecture

```
charm-pokedex/
├── go.mod                      # Go module definition
├── go.sum                      # Dependency checksums
├── main.go                     # Application entry point
├── .gitignore                  # Git ignore rules
├── helpers/
│   └── format.go              # Utility formatting functions
└── internal/
    ├── globals/
    │   └── messages.go        # Global message types
    ├── models/
    │   └── pokedex.go         # Data models and structs
    ├── service/
    │   ├── pokedx_api.go      # PokeAPI integration
    │   └── pokedx_api_test.go # API service tests
    └── ui/
        ├── model.go           # UI state management
        ├── styles.go          # Visual styling
        ├── update.go          # Business logic
        └── view.go            # Rendering logic
```

### Technology Stack

- **[Go](https://golang.org/)** - Programming language
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - TUI framework
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - TUI components
- **[PokeAPI](https://pokeapi.co/)** - Pokemon data source

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/service

# Verbose test output
go test -v ./...
```

## 🎨 Customization

### Color Scheme

Edit `internal/ui/styles.go` to customize colors:

```go
var (
    TitleStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFFDF5")).
        Background(lipgloss.Color("#1078eeff"))
    
    // Customize other styles...
)
```

### Data Source

The app uses [PokeAPI](https://pokeapi.co/) by default. To change the data source, modify the API endpoints in `internal/service/pokedx_api.go`.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 🙏 Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing free Pokemon data
- [Charm](https://charm.sh/) for the amazing TUI libraries
- Pokemon Company for creating the amazing Pokemon universe