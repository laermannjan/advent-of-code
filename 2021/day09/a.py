import sys

import numpy as np
from utils import input


def get_neighbors(row, col, height, width):
    neighbors = [
        (nr, nc)
        for (r, c) in [(0, 1), (1, 0), (0, -1), (-1, 0)]
        if 0 <= (nr := row + r) < height and 0 <= (nc := col + c) < width
    ]
    return neighbors


def get_low_points(heightmap):
    low_points = []
    for row in range(heightmap.shape[0]):
        for col in range(heightmap.shape[1]):
            neighbors = [
                heightmap[neigh[0], neigh[1]]
                for neigh in get_neighbors(row, col, *heightmap.shape)
            ]
            if all([heightmap[row, col] < n for n in neighbors]):
                low_points.append((row, col))
    return low_points


heightmap = np.array([[int(x) for x in row] for row in input.lines()])
low_points = get_low_points(heightmap)
print(sum([heightmap[lp] + 1 for lp in low_points]), sys.stderr)
