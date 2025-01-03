import sys

from utils import input


def main():
    grid = dict(input.coords())
    for (r, c), tile in grid.items():
        if tile == "S":
            break

    dists = {(r, c): 0}

    while grid[(r, c)] != "E":
        for nr, nc in [(r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)]:
            if grid[(nr, nc)] == "#":
                continue
            if (nr, nc) in dists:
                continue
            dists[(nr, nc)] = dists[(r, c)] + 1
            r, c = nr, nc

    count = 0
    for r, c in dists:
        for nr, nc in [
            (r - 2, c),
            (r - 1, c - 1),
            (r, c - 2),
            (r + 1, c - 1),
            (r + 2, c),
            (r + 1, c + 1),
            (r, c + 2),
            (r - 1, c + 1),
        ]:
            # NOTE: lower threshold to 3 for example
            if dists.get((nr, nc), 0) - dists[(r, c)] >= 102:
                count += 1

    print(dists)
    print(count, file=sys.stderr)


if __name__ == "__main__":
    main()
