# GameOfLife

Terminal-based implementation of Conway's Game of Life written in Go.
The simulation runs in real-time in the terminal, with customizable grid size.

When number of rows and columns is set, initial grid is generated and next iteration occurs after 200 miliseconds which animate Game of Life rules.

## Features
- Randomly seeded initial grid
- Toroidal (wrapping) world boundaries
- Real-time terminal display with ASCII graphics
- Clean, modular code structure using Go packages

## Rules (Conway’s Game of Life)

Each cell can be alive (`■`) or dead (` `), and evolves based on its 8 neighbors:

- **Live cell + 2 or 3 neighbors → lives**
- **Live cell + <2 or >3 neighbors → dies**
- **Dead cell + exactly 3 neighbors → becomes alive**

## Project structure
```
GameOfLife/
├── go.mod
├── main.go # Entry point: user input & game loop
├── internal/
│ ├── grid/
│ │ └── grid.go # Grid creation, printing, logic
│ └── util/
│ │ └── input.go # Helper for user input
└── README.md # You're here!
```

## How to run?
### 1. Clone the repositorz
```bash
git clone https://github.com/goretante/GameOfLife
cd GameOfLife
```

### 2. Initialize Go modules
```bash
go mod tidy
```

### 3. Run the program
```bash
go run main.go
```

## Dependencies
- Go 1.18+
- No third-party packages required

## Planned improvements
- Usage of goroutines
- Add keyboard input to pause/quit
- Creating test case packages and making simple CI/CD pipeline