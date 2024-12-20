import sys
from itertools import product

from utils import input

grid = list(input.lines())
cnt = 0
for (row, col), letter in input.coords():
    if letter != "X":
        continue
    for dr, dc in product([-1, 0, 1], repeat=2):
        if dc == dr == 0:
            continue
        if not (0 <= row + 3 * dr < len(grid[row]) and 0 <= col + 3 * dc < len(grid)):
            continue
        for i, l in enumerate("MAS", start=1):
            if grid[row + i * dr][col + i * dc] != l:
                break
        else:
            cnt += 1
print(cnt, file=sys.stderr)
