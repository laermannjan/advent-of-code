import sys
from functools import cache

from utils import input


@cache
def blink(stone: str, blinks: int):
    next_stones: list[str | list[str]]

    if stone == "0":
        next_stones = ["1"]
    elif len(stone) % 2 == 0:
        next_stones = [stone[: len(stone) // 2], str(int(stone[len(stone) // 2 :]))]
    else:
        next_stones = [str(int(stone) * 2024)]

    print(f"computing {stone=} - {blinks=} - {next_stones=}")
    blinked_stones = []
    for next_stone in next_stones:
        blinked_stones.extend(
            blink(next_stone, blinks - 1) if blinks > 1 else [next_stone]
        )

    print(f"returning {stone=} - {blinks=} - {blinked_stones=}")
    return blinked_stones


def main():
    BLINKS = 6
    stones = input.lines()[0].split(" ")

    print(f"{stones=}")
    final_stones = []
    for s, stone in enumerate(stones):
        print(f"OG STONE {stone=}")
        final_stones.extend(blink(stone, BLINKS))

    print(f"{final_stones=}")
    print(f"{len(final_stones)} - after {BLINKS} blinks", file=sys.stderr)


if __name__ == "__main__":
    main()
