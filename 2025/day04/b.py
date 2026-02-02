import sys

from utils import input


def main():
    grid = {k: v for k, v in input.coords()}
    accessible = set()
    removed = True

    while removed:
        removed = False
        for (x, y), cell in grid.items():
            if cell != "@":
                continue
            rolls = 0
            for dx in [-1, 0, 1]:
                for dy in [-1, 0, 1]:
                    if dx == dy == 0:
                        continue
                    if (x + dx, y + dy) in grid and grid[x + dx, y + dy] == "@":
                        rolls += 1
            if rolls < 4:
                accessible.add((x, y))
                grid[(x, y)] = "."
                removed = True

    print(len(accessible), file=sys.stderr)


if __name__ == "__main__":
    main()
