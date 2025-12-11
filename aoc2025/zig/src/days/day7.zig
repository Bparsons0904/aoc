const std = @import("std");
const grid = @import("../grid.zig");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const TachyonGraph = std.AutoHashMap(grid.Point, u32);

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    var g = try grid.Grid.init(allocator, "day7.part1");
    defer g.deinit();
    
    const part1Count = processTachyonBeamSplitCounter(g);
    const part2Count = processTachyonBeamRoutesCounter(g);
    
    print("Part 1: {d}, Part 2: {d}\n", .{ part1Count, part2Count });
}

fn processTachyonBeamSplitCounter(g: grid.Grid) u32 {
    _ = g;
    return 0;
}

fn processTachyonBeamRoutesCounter(g: grid.Grid) u32 {
    _ = g;
    return 0;
}