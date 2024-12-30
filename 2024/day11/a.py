import sys

from utils import input


def blink(stones: list[int]) -> list[int]:
    print(f"blinking on {stones}")
    next_stones = []
    for stone in stones:
        if stone == 0:
            next_stones.append(1)
            continue

        string = str(stone)
        length = len(string)

        if length % 2 == 0:
            next_stones.append(int(string[: length // 2]))
            next_stones.append(int(string[length // 2 :]))
        else:
            next_stones.append(stone * 2024)
        # print(f"{stone=}, {stones}")  # NOTE: uncommenting will let this run EXTREMELY slower

    return next_stones


def main():
    BLINKS = 25
    stones = [int(x) for x in input.lines()[0].split()]

    for b in range(BLINKS):
        print(f"BLINK = '{b}'", file=sys.stderr)
        stones = blink(stones)
    print(len(stones), file=sys.stderr)


if __name__ == "__main__":
    main()
