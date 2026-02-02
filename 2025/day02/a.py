import sys

from utils import input


def main():
    ranges = [list(map(int, r.split("-"))) for r in input.stdin().read().split(",")]
    numbers = sum([list(range(start, end + 1)) for start, end in ranges], [])

    invalids = 0

    for num in numbers:
        s = str(num)
        if len(s) % 2 == 0 and s[: len(s) // 2] * 2 == s:
            print(num)
            invalids += num
    print(invalids, file=sys.stderr)


if __name__ == "__main__":
    main()
