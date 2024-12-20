import sys

from utils import input


def count_increases(numbers: list[int]) -> int:
    return sum([this < next for this, next in zip(numbers, numbers[1:])])


def window_sums(numbers: list[int], k: int) -> list[int]:
    return [sum(numbers[i : i + 3]) for i in range(len(numbers) - 2)]


print(
    count_increases(window_sums([int(line) for line in input.lines()], k=3)),
    file=sys.stderr,
)
