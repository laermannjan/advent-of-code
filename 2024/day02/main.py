from itertools import pairwise

from utils.execute import Day
from utils.input import Input


def is_safe(report: list[int]) -> bool:
    diffs = [y - x for x, y in pairwise(report)]
    return all(1 <= d <= 3 for d in diffs) or all(-1 >= d >= -3 for d in diffs)


def part_one(input: Input) -> int | None:
    reports = [list(map(int, line.split())) for line in input.lines()]
    return sum([is_safe(report) for report in reports])


def part_two(input: Input) -> int | None:
    reports = [list(map(int, line.split())) for line in input.lines()]
    # NOTE: safe reports are still safe when skipping first/last element
    # so we can skip checking the entire report
    return sum(
        any(
            is_safe(report[:index] + report[index + 1 :])
            for index in range(len(report))
        )
        for report in reports
    )


## Above is a simple and straightforward solution. Depending on the input's size it's likely slower,
## but much easier to write, understand, and maintain.
##########################
## Below checks each report entirely, but realizes that we only need to
## try remove 3 indices when encountering a problem.
## If the index pair (i, i+1) make a report unsafe, we can determine that removing either i-1, i, or i+1 must make it safe
## removing any other index would not affect this problem encountered here


# def has_problem(report: list[int]) -> int | None:
#     """Returns None if report is safe. Index of the left element of the pair that resulted in a problem otherwise."""
#     increasing = decreasing = True
#     for i, (prev, next) in enumerate(pairwise(report)):
#         if not (1 <= abs(next - prev) <= 3):
#             print("step too large")
#             return i
#
#         if next < prev:
#             increasing = False
#         if next > prev:
#             decreasing = False
#         if not increasing and not decreasing:
#             print("not monotone")
#             return i
#     return None
#
#
# def part_one(input: Input) -> int | None:
#     reports = [list(map(int, line.split())) for line in input.lines()]
#     return sum([has_problem(report) is None for report in reports])
#
#
# def part_two(input: Input) -> int | None:
#     reports = [list(map(int, line.split())) for line in input.lines()]
#     safe = 0
#     for r, report in enumerate(reports):
#         idx = has_problem(report)
#         if idx is None:
#             safe += 1
#             print(f"report {r+1} safe")
#         else:
#             print(f"report {r+1} not safe, checking tolerance")
#             for i in [idx - 1, idx, idx + 1]:
#                 if i < 0:
#                     continue
#                 new_report = report[:i] + report[i + 1 :]
#                 if has_problem(new_report) is None:
#                     safe += 1
#                     print(f"  removing index {i} makes report safe")
#                     break
#             print("  report not safe")
#     return safe


if __name__ == "__main__":
    Day(part_one, part_two).run()
