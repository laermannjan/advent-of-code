from itertools import pairwise

from utils.execute import Day
from utils.input import Input


def has_problem(report: list[int]) -> int | None:
    """Returns None if report is safe. Index of the left element of the pair that resulted in a problem otherwise."""
    increasing = decreasing = True
    for i, (prev, next) in enumerate(pairwise(report)):
        if not (1 <= abs(next - prev) <= 3):
            print("step too large")
            return i

        if next < prev:
            increasing = False
        if next > prev:
            decreasing = False
        if not increasing and not decreasing:
            print("not monotone")
            return i
    return None


def part_one(input: Input) -> int | None:
    reports = [list(map(int, line.split())) for line in input.lines()]
    return sum([has_problem(report) is None for report in reports])


def part_two(input: Input) -> int | None:
    reports = [list(map(int, line.split())) for line in input.lines()]
    safe = 0
    for r, report in enumerate(reports):
        idx = has_problem(report)
        if idx is None:
            safe += 1
            print(f"report {r+1} safe")
        else:
            print(f"report {r+1} not safe, checking tolerance")
            for i in [idx - 1, idx, idx + 1]:
                if i < 0:
                    continue
                new_report = report[:i] + report[i + 1 :]
                if has_problem(new_report) is None:
                    safe += 1
                    print(f"  removing index {i} makes report safe")
                    break
            print("  report not safe")
    return safe


if __name__ == "__main__":
    Day(part_one, part_two).run()
