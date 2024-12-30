import sys
from functools import cache

from utils import input


@cache
def blink(stone: int, blinks: int):
    # print(f"{stone=}, {blinks=}")
    if blinks < 1:
        return [stone]
    string = str(stone)
    length = len(string)

    if stone == 0:
        return blink(1, blinks - 1)
    elif length % 2 == 0:
        return [
            *blink(int(string[: length // 2]), blinks - 1),
            *blink(int(string[length // 2 :]), blinks - 1),
        ]
    else:
        return blink(stone * 2024, blinks - 1)


def main():
    BLINKS = 75

    stones = [int(x) for x in input.lines()[0].split()]

    print(f"{stones=}")
    final_stones = []
    for stone in stones:
        print(f"OG STONE {stone=}")
        final_stones.extend(blink(stone, BLINKS))
        # print(f"{final_stones=}")

    print(f"{len(final_stones)} - after {BLINKS} blinks", file=sys.stderr)


if __name__ == "__main__":
    main()
