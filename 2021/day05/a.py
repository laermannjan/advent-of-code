import sys
from collections import Counter

from utils import input


def parse_points(line: str) -> tuple[tuple[int, int], tuple[int, int]]:
    points = line.split(" -> ")
    p1 = points[0]
    p2 = points[1]

    p1_coords = [int(x) for x in p1.split(",")]
    p2_coords = [int(x) for x in p2.split(",")]

    x1 = p1_coords[0]
    y1 = p1_coords[1]
    x2 = p2_coords[0]
    y2 = p2_coords[1]

    return (x1, y1), (x2, y2)


def point2line(
    x1: int, y1: int, x2: int, y2: int, with_diag: bool
) -> list[tuple[int, int]]:
    if x1 == x2:
        return [(x1, y) for y in range(min(y1, y2), max(y1, y2) + 1)]
    elif y1 == y2:
        return [(x, y1) for x in range(min(x1, x2), max(x1, x2) + 1)]
    elif with_diag:
        if x1 - x2 == y1 - y2:
            return list(
                zip(
                    range(min(x1, x2), max(x1, x2) + 1),
                    range(min(y1, y2), max(y1, y2) + 1),
                )
            )
        elif x1 - x2 == y2 - y1:
            return list(
                zip(
                    range(min(x1, x2), max(x1, x2) + 1),
                    reversed(range(min(y1, y2), max(y1, y2) + 1)),
                )
            )
        else:
            return []
    else:
        return []


def count_crossings(input: list[str], with_diag: bool) -> int:
    cntr = Counter()
    for line in input:
        p1, p2 = parse_points(line)
        points_on_line = point2line(*p1, *p2, with_diag=with_diag)
        cntr.update(points_on_line)

    return len([freq for _, freq in cntr.items() if freq >= 2])


print(count_crossings(input.lines(), False), file=sys.stderr)
