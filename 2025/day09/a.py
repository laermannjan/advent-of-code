import sys
from itertools import combinations

from utils import input


def main():
    points = [tuple(map(int, point.split(","))) for point in input.lines()]
    pairs = list(combinations(points, 2))

    print(points)

    def area(a, b):
        (x1, y1), (x2, y2) = a, b
        return (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)

    pairs.sort(key=lambda p: area(*p))

    print(pairs[-1])

    print(area(*pairs[-1]), file=sys.stderr)


if __name__ == "__main__":
    main()
