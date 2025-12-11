const std = @import("std");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const WorksheetMap = std.StringHashMap(Worksheet);

const Worksheet = struct {
    values: []u32,
    cephalopodValues: []u32,
    operator: u8,
    allocator: std.mem.Allocator,
    
    fn init(allocator: std.mem.Allocator) Worksheet {
        return Worksheet{
            .values = &[_]u32{},
            .cephalopodValues = &[_]u32{},
            .operator = '+',
            .allocator = allocator,
        };
    }
    
    fn deinit(self: *Worksheet) void {
        self.allocator.free(self.values);
        self.allocator.free(self.cephalopodValues);
    }
};

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    
    var worksheet = WorksheetMap.init(allocator);
    defer worksheet.deinit();
    
    const part1Total, const part2Total = calculateWorksheets(worksheet);
    
    print("Part 1: {d}, Part 2: {d}\n", .{ part1Total, part2Total });
}

fn calculateWorksheets(worksheet: WorksheetMap) struct { u32, u32 } {
    _ = worksheet;
    return .{ 0, 0 };
}