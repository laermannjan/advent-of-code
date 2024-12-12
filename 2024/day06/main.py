from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
    moves = {"^": (-1, 0), ">": (0, 1), "v": (1, 0), "<": (0, -1)}
    grid = list(input.lines())

    def in_bounds(row, col):
        return 0 <= row < len(grid) and 0 <= col < len(grid[0])

    def next_pos(row, col, direction):
        return row + moves[direction][0], col + moves[direction][1]

    for row in range(len(grid)):
        for col in range(len(grid[row])):
            ele = grid[row][col]
            if ele in ("^", "v", "<", ">"):
                guard_row, guard_col = (row, col)
                direction = ele
    else:
        RuntimeError("guard not found")

    positions = {(guard_row, guard_col)}

    while True:
        next_row_maybe, next_col_maybe = next_pos(guard_row, guard_col, direction)
        print(f"looking at ({next_row_maybe}, {next_col_maybe})")

        if not in_bounds(next_row_maybe, next_col_maybe):
            print(f"({next_row_maybe}, {next_col_maybe}) is out of bounds")
            break

        if grid[next_row_maybe][next_col_maybe] == "#":
            print(f"({next_row_maybe}, {next_col_maybe}) is #")
            if direction == "^":
                print("turning >")
                direction = ">"
            elif direction == ">":
                print("turning v")
                direction = "v"
            elif direction == "v":
                print("turning <")
                direction = "<"
            elif direction == "<":
                print("turning ^")
                direction = "^"
            continue
        guard_row, guard_col = next_row_maybe, next_col_maybe
        positions.add((guard_row, guard_col))
        # sleep(0.5)

    return len(positions)


def part_two(input: Input) -> int | None:
    pass


if __name__ == "__main__":
    assert (
        p1 := part_one(Input.from_file_relpath("example.txt"))
    ) == 41, f"Part one failed, {p1=}"
    assert (
        p2 := part_two(Input.from_file_relpath("example.txt"))
    ) == None, f"Part two failed, {p2=}"
    Day(part_one, part_two).run()
