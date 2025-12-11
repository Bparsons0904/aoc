const std = @import("std");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const ProductIDRange = struct {
    min: u64,
    max: u64,
};

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    
    const productIDRanges = try processDay2(allocator);
    defer allocator.free(productIDRanges);
    
    const part1Count = calculatePart1(productIDRanges);
    const part2Count = calculatePart2(productIDRanges);
    const part2_1Count = calculatePart2_1(productIDRanges);
    const part2_2Count = calculatePart2_2(productIDRanges);
    const part2_3Count = calculatePart2_3(productIDRanges);
    
    print("Day 2 - Part1: {d}, Part2: {d}, Part2_1: {d}, Part2_2: {d}, Part2_3: {d}\n", 
        .{ part1Count, part2Count, part2_1Count, part2_2Count, part2_3Count });
}

fn processDay2(allocator: std.mem.Allocator) ![]ProductIDRange {
    const lines = try file_parser.readFile(allocator, "day2.part1");
    defer allocator.free(lines);
    
    if (lines.len == 0) {
        return error.NoInput;
    }
    
    const row = lines[0];
    
    // Parse the ranges - count them first
    var range_count: usize = 1;
    for (row) |byte| {
        if (byte == ',') range_count += 1;
    }
    
    var ranges = try allocator.alloc(ProductIDRange, range_count);
    var ranges_idx: usize = 0;
    
    var start: usize = 0;
    for (row, 0..) |byte, i| {
        if (byte == ',') {
            const range_str = row[start..i];
            
            var parts_start: usize = start;
            var min_end: ?usize = null;
            
            // Find the dash separator
            for (range_str, 0..) |c, j| {
                if (c == '-') {
                    min_end = start + j;
                    parts_start = start + j + 1;
                    break;
                }
            }
            
            if (min_end) |end| {
                const min_str = row[start..end];
                const max_str = row[parts_start..i];
                
                const min = try std.fmt.parseInt(u64, min_str, 10);
                const max = try std.fmt.parseInt(u64, max_str, 10);
                
                ranges[ranges_idx] = ProductIDRange{ .min = min, .max = max };
                ranges_idx += 1;
            }
            
            start = i + 1;
        }
    }
    
    // Handle the last range
    const range_str = row[start..];
    var parts_start: usize = start;
    var min_end: ?usize = null;
    
    // Find the dash separator
    for (range_str, 0..) |c, j| {
        if (c == '-') {
            min_end = start + j;
            parts_start = start + j + 1;
            break;
        }
    }
    
    if (min_end) |end| {
        const min_str = row[start..end];
        const max_str = row[parts_start..];
        
        const min = try std.fmt.parseInt(u64, min_str, 10);
        const max = try std.fmt.parseInt(u64, max_str, 10);
        
        ranges[ranges_idx] = ProductIDRange{ .min = min, .max = max };
        ranges_idx += 1;
    }
    
    return ranges;
}

fn calculatePart1(product_id_ranges: []const ProductIDRange) u64 {
    var result: u64 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = std.fmt.bufPrint(&buf, "{d}", .{i}) catch continue;
            
            if (i_str.len % 2 != 0) continue;
            
            const half = i_str.len / 2;
            if (std.mem.eql(u8, i_str[0..half], i_str[half..])) {
                result += i;
            }
        }
    }
    
    return result;
}

// My implementation
fn calculatePart2(product_id_ranges: []const ProductIDRange) u64 {
    var result: u64 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = std.fmt.bufPrint(&buf, "{d}", .{i}) catch continue;
            
            var j: usize = 0;
            while (j + 1 < i_str.len) : (j += 1) {
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
fn calculatePart2_1(product_id_ranges: []const ProductIDRange) u64 {
    var result: u64 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = std.fmt.bufPrint(&buf, "{d}", .{i}) catch continue;
            
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
fn calculatePart2_2(product_id_ranges: []const ProductIDRange) u64 {
    var result: u64 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = std.fmt.bufPrint(&buf, "{d}", .{i}) catch continue;
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
fn calculatePart2_3(product_id_ranges: []const ProductIDRange) u64 {
    var result: u64 = 0;
    
    for (product_id_ranges) |range| {
        var i = range.min;
        while (i <= range.max) : (i += 1) {
            var buf: [20]u8 = undefined;
            const i_str = std.fmt.bufPrint(&buf, "{d}", .{i}) catch continue;
            const n = i_str.len;
            
            if (n < 2) continue;
            
            // Create s+s and search for s starting from index 1
            // Simplified approach without alloc
            var period: usize = 1;
            while (period < n) : (period += 1) {
                if (n % period != 0) continue;
                
                const pattern = i_str[0..period];
                var is_repeating = true;
                
                var offset: usize = period;
                while (offset < n) : (offset += period) {
                    if (!std.mem.eql(u8, i_str[offset..offset+period], pattern)) {
                        is_repeating = false;
                        break;
                    }
                }
                
                if (is_repeating and period < n) {
                    result += i;
                    break;
                }
            }
        }
    }
    
    return result;
}