import sys

import numpy as np
from utils import input


def step(grid):
    grid += 1  # regular step increase
    energized = True
    while energized:
        energized = False
        energized_grid = np.copy(
            grid
        )  # keeps track of energy transferal in this evaluation round
        for row in range(grid.shape[0]):
            for col in range(grid.shape[1]):
                flashing_neighbors = (
                    grid[max(0, row - 1) : row + 2, max(0, col - 1) : col + 2] > 9
                )
                adj_energy = np.sum(flashing_neighbors)
                if adj_energy > 0:
                    energized_grid[row, col] += adj_energy
                    energized = True

        already_flashed = grid > 9
        grid = energized_grid
        grid[already_flashed] = -100

    flashes = np.sum(grid < 0)
    grid[grid < 0] = 0
    return grid, flashes


grid = np.array([[int(x) for x in row] for row in input.lines()])
steps = 100
flashes = []

for s in range(steps):
    grid, fs = step(grid)
    flashes.append(fs)
    total_flashes = sum(flashes)
print(total_flashes, file=sys.stderr)
