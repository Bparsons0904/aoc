const std = @import("std");
const file_parser = @import("../file_parser.zig");
const print = std.debug.print;

const IngredientRange = struct {
    min: u32,
    max: u32,
};

const Ingredients = struct {
    freshIngredientRanges: []IngredientRange,
    ingredients: []u32,
    allocator: std.mem.Allocator,
    
    fn init(allocator: std.mem.Allocator) Ingredients {
        return Ingredients{
            .freshIngredientRanges = &[_]IngredientRange{},
            .ingredients = &[_]u32{},
            .allocator = allocator,
        };
    }
    
    fn deinit(self: *Ingredients) void {
        self.allocator.free(self.freshIngredientRanges);
        self.allocator.free(self.ingredients);
    }
};

pub fn run() !void {
    const allocator = std.heap.page_allocator;
    
    var ingredients = Ingredients.init(allocator);
    defer ingredients.deinit();
    
    const part1Count, const part2Count = countFreshIngredients(ingredients);
    
    print("Part 1: {d}, Part 2: {d}\n", .{ part1Count, part2Count });
}

fn countFreshIngredients(ingredients: Ingredients) struct { u32, u32 } {
    _ = ingredients;
    return .{ 0, 0 };
}