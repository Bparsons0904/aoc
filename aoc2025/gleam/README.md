# Advent of Code 2025 - Gleam Solutions

This directory contains Gleam implementations of the Advent of Code 2025 solutions, ported from the Go implementations.

## Setup

### Install Gleam

```bash
# On Linux/macOS with Homebrew
brew install gleam

# Or download from https://gleam.run/getting-started/installing/
```

### Install Dependencies

```bash
cd gleam
gleam deps download
```

## Running Solutions

```bash
# Run all solutions
gleam run

# Run in development mode with auto-rebuild
gleam run --watch

# Build the project
gleam build
```

## Project Structure

- `src/aoc2025.gleam` - Main entry point
- `src/day1.gleam` - Day 1: Dial puzzle solution
- `src/day2.gleam` - Day 2: Product ID ranges solution
- `src/day3.gleam` - Day 3: Battery pack joltage solution
- `src/day5.gleam` - Day 5: Fresh ingredients ranges solution
- `src/day6.gleam` - Day 6: Worksheet calculations solution
- `gleam.toml` - Project configuration
- `input/` - Input files (copied from ../go/files/)

## Implemented Solutions

### Day 1: Dial Puzzle
Circular dial with 100 positions (0-99) processing left/right movement instructions. Uses mathematical calculations rather than step-by-step simulation.

### Day 2: Product ID Validation
Finds product IDs with specific patterns:
- Part 1: IDs where first half equals second half
- Part 2: IDs with repeating patterns (e.g., "123123", "7777")

### Day 3: Battery Pack Joltage
Selects maximum joltage digits from battery packs to form optimal battery configurations.

### Day 5: Fresh Ingredient Ranges
Merges overlapping ingredient ID ranges and counts:
- Part 1: How many actual ingredients fall within fresh ranges
- Part 2: Total capacity of all merged fresh ingredient ranges

### Day 6: Worksheet Calculations
Processes worksheets with vertical and horizontal value arrangements:
- Part 1: Calculate totals using horizontal values
- Part 2: Calculate totals using vertical (cephalopod) values

## Not Yet Implemented

- **Day 4**: Grid-based paper roll simulation (requires complex grid system)
- **Day 7**: Tachyon beam grid pathfinding (requires grid system)
- **Day 8**: Junction box 3D clustering (complex union-find algorithm)
- **Day 9+**: Not yet implemented in Go

## Key Gleam Patterns Used

- **Pattern matching**: Used extensively for direction handling and conditional logic
- **Immutable state**: State is passed through fold operations rather than mutated
- **Pipelines**: String processing uses the `|>` operator for readable transformations
- **Result types**: File I/O and parsing use Result types for error handling
- **Custom types**: Type-safe data structures (DialInstruction, ProductIDRange, etc.)
- **Recursive functions**: For list processing and iterative calculations
- **Dict operations**: For maintaining indexed collections of worksheets
