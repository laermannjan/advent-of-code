import sys

from utils import input


def blink(stones: list[str]) -> list[str]:
    print(f"blinking on {stones}")
    it = enumerate(stones)
    for s, stone in it:
        if stone == "0":
            stones[s] = "1"
        elif len(stone) % 2 == 0:
            stones[s] = stone[: len(stone) // 2]
            stones.insert(s + 1, str(int(stone[len(stone) // 2 :])))
            next(it)
        else:
            stones[s] = str(int(stone) * 2024)
        print(f"{s=}, {stone=}, {stones}")

    return stones


def main():
    BLINKS = 25
    stones = input.lines()[0].split(" ")

    for b in range(BLINKS):
        print(f"BLINK = '{b}'", file=sys.stderr)
        stones = blink(stones)
    print(len(stones), file=sys.stderr)


if __name__ == "__main__":
    main()
