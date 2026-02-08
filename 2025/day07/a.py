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

    beams = deque([S])
    print(f"{beams=}")
    seen = set()
    count = 0

    while len(beams) > 0:
        (r, c) = beams.popleft()
        print(f"{(r, c)=}")
        if (r, c) in seen:
            continue

        if grid[r][c] in ".S":
            if r == len(grid) - 1:
                continue
            beams.append((r + 1, c))
            print(f"{r=}, {count=}")
        elif grid[r][c] == "^":
            count += 1
            beams.append((r, c - 1))
            beams.append((r, c + 1))

        seen.add((r, c))
    print(count, file=sys.stderr)


if __name__ == "__main__":
    main()
