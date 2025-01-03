import sys
from itertools import product

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

    cheats = {}
    for r, c in dists:
        for dr, dc in product(range(-20, 21), repeat=2):
            dist = abs(dr) + abs(dc)
            if not 2 <= dist <= 20:
                continue
            nr, nc = r + dr, c + dc
            # NOTE: set thresh to 50 for example
            if (diff := dists.get((nr, nc), 0) - dists[(r, c)] - dist) >= 100:
                cheats[(r, c, nr, nc)] = diff

    diffs = {}
    for cheat, diff in cheats.items():
        if diff not in diffs:
            diffs[diff] = set()
        diffs[diff].add(cheat)

    for diff in sorted(diffs):
        print(f"{diff=} {len(diffs[diff])=}")

    print(len(cheats), file=sys.stderr)


if __name__ == "__main__":
    main()
