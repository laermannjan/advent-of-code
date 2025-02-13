const std = @import("std");

fn solve(input: []const u8, allocator: std.mem.Allocator) !u32 {
    var arena = std.heap.ArenaAllocator.init(allocator);
    defer arena.deinit();

    var l1 = std.ArrayList(u32).init(arena.allocator());
    var counts = std.AutoHashMap(u32, u32).init(arena.allocator());

    var lines = std.mem.tokenizeScalar(u8, input, '\n');
    while (lines.next()) |line| {
        var nums = std.mem.tokenizeScalar(u8, line, ' ');
        try l1.append(try std.fmt.parseUnsigned(u32, nums.next().?, 10));
        const v = try std.fmt.parseUnsigned(u32, nums.next().?, 10);
        try counts.put(v, (counts.get(v) orelse 0) + 1);
    }
    var sum: u32 = 0;
    for (l1.items) |a| {
        sum += a * (counts.get(a) orelse 0);
    }
    return sum;
}

pub fn main() !void {
    const stdin = std.io.getStdIn();
    const input = try stdin.reader().readAllAlloc(std.heap.page_allocator, 1024 * 1024); // max 1MB
    defer std.heap.page_allocator.free(input);
    std.debug.print("{!}", .{solve(input, std.heap.page_allocator)});
}

test "example" {
    const input = @embedFile("example.txt");
    try std.testing.expect(try solve(input, std.heap.page_allocator) == 11);
}
