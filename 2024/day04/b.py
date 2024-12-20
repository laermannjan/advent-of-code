import sys

from utils import input

grid = list(input.lines())
cnt = 0
for (row, col), letter in input.coords():
    if row == 0 or row == len(grid) - 1 or col == 0 or col == len(grid[0]) - 1:
        continue

    if letter != "A":
        continue

    surround = "".join(
        grid[row + dr][col + dc] for (dr, dc) in [(-1, -1), (-1, 1), (1, 1), (1, -1)]
    )
    if surround in ["MMSS", "SMMS", "SSMM", "MSSM"]:
        cnt += 1
print(cnt, file=sys.stderr)
