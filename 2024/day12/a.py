import sys

from utils import input


def region(
    grid: dict[tuple[int, int], str],
    start: tuple[int, int],
):
    area = set()
    stack = [start]
    while stack:
        row, col = stack.pop()
        plant = grid[(row, col)]
        area.add((row, col))

        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = row + dr, col + dc
            if (nr, nc) not in grid:
                continue
            if (nr, nc) in area:
                continue
            if grid[(nr, nc)] == plant:
                stack.append((nr, nc))
    return area


def main():
    grid = dict(input.coords())
    covered = set()

    price = 0

    for (row, col), plant in grid.items():
        if (row, col) in covered:
            continue
        area = region(grid, (row, col))
        covered |= area

        perimiter = 0

        for arow, acol in area:
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                nr, nc = arow + dr, acol + dc
                if (nr, nc) not in area:
                    perimiter += 1
        print(f"found region {plant=}, {area=}, {perimiter=}")
        price += len(area) * perimiter

    print(price, file=sys.stderr)


if __name__ == "__main__":
    main()
