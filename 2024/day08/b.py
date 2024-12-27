import sys
from itertools import combinations

from utils import input

antenna = {}

antinodes = set()

for loc, val in input.coords():
    if val == ".":
        continue
    if val not in antenna:
        antenna[val] = []
    antenna[val].append(loc)

max_r, max_c = loc

for freq, locs in antenna.items():
    print(freq)
    for (rx, cx), (ry, cy) in combinations(locs, 2):
        print(f"({rx},{cx}), ({ry},{cy})")
        rd = ry - rx
        cd = cy - cx

        r1 = rx
        c1 = cx

        while 0 <= r1 <= max_r and 0 <= c1 <= max_c:
            print(f"antinode at ({r1},{c1})")
            antinodes.add((r1, c1))
            r1 -= rd
            c1 -= cd

        r2 = ry
        c2 = cy

        while 0 <= r2 <= max_r and 0 <= c2 <= max_c:
            print(f"antinode at ({r2},{c2})")
            antinodes.add((r2, c2))
            r2 += rd
            c2 += cd


print(len(antinodes), file=sys.stderr)
