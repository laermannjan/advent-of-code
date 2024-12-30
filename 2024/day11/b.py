import sys
from functools import cache

from utils import input


@cache
def count(stone: int, blinks: int) -> int:
    # NOTE: we need to compute the count of stones
    # instead of the actual stones, otherwise the list of stones
    # grows too large and bottlenecks

    # print(f"{stone=}, {blinks=}")
    if blinks == 0:
        return 1
    if stone == 0:
        return count(1, blinks - 1)

    string = str(stone)
    length = len(string)
    if length % 2 == 0:
        c1 = count(int(string[: length // 2]), blinks - 1)
        c2 = count(int(string[length // 2 :]), blinks - 1)
        return c1 + c2
    return count(stone * 2024, blinks - 1)


def main():
    BLINKS = 75

    stones = [int(x) for x in input.lines()[0].split()]

    print(f"{stones=}")
    n_stones = 0
    for stone in stones:
        print(f"OG STONE {stone=}")
        n_stones += count(stone, BLINKS)

    print(f"{n_stones} - after {BLINKS} blinks", file=sys.stderr)


if __name__ == "__main__":
    main()
