import sys
from itertools import pairwise

from utils import input


def is_safe(report: list[int]) -> bool:
    diffs = [y - x for x, y in pairwise(report)]
    return all(1 <= d <= 3 for d in diffs) or all(-1 >= d >= -3 for d in diffs)


reports = [list(map(int, line.split())) for line in input.lines()]
# NOTE: safe reports are still safe when skipping first/last element
# so we can skip checking the entire report
safe_reports = [
    any(is_safe(report[:index] + report[index + 1 :]) for index in range(len(report)))
    for report in reports
]
print(sum(safe_reports), file=sys.stderr)
