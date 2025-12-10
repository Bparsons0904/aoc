const std = @import("std");
const file_parser = @import("file_parser.zig");
const print = std.debug.print;

// Day 1 implementations
pub fn day1() !void {
    print("Running Day 1...\n");
    
    // Create a stub implementation since day1.part1 file doesn't exist
    const example_instructions = [_][]const u8{ "R4", "L2", "R3" };
    
    // Create linked list representation
    var dial_list = DialList.init();
    
    // Position at 50 like in the Go version
    while (dial_list.head.value != 50) {
        dial_list.next();
    }
    
    var day1part1Count: u32 = 0;
    var day1part2Count: u32 = 0;
    
    for (example_instructions) |instruction| {
        const direction = instruction[0];
        const step = try std.fmt.parseInt(u32, instruction[1..], 10);
        
        switch (direction) {
            'R' => {
                for (0..step) |_| {
                    dial_list.next();
                    if (dial_list.head.value == 0) {
                        day1part2Count += 1;
                    }
                }
            },
            'L' => {
                for (0..step) |_| {
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
    
    print("Day 1 Part 1: {d}, Part 2: {d}\n", .{ day1part1Count, day1part2Count });
}

pub fn day1_1() !void {
    print("Running Day 1 Part 1 optimized...\n");
    const example_instructions = [_][]const u8{ "R4", "L2", "R3" };
    
    var step1Count: u32 = 0;
    var step2Count: u32 = 0;
    var currentValue: i32 = 50;
    
    for (example_instructions) |instruction| {
        const direction = instruction[0];
        const step = try std.fmt.parseInt(i32, instruction[1..], 10);
        
        switch (direction) {
            'R' => {
                const newValue = currentValue + step;
                if (newValue >= 100) {
                    if (currentValue == 0) {
                        step2Count += @intCast(step / 100);
                    } else {
                        step2Count += @intCast((step + currentValue) / 100);
                    }
                }
                currentValue = newValue % 100;
            },
            'L' => {
                const newValue = currentValue - step;
                if (currentValue == 0) {
                    step2Count += @intCast(step / 100);
                } else if (step >= currentValue) {
                    step2Count += @intCast((step - currentValue) / 100 + 1);
                }
                currentValue = ((newValue % 100) + 100) % 100;
            },
            else => return error.InvalidDirection,
        }
        
        if (currentValue == 0) {
            step1Count += 1;
        }
    }
    
    print("Day 1 Part 1 Optimized: {d}, Part 2 Optimized: {d}\n", .{ step1Count, step2Count });
}

// Day 2 implementation
pub fn day2() !void {
    print("Running Day 2...\n");
    
    const allocator = std.heap.page_allocator;
    
    // Read the input file
    const lines = try file_parser.readFile(allocator, "day2.part1");
    defer allocator.free(lines);
    
    if (lines.len == 0) {
        print("No input found in day2.part1\n");
        return;
    }
    
    const row = lines[0];
    
    // Parse the ranges
    var ranges = std.mem.tokenize(u8, row, ",");
    var product_id_ranges = std.ArrayList(ProductIDRange).init(allocator);
    defer product_id_ranges.deinit();
    
    while (ranges.next()) |range_str| {
        var parts = std.mem.tokenize(u8, range_str, "-");
        const min_str = parts.next() orelse return error.InvalidInput;
        const max_str = parts.next() orelse return error.InvalidInput;
        
        const min = try std.fmt.parseInt(u32, min_str, 10);
        const max = try std.fmt.parseInt(u32, max_str, 10);
        
        try product_id_ranges.append(ProductIDRange{ .min = min, .max = max });
    }
    
    const part1Count = calculatePart1(product_id_ranges.items);
    const part2Count = calculatePart2(product_id_ranges.items);
    const part2_1Count = calculatePart2_1(product_id_ranges.items);
    const part2_2Count = calculatePart2_2(product_id_ranges.items);
    const part2_3Count = calculatePart2_3(product_id_ranges.items);
    
    print("Day 2 - Part1: {d}, Part2: {d}, Part2_1: {d}, Part2_2: {d}, Part2_3: {d}\n", 
        .{ part1Count, part2Count, part2_1Count, part2_2Count, part2_3Count });
}

// Placeholder for other days
pub fn day3() !void {
    print("Day 3 not implemented yet\n");
}

pub fn day4() !void {
    print("Day 4 not implemented yet\n");
}

pub fn day5() !void {
    print("Day 5 not implemented yet\n");
}

pub fn day6() !void {
    print("Day 6 not implemented yet\n");
}

pub fn day7() !void {
    print("Day 7 not implemented yet\n");
}

pub fn day8() !void {
    print("Day 8 not implemented yet\n");
}

pub fn day9() !void {
    print("Day 9 not implemented yet\n");
}

pub fn day10() !void {
    print("Day 10 not implemented yet\n");
}

pub fn day11() !void {
    print("Day 11 not implemented yet\n");
}

pub fn day12() !void {
    print("Day 12 not implemented yet\n");
}

// Helper functions for Day 1
const Dial = struct {
    value: u32,
    next: ?*Dial,
    prev: ?*Dial,
};

const DialList = struct {
    head: *Dial,
    tail: *Dial,
    
    fn init() DialList {
        const allocator = std.heap.page_allocator;
        
        // Create 100 nodes in a circular linked list
        var nodes: [100]*Dial = undefined;
        for (0..100) |i| {
            const node = allocator.create(Dial) catch @panic("OOM");
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

// Helper functions for Day 2
const ProductIDRange = struct {
    min: u32,
    max: u32,
};

fn calculatePart1(product_id_ranges: []const ProductIDRange) u32 {
    var result: u32 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = try std.fmt.bufPrint(&buf, "{d}", .{i});
            
            if (i_str.len % 2 != 0) continue;
            
            const half = i_str.len / 2;
            if (std.mem.eql(u8, i_str[0..half], i_str[half..])) {
                result += i;
            }
        }
    }
    
    return result;
}

fn calculatePart2(product_id_ranges: []const ProductIDRange) u32 {
    var result: u32 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = try std.fmt.bufPrint(&buf, "{d}", .{i});
            
            var j: usize = 0;
            while (j <= i_str.len - 2) : (j += 1) {
                const to_check = i_str[0..j+1];
                const expected_count = (i_str.len + to_check.len - 1) / to_check.len;
                var count: usize = 0;
                var pos: usize = 0;
                
                while (pos < i_str.len) {
                    if (pos + to_check.len <= i_str.len and 
                        std.mem.eql(u8, i_str[pos..pos+to_check.len], to_check)) {
                        count += 1;
                        pos += to_check.len;
                    } else {
                        break;
                    }
                }
                
                if (count >= 2 and count == expected_count) {
                    result += i;
                    break;
                }
            }
        }
    }
    
    return result;
}

// Combo Bob and Derek implementation
fn calculatePart2_1(product_id_ranges: []const ProductIDRange) u32 {
    var result: u32 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = try std.fmt.bufPrint(&buf, "{d}", .{i});
            
            var j: usize = 1;
            while (j <= i_str.len / 2) : (j += 1) {
                const pattern = i_str[0..j];
                var remaining = i_str[j..];
                
                var all_match = true;
                while (remaining.len > 0) {
                    if (remaining.len < pattern.len or 
                        !std.mem.eql(u8, remaining[0..pattern.len], pattern)) {
                        all_match = false;
                        break;
                    }
                    remaining = remaining[pattern.len..];
                }
                
                if (all_match) {
                    result += i;
                    break;
                }
            }
        }
    }
    
    return result;
}

// Claude implementation
fn calculatePart2_2(product_id_ranges: []const ProductIDRange) u32 {
    var result: u32 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = try std.fmt.bufPrint(&buf, "{d}", .{i});
            const str_len = i_str.len;
            
            var pattern_len: usize = 1;
            while (pattern_len <= str_len / 2) : (pattern_len += 1) {
                if (str_len % pattern_len != 0) continue;
                
                const pattern = i_str[0..pattern_len];
                var is_repeating = true;
                
                var offset: usize = pattern_len;
                while (offset < str_len) : (offset += pattern_len) {
                    if (!std.mem.eql(u8, i_str[offset..offset+pattern_len], pattern)) {
                        is_repeating = false;
                        break;
                    }
                }
                
                if (is_repeating) {
                    result += i;
                    break;
                }
            }
        }
    }
    
    return result;
}

// Geminis implementation
fn calculatePart2_3(product_id_ranges: []const ProductIDRange) u32 {
    var result: u32 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = try std.fmt.bufPrint(&buf, "{d}", .{i});
            const n = i_str.len;
            
            if (n < 2) continue;
            
            // Create s+s and search for s starting from index 1
            var doubled = try std.heap.page_allocator.alloc(u8, n * 2);
            defer std.heap.page_allocator.free(doubled);
            
            std.mem.copy(u8, doubled[0..n], i_str);
            std.mem.copy(u8, doubled[n..n*2], i_str);
            
            // Search for the pattern starting from index 1
            if (std.mem.indexOf(u8, doubled[1..], i_str)) |idx| {
                const period = idx + 1;
                if (period > 0 and period < n) {
                    result += i;
                }
            }
        }
    }
    
    return result;
}