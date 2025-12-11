const std = @import("std");
const file_parser = @import("file_parser.zig");

pub const EMPTY = '.';
pub const PAPER_ROLL = '@';
pub const MOVED_UP = '^';
pub const MOVED_DOWN = 'v';
pub const MOVED_LEFT = '<';
pub const MOVED_RIGHT = '>';
pub const START = 'S';
pub const TACHYON = MOVED_UP;

pub const Point = struct {
    x: i32,
    y: i32,
    
    pub fn init(x: i32, y: i32) Point {
        return Point{ .x = x, .y = y };
    }
};

pub const RIGHT = Point{ .x = 1, .y = 0 };
pub const LEFT = Point{ .x = -1, .y = 0 };
pub const DOWN = Point{ .x = 0, .y = 1 };
pub const UP = Point{ .x = 0, .y = -1 };
pub const RIGHT_DOWN = Point{ .x = 1, .y = 1 };
pub const RIGHT_UP = Point{ .x = 1, .y = -1 };
pub const LEFT_DOWN = Point{ .x = -1, .y = 1 };
pub const LEFT_UP = Point{ .x = -1, .y = -1 };

pub const DIRECTIONS = [_]Point{
    LEFT, RIGHT, UP, DOWN, LEFT_DOWN, LEFT_UP, RIGHT_DOWN, RIGHT_UP,
};

pub const Visit = struct {
    point: Point,
    direction: Point,
};

pub const Grid = struct {
    width: usize,
    height: usize,
    start: Point,
    current: Point,
    visited: []Visit,
    visited_count: usize,
    visited_capacity: usize,
    map: [][]u8,
    allocator: std.mem.Allocator,
    
    pub fn init(allocator: std.mem.Allocator, filename: []const u8) !Grid {
        const lines = try file_parser.readFile(allocator, filename);
        defer allocator.free(lines);
        
        if (lines.len == 0) {
            return error.EmptyFile;
        }
        
        const width = lines[0].len;
        const height = lines.len;
        
        var map = try allocator.alloc([]u8, height);
        for (0..height) |y| {
            map[y] = try allocator.alloc(u8, width);
            for (0..width) |x| {
                map[y][x] = lines[y][x];
            }
        }
        
        var start = Point{ .x = 0, .y = 0 };
        for (0..height) |y| {
            for (0..width) |x| {
                if (map[y][x] == START) {
                    start = Point{ .x = @intCast(x), .y = @intCast(y) };
                    break;
                }
            }
        }
        
        const visited_capacity = 16;
        var visited = try allocator.alloc(Visit, visited_capacity);
        visited[0] = Visit{ .point = start, .direction = Point{ .x = 0, .y = 0 } };
        
        return Grid{
            .width = width,
            .height = height,
            .start = start,
            .current = start,
            .visited = visited,
            .visited_count = 1,
            .visited_capacity = visited_capacity,
            .map = map,
            .allocator = allocator,
        };
    }
    
    pub fn deinit(self: *Grid) void {
        for (self.map) |row| {
            self.allocator.free(row);
        }
        self.allocator.free(self.map);
        self.allocator.free(self.visited);
    }
    
    pub fn setStart(self: *Grid, point: Point) void {
        self.start = point;
        self.current = point;
        self.visited_count = 0;
        self.visited[0] = Visit{ .point = point, .direction = Point{ .x = 0, .y = 0 } };
        self.visited_count = 1;
    }
    
    pub fn setObject(self: *Grid, point: Point, object: u8) void {
        if (self.pointWithinBounds(point)) {
            self.map[@as(usize, @intCast(point.y))][@as(usize, @intCast(point.x))] = object;
        }
    }
    
    pub fn positionContainsObject(self: *Grid, point: Point, object: u8) bool {
        if (!self.pointWithinBounds(point)) {
            return false;
        }
        return self.map[@as(usize, @intCast(point.y))][@as(usize, @intCast(point.x))] == object;
    }
    
    pub fn pointWithinBounds(self: *Grid, point: Point) bool {
        return point.x >= 0 and point.x < @as(i32, @intCast(self.width)) and
               point.y >= 0 and point.y < @as(i32, @intCast(self.height));
    }
    
    pub fn canMove(self: *Grid, direction: Point) bool {
        const newPoint = Point{
            .x = self.current.x + direction.x,
            .y = self.current.y + direction.y,
        };
        return self.pointWithinBounds(newPoint);
    }
    
    pub fn move(self: *Grid, direction: Point) !bool {
        if (!self.canMove(direction)) {
            return false;
        }
        
        const newPoint = Point{
            .x = self.current.x + direction.x,
            .y = self.current.y + direction.y,
        };
        
        const visit = Visit{
            .point = newPoint,
            .direction = direction,
        };
        
        if (self.visited_count >= self.visited_capacity) {
            const new_visited = try self.allocator.realloc(self.visited, self.visited_capacity * 2);
            self.visited = new_visited;
            self.visited_capacity *= 2;
        }
        
        self.visited[self.visited_count] = visit;
        self.visited_count += 1;
        self.current = newPoint;
        
        return true;
    }
    
    pub fn print(self: *Grid) void {
        for (0..self.height) |y| {
            for (0..self.width) |x| {
                std.debug.print("{c}", .{self.map[y][x]});
            }
            std.debug.print("\n");
        }
    }
};