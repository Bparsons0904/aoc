const std = @import("std");
const grid = @import("../grid.zig");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const Interval = struct {
    start: u32,
    end: u32,
};

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    
    const filename = "day9.part1";
    const redTiles = try getRedTiles(allocator, filename);
    defer allocator.free(redTiles);
    
    const largestAreaPart1 = getLargestArea(redTiles);
    
    const intervals = try getInsideIntervals(allocator, redTiles);
    defer allocator.free(intervals);
    
    print("Part 1: {d}, Part 2: {d}\n", .{ largestAreaPart1, intervals.len });
}

fn getRedTiles(allocator: std.mem.Allocator, filename: []const u8) ![]grid.Point {
    _ = allocator;
    _ = filename;
    return &[_]grid.Point{};
}

fn getLargestArea(redTiles: []const grid.Point) u32 {
    _ = redTiles;
    return 0;
}

fn getInsideIntervals(allocator: std.mem.Allocator, redTiles: []const grid.Point) ![]Interval {
    _ = allocator;
    _ = redTiles;
    return &[_]Interval{};
}