import sys
from collections import deque

from utils import input

DIM = 71
BYTES = 991024

# DIM = 7
# BYTES = 12


def reachable(grid):
    q = deque([(0, 0, 0)])
    seen = set()

    while q:
        dist, r, c = q.popleft()
        if (r, c) in seen:
            continue
        if (r, c) == (DIM - 1, DIM - 1):
            return True
        seen.add((r, c))
        for nr, nc in [(r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)]:
            if not 0 <= nr < DIM or not 0 <= nc < DIM:
                continue
            if grid[nr][nc] == "#":
                continue
            q.append((dist + 1, nr, nc))
    return False


def main():
    G = [["." for _ in range(DIM)] for _ in range(DIM)]
    BLOCKS = [tuple(map(int, line.split(","))) for line in input.lines()]

    # binary search to find at which byte we cannot reach the goal
    fail_at = len(BLOCKS)
    success_at = 0

    def apply(bytes):
        grid = [
            ["#" if (r, c) in BLOCKS[:bytes] else x for c, x in enumerate(row)]
            for r, row in enumerate(G)
        ]
        return grid

    while success_at + 1 != fail_at:
        bytes = (fail_at + success_at) // 2
        grid = apply(bytes)
        if reachable(grid):
            success_at = bytes
        else:
            fail_at = bytes

    print(*BLOCKS[success_at], sep=",", file=sys.stderr)


if __name__ == "__main__":
    main()
