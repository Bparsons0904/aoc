# Advent of Code 2025 - Gleam Solutions

This directory contains Gleam implementations of the Advent of Code 2025 solutions.

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
- `src/day1.gleam` - Day 1 solution
- `gleam.toml` - Project configuration
- `../input/` - Input files (shared with other language implementations)

## Implementation Notes

### Day 1: Dial Puzzle

The solution implements a circular dial with 100 positions (0-99) and processes left/right movement instructions:

- **Part 1**: Count how many instructions end with the dial at position 0
- **Part 2**: Count total zero-crossings during all movements

The Gleam implementation uses a functional approach with immutable state, mirroring the optimized version from the Go solution that calculates position changes mathematically rather than simulating each step.

## Key Gleam Patterns Used

- **Pattern matching**: Used extensively for direction handling and conditional logic
- **Immutable state**: State is passed through fold operations rather than mutated
- **Pipelines**: String processing uses the `|>` operator for readable transformations
- **Result types**: File I/O and parsing use Result types for error handling
- **Assert unwrapping**: Using `assert Ok(value)` for cases where failure indicates a bug
