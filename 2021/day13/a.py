import sys

from utils import input


def parse_input() -> tuple[set[tuple[int, int]], list[tuple[str, int]]]:
    dot_data, instructions_data = input.stdin().read().split("\n\n")

    dots: set[tuple[int, int]] = {
        (int(x), int(y)) for x, y in (line.split(",") for line in dot_data.splitlines())
    }

    folds: list[tuple[str, int]] = [
        (axis, int(index))
        for axis, index in (
            instruction.split("fold along ")[1].split("=")
            for instruction in instructions_data.splitlines()
        )
    ]

    return dots, folds


def fold(dots: set[tuple[int, int]], axis: str, index: int) -> set[tuple[int, int]]:
    if axis == "x":
        dots = {(2 * index - dx, dy) if dx > index else (dx, dy) for dx, dy in dots}
    else:
        assert axis == "y"
        dots = {(dx, 2 * index - dy) if dy > index else (dx, dy) for dx, dy in dots}
    return dots


dots, folds = parse_input()
axis, index = folds[0]
dots = fold(dots, axis, index)
print(len(dots), file=sys.stderr)
