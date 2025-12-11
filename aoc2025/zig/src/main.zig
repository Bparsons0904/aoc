const std = @import("std");
const challenges = @import("challenges.zig");
const print = std.debug.print;

pub fn main() !void {
    print("=== Day 4 ===\n", .{});
    try challenges.day4.run();
}