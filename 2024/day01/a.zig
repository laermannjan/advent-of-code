const std = @import("std");

fn solve(input: []const u8, allocator: std.mem.Allocator) !u32 {
    var arena = std.heap.ArenaAllocator.init(allocator);
    defer arena.deinit();

    var l1 = std.ArrayList(i32).init(arena.allocator());
    var l2 = std.ArrayList(i32).init(arena.allocator());

    var lines = std.mem.tokenizeScalar(u8, input, '\n');
    while (lines.next()) |line| {
        var nums = std.mem.tokenizeScalar(u8, line, ' ');
        try l1.append(try std.fmt.parseUnsigned(i32, nums.next().?, 10));
        try l2.append(try std.fmt.parseUnsigned(i32, nums.next().?, 10));
    }

    std.mem.sort(i32, l1.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, l2.items, {}, comptime std.sort.asc(i32));

    var sum: u32 = 0;
    for (l1.items, l2.items) |a, b| {
        sum += @abs(a - b);
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
