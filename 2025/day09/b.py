import sys
from itertools import combinations

from utils import input


def main():
    points = [tuple(map(int, point.split(","))) for point in input.lines()]
    pairs = list(combinations(points, 2))
    edges = list(zip(points, points[1:])) + [(points[-1], points[0])]

    def area(a, b):
        return (abs(a[0] - b[0]) + 1) * (abs(a[1] - b[1]) + 1)

    def inside(a, b):
        # AABB test (axis-aligned bounding box)
        (rect_x_min, rect_x_max), (rect_y_min, rect_y_max) = map(sorted, zip(a, b))

        for edge_start, edge_end in edges:
            edge_x_min, edge_x_max = sorted([edge_start[0], edge_end[0]])
            edge_y_min, edge_y_max = sorted([edge_start[1], edge_end[1]])

            if (
                rect_x_min < edge_x_max
                and rect_x_max > edge_x_min
                and rect_y_min < edge_y_max
                and rect_y_max > edge_y_min
            ):
                return True
        return False

    max_area = 0
    max_pair = None
    for a, b in pairs:
        this_area = area(a, b)
        if this_area > max_area:
            if not inside(a, b):
                max_area = this_area
                max_pair = (a, b)

    print(max_pair)
    print(max_area, file=sys.stderr)


if __name__ == "__main__":
    main()
