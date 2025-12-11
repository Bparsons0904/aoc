const std = @import("std");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const JunctionBox = struct {
    x: u32,
    y: u32,
    z: u32,
};

const JunctionBoxGroup = struct {
    connections: []JunctionBox,
    allocator: std.mem.Allocator,
    
    fn init(allocator: std.mem.Allocator) JunctionBoxGroup {
        return JunctionBoxGroup{
            .connections = &[_]JunctionBox{},
            .allocator = allocator,
        };
    }
    
    fn deinit(self: *JunctionBoxGroup) void {
        self.allocator.free(self.connections);
    }
};

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    
    const junctionBoxes = try allocator.alloc(JunctionBox, 16);
    
    var junctionBoxConnectionIndex = std.AutoHashMap(u32, u32).init(allocator);
    defer junctionBoxConnectionIndex.deinit();
    
    const connections = getSortedConnections(junctionBoxes);
    
    const part1Results = buildJunctionBoxConnectionsPart1(connections);
    const part2Results = buildJunctionBoxConnectionsPart2(connections);
    
    print("Part 1: {d}, Part 2: {d}\n", .{ part1Results, part2Results });
    
    allocator.free(junctionBoxes);
}

fn getSortedConnections(junctionBoxes: []const JunctionBox) []const JunctionBox {
    _ = junctionBoxes;
    return &[_]JunctionBox{};
}

fn buildJunctionBoxConnectionsPart1(connections: []const JunctionBox) u32 {
    _ = connections;
    return 0;
}

fn buildJunctionBoxConnectionsPart2(connections: []const JunctionBox) u32 {
    _ = connections;
    return 0;
}