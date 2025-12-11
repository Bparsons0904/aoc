const std = @import("std");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const BatteryPack = []u32;

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    
    var batteryPacks = try allocator.alloc(BatteryPack, 16);
    var batteryPacks_count: usize = 0;
    
    try processDay3File(allocator, &batteryPacks, &batteryPacks_count);
    
    const part1Result = try calculate12PackJoltage(batteryPacks[0..batteryPacks_count], 2);
    const part2Result = try calculate12PackJoltage(batteryPacks[0..batteryPacks_count], 12);
    
    print("Results Part1: {d}, Part2: {d}\n", .{ part1Result, part2Result });
    
    // Free all allocated memory
    for (batteryPacks[0..batteryPacks_count]) |pack| {
        allocator.free(pack);
    }
    allocator.free(batteryPacks);
}

fn calculate12PackJoltage(batteryPacks: []const BatteryPack, batterySize: usize) !u64 {
    var maxJoltage: u64 = 0;
    for (batteryPacks) |batteryPack| {
        maxJoltage += try getLargestPackJoltage(batteryPack, batterySize);
    }
    return maxJoltage;
}

fn processDay3File(allocator: std.mem.Allocator, batteryPacks: *[]BatteryPack, batteryPacks_count: *usize) !void {
    const file = try file_parser.readFile(allocator, "day3.part1");
    defer allocator.free(file);
    
    for (file) |line| {
        var batteryPack = try allocator.alloc(u32, line.len);
        
        for (line, 0..) |batteryChar, i| {
            const battery = try std.fmt.charToDigit(batteryChar, 10);
            batteryPack[i] = @intCast(battery);
        }
        
        if (batteryPacks_count.* >= batteryPacks.len) {
            const new_packs = try allocator.realloc(batteryPacks.*, batteryPacks.len * 2);
            batteryPacks.* = new_packs;
        }
        batteryPacks.*[batteryPacks_count.*] = batteryPack;
        batteryPacks_count.* += 1;
    }
}

// Extension method for BatteryPack
fn getLargestPackJoltage(batteryPack: BatteryPack, batterySize: usize) !u64 {
    var maxJoltage = try std.heap.page_allocator.alloc(u32, batterySize);
    defer std.heap.page_allocator.free(maxJoltage);
    var maxJoltageCount: usize = 0;
    
    var index: usize = 0;
    var iteration: usize = 0;
    
    while (maxJoltageCount < batterySize) {
        const result = try iterateBatteryPackCheck(batteryPack, index, iteration, batterySize);
        maxJoltage[maxJoltageCount] = result.joltage;
        index = result.endIndex + 1;
        iteration += 1;
        maxJoltageCount += 1;
    }
    
    // Convert array of numbers to a single integer
    var result: u64 = 0;
    for (maxJoltage[0..maxJoltageCount]) |joltage| {
        var buf: [20]u8 = undefined;
        const joltageStr = std.fmt.bufPrint(&buf, "{d}", .{joltage}) catch continue;
        for (joltageStr) |digit| {
            const digitVal = try std.fmt.charToDigit(digit, 10);
            result = result * 10 + @as(u64, @intCast(digitVal));
        }
    }
    
    return result;
}

const JoltageResult = struct {
    joltage: u32,
    endIndex: usize,
};

fn iterateBatteryPackCheck(
    battery: BatteryPack,
    startingIndex: usize,
    currentBatterySize: usize,
    batterySize: usize,
) !JoltageResult {
    var maxFoundJoltage: u32 = 0;
    var endIndex: usize = startingIndex;
    const stop = battery.len + currentBatterySize - (batterySize - 1);
    
    var i = startingIndex;
    while (i < stop and i < battery.len) : (i += 1) {
        if (battery[i] > maxFoundJoltage) {
            maxFoundJoltage = battery[i];
            endIndex = i;
        }
    }
    
    return JoltageResult{
        .joltage = maxFoundJoltage,
        .endIndex = endIndex,
    };
}