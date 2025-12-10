const std = @import("std");

pub fn readFile(allocator: std.mem.Allocator, file_name: []const u8) ![][]const u8 {
    const file_path = try std.fmt.allocPrint(allocator, "files/{s}", .{file_name});
    defer allocator.free(file_path);
    
    const file = std.fs.cwd().openFile(file_path, .{}) catch |err| switch (err) {
        error.FileNotFound => {
            std.log.err("File not found: {s}", .{file_path});
            return err;
        },
        else => return err,
    };
    defer file.close();
    
    const buffer = try file.readToEndAlloc(allocator, 1024 * 1024);
    defer allocator.free(buffer);
    
    var lines = std.ArrayList([]const u8).init(allocator);
    defer lines.deinit();
    
    var iter = std.mem.split(u8, buffer, "\n");
    while (iter.next()) |line| {
        const trimmed = std.mem.trim(u8, line, "\r");
        if (trimmed.len > 0) {
            const owned_line = try allocator.dupe(u8, trimmed);
            try lines.append(owned_line);
        }
    }
    
    return lines.toOwnedSlice();
}

pub fn readFileToOwnedSlice(allocator: std.mem.Allocator, file_name: []const u8) ![]u8 {
    const file_path = try std.fmt.allocPrint(allocator, "files/{s}", .{file_name});
    defer allocator.free(file_path);
    
    const file = std.fs.cwd().openFile(file_path, .{}) catch |err| switch (err) {
        error.FileNotFound => {
            std.log.err("File not found: {s}", .{file_path});
            return err;
        },
        else => return err,
    };
    defer file.close();
    
    return file.readToEndAlloc(allocator, 1024 * 1024);
}