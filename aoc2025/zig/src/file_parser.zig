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
    
    // Simple array list implementation for compatibility
    var lines = try allocator.alloc([]const u8, 128);
    var lines_count: usize = 0;
    
    var start: usize = 0;
    for (buffer, 0..) |byte, i| {
        if (byte == '\n') {
            const line = buffer[start..i];
            if (line.len > 0 and line[line.len-1] == '\r') {
                const trimmed = line[0..line.len-1];
                if (trimmed.len > 0) {
                    const owned_line = try allocator.dupe(u8, trimmed);
                    if (lines_count >= lines.len) {
                        const new_lines = try allocator.realloc(lines, lines.len * 2);
                        lines = new_lines;
                    }
                    lines[lines_count] = owned_line;
                    lines_count += 1;
                }
            } else if (line.len > 0) {
                const owned_line = try allocator.dupe(u8, line);
                if (lines_count >= lines.len) {
                    const new_lines = try allocator.realloc(lines, lines.len * 2);
                    lines = new_lines;
                }
                lines[lines_count] = owned_line;
                lines_count += 1;
            }
            start = i + 1;
        }
    }
    
    // Handle last line if it doesn't end with newline
    if (start < buffer.len) {
        const line = buffer[start..];
        if (line.len > 0) {
            const owned_line = try allocator.dupe(u8, line);
            if (lines_count >= lines.len) {
                const new_lines = try allocator.realloc(lines, lines.len * 2);
                lines = new_lines;
            }
            lines[lines_count] = owned_line;
            lines_count += 1;
        }
    }
    
    return allocator.realloc(lines, lines_count);
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