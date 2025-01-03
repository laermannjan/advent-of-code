import sys
from collections import deque

from utils import input

DIM = 71
BYTES = 1024

# DIM = 7
# BYTES = 12


def main():
    grid = [["." for _ in range(DIM)] for _ in range(DIM)]
    for i, line in enumerate(input.lines()):
        if i == BYTES:
            break
        x, y = map(int, line.split(","))
        grid[y][x] = "#"

    for row in grid:
        print("".join(row))

    q = deque([(0, 0, 0)])
    seen = set()

    while q:
        dist, r, c = q.popleft()
        if (r, c) in seen:
            continue
        if (r, c) == (DIM - 1, DIM - 1):
            print(dist, file=sys.stderr)
            break
        seen.add((r, c))
        for nr, nc in [(r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)]:
            if not 0 <= nr < DIM or not 0 <= nc < DIM:
                continue
            if grid[nr][nc] == "#":
                continue
            q.append((dist + 1, nr, nc))


if __name__ == "__main__":
    main()
