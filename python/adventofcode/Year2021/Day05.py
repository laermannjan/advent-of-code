"""
-- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.
They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for you to review. For example:
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends. In other words:
An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.
So, the horizontal and vertical lines from the above list would produce the following diagram:
.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.
To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.
Consider only horizontal and vertical lines. At how many points do at least two lines overlap?
--- Part Two ---
Unfortunately, considering only horizontal and vertical lines doesn't give you the full picture; you need to also consider diagonal lines.
Because of the limits of the hydrothermal vent mapping system, the lines in your list will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees. In other words:
An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
Considering all lines from the above example would now produce the following diagram:
1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
You still need to determine the number of points where at least two lines overlap. In the above example, this is still anywhere in the diagram with a 2 or larger - now a total of 12 points.
Consider all of the lines. At how many points do at least two lines overlap?
"""

from collections import Counter


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


def parse_input(filename: str) -> list[str]:
    with open(filename, 'r') as f:
        data = f.read().splitlines()
    return data


def day05a(filename: str) -> int:
    return count_crossings(parse_input(filename), False)


def day05b(filename: str) -> int:
    return count_crossings(parse_input(filename), True)
