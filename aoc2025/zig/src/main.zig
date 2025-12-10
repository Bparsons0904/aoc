const std = @import("std");
const challenges = @import("challenges.zig");

pub fn main() !void {
    const stdout = std.io.getStdErr().writer();
    
    // Get current time and add one hour (like in the Go version)
    const now = std.time.timestamp();
    const calendar_time = std.time.epoch.EpochSeconds{ .secs = now };
    const epoch_day = calendar_time.getEpochDay();
    const year_day = epoch_day.calculateYearDay();
    const day_of_month = @as(u5, @intCast(year_day.day_of_year % 31 + 1));
    
    switch (day_of_month) {
        1 => {
            try challenges.day1();
            try challenges.day1_1();
        },
        2 => try challenges.day2(),
        3 => try challenges.day3(),
        4 => try challenges.day4(),
        5 => try challenges.day5(),
        6 => try challenges.day6(),
        7 => try challenges.day7(),
        8 => try challenges.day8(),
        9 => try challenges.day9(),
        10 => try challenges.day10(),
        11 => try challenges.day11(),
        12 => try challenges.day12(),
        else => try stdout.print("No challenge for day {d}\n", .{day_of_month}),
    }
}