from itertools import product

from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
    grid = list(input.lines())
    cnt = 0
    for (row, col), letter in input.coords():
        if letter != "X":
            continue
        for dr, dc in product([-1, 0, 1], repeat=2):
            if dc == dr == 0:
                continue
            if not (
                0 <= row + 3 * dr < len(grid[row]) and 0 <= col + 3 * dc < len(grid)
            ):
                continue
            for i, l in enumerate("MAS", start=1):
                if grid[row + i * dr][col + i * dc] != l:
                    break
            else:
                cnt += 1
    return cnt


def part_two(input: Input) -> int | None:
    grid = list(input.lines())
    cnt = 0
    for (row, col), letter in input.coords():
        if row == 0 or row == len(grid) - 1 or col == 0 or col == len(grid[0]) - 1:
            continue

        if letter != "A":
            continue

        surround = "".join(
            grid[row + dr][col + dc]
            for (dr, dc) in [(-1, -1), (-1, 1), (1, 1), (1, -1)]
        )
        if surround in ["MMSS", "SMMS", "SSMM", "MSSM"]:
            cnt += 1
    return cnt


if __name__ == "__main__":
    assert (
        p1 := part_one(Input.from_file_relpath("example.txt"))
    ) == 18, f"Part one failed, {p1=}"
    assert (
        p2 := part_two(Input.from_file_relpath("example.txt"))
    ) == 9, f"Part two failed, {p2=}"
    Day(part_one, part_two).run()
