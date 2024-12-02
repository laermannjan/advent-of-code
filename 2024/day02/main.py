from itertools import pairwise

from utils.execute import Day
from utils.input import Input


def is_safe(report: list[int]) -> bool:
    increasing = decreasing = True
    for prev, next in pairwise(report):
        if not (1 <= abs(next - prev) <= 3):
            print("step too large")
            return False

        if next < prev:
            increasing = False
        if next > prev:
            decreasing = False
        if not increasing and not decreasing:
            print("not monotone")
            return False
    return True


def part_one(input: Input) -> int | None:
    reports = [list(map(int, line.split())) for line in input.lines()]
    return sum([is_safe(report) for report in reports])


def part_two(input: Input) -> int | None:
    pass


if __name__ == "__main__":
    Day(part_one, part_two).run()
