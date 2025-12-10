# Advent of Code 2025 - Zig Implementation

This is a Zig implementation of the Advent of Code 2025 challenges, converted from the original Go version.

## Prerequisites

You need to have Zig installed on your system. You can install it from https://ziglang.org/download/

## Building and Running

```bash
# Build the project
zig build

# Run the current day's challenge
zig build run

# Or build and run directly
zig build-exe src/main.zig --name aoc2025
./aoc2025
```

## Structure

- `src/main.zig` - Main entry point that runs the appropriate day's challenge
- `src/challenges.zig` - Implementation of all the day challenges
- `src/file_parser.zig` - Utilities for reading input files
- `files/` - Input files for the challenges (copied from Go version)

## Progress

Currently implemented:
- Day 1 (both parts)
- Day 2 (all 5 variants)

## Differences from Go Version

- Uses Zig's memory management with allocators instead of Go's garbage collection
- Uses Zig's error handling system instead of Go's error returns
- Uses Zig's structured error values instead of Go's error types
- Linked list implementation uses explicit memory allocation
- String operations use Zig's slice semantics

## Running Specific Days

To run a specific day's challenge, you can modify `src/main.zig` or create a new file with just the day's functions.