const std = @import("std");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const Dial = struct {
    value: u32,
    next: ?*Dial,
    prev: ?*Dial,
};

const DialList = struct {
    head: *Dial,
    tail: *Dial,
    
    fn init() !DialList {
        const allocator = std.heap.page_allocator;
        
        // Create 100 nodes in a circular linked list
        var nodes: [100]*Dial = undefined;
        for (0..100) |i| {
            const node = try allocator.create(Dial);
            node.value = @intCast(i);
            node.next = null;
            node.prev = null;
            nodes[i] = node;
        }
        
        // Link them together
        for (0..100) |i| {
            const next_idx = (i + 1) % 100;
            const prev_idx = if (i == 0) 99 else i - 1;
            
            nodes[i].next = nodes[next_idx];
            nodes[i].prev = nodes[prev_idx];
        }
        
        return DialList{
            .head = nodes[0],
            .tail = nodes[99],
        };
    }
    
    fn next(self: *DialList) void {
        self.head = self.head.next.?;
    }
    
    fn prev(self: *DialList) void {
        self.head = self.head.prev.?;
    }
};

const DialInstruction = struct {
    direction: u8,
    step: u32,
};

pub fn run() !void {
    try day1();
    try day1_1();
}

pub fn day1() !void {
    const allocator = std.heap.page_allocator;
    
    const files = try file_parser.readFile(allocator, "day1.part1");
    defer allocator.free(files);
    
    const dial_instructions = try allocator.alloc(DialInstruction, files.len);
    
    try parseInstructions(files, dial_instructions);
    
    var dial_list = try DialList.init();
    
    // Position at 50 like in the Go version
    while (dial_list.head.value != 50) {
        dial_list.next();
    }
    
    var day1part1Count: u32 = 0;
    var day1part2Count: u32 = 0;
    
    for (dial_instructions) |dial_instruction| {
        switch (dial_instruction.direction) {
            'R' => {
                for (0..dial_instruction.step) |_| {
                    dial_list.next();
                    if (dial_list.head.value == 0) {
                        day1part2Count += 1;
                    }
                }
            },
            'L' => {
                for (0..dial_instruction.step) |_| {
                    dial_list.prev();
                    if (dial_list.head.value == 0) {
                        day1part2Count += 1;
                    }
                }
            },
            else => return error.InvalidDirection,
        }
        
        if (dial_list.head.value == 0) {
            day1part1Count += 1;
        }
    }
    
    print("day1 part1: {d}, part2: {d}\n", .{ day1part1Count, day1part2Count });
}

pub fn day1_1() !void {
    const allocator = std.heap.page_allocator;
    
    const files = try file_parser.readFile(allocator, "day1.part1");
    defer allocator.free(files);
    
    const dial_instructions = try allocator.alloc(DialInstruction, files.len);
    
    try parseInstructions(files, dial_instructions);
    
    var step1Count: u32 = 0;
    var step2Count: u32 = 0;
    var currentValue: i32 = 50;
    
    for (dial_instructions) |dial_instruction| {
        switch (dial_instruction.direction) {
            'R' => {
                const newValue = currentValue + @as(i32, @intCast(dial_instruction.step));
                if (newValue >= 100) {
                    if (currentValue == 0) {
                        step2Count += @intCast(dial_instruction.step / 100);
                    } else {
                        step2Count += @intCast((dial_instruction.step + @as(u32, @intCast(currentValue))) / 100);
                    }
                }
                currentValue = @rem(newValue, 100);
            },
            'L' => {
                const newValue = currentValue - @as(i32, @intCast(dial_instruction.step));
                if (currentValue == 0) {
                    step2Count += @intCast(dial_instruction.step / 100);
                } else if (dial_instruction.step >= @as(u32, @intCast(currentValue))) {
                    step2Count += @intCast((dial_instruction.step - @as(u32, @intCast(currentValue))) / 100 + 1);
                }
                currentValue = @mod(newValue, 100);
            },
            else => return error.InvalidDirection,
        }
        
        if (currentValue == 0) {
            step1Count += 1;
        }
    }
    
    print("day1 part1 optimized: {d}, part2 optimized: {d}\n", .{ step1Count, step2Count });
}

fn parseInstructions(file_lines: []const []const u8, dial_instructions: []DialInstruction) !void {
    for (file_lines, 0..) |line, i| {
        if (line.len == 0) continue;
        
        dial_instructions[i] = DialInstruction{
            .direction = line[0],
            .step = try std.fmt.parseInt(u32, line[1..], 10),
        };
    }
}