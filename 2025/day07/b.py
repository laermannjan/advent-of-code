from functools import cache
from collections import deque
import sys

from utils import input


def main():
    grid: list[list[str]] = [list(line.strip()) for line in input.lines()]

    S = [
        (r, c)
        for r, row in enumerate(grid)
        for c, char in enumerate(row)
        if char == "S"
    ].pop()

    # NOTE: each particle sub-path starting from (r, c) might have already been traversed by an earlier particle
    # hence we use caching to reduce total compute
    @cache
    def solve(r, c):
        if grid[r][c] in ".S":
            if r == len(grid) - 1:
                return 1  # if we reach the bottom, plus 1
            else:
                return solve(r + 1, c)  # traverse straight down, no additional path
        elif grid[r][c] == "^":
            return solve(r, c - 1) + solve(r, c + 1)  # sum of both legs of the splitter

    print(solve(*S), file=sys.stderr)


if __name__ == "__main__":
    main()
