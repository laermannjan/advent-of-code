import sys
from itertools import combinations
from math import prod

from utils import input


def main():
    points = [tuple(map(int, point.split(","))) for point in input.lines()]

    def dist(a, b):
        (x1, y1, z1), (x2, y2, z2) = a, b
        return (x1 - x2) ** 2 + (y1 - y2) ** 2 + (z1 - z2) ** 2

    pairs = sorted([(a, b) for a, b in combinations(points, 2)], key=lambda p: dist(*p))
    circuits = []

    for a, b in pairs[:1000]:  # NOTE: change to 10 for example
        print(f"{a=}, {b=}")

        a_in, b_in = None, None
        for circuit in circuits:
            if a in circuit:
                a_in = circuit
            if b in circuit:
                b_in = circuit

        if a_in is None and b_in is None:
            circuits.append({a, b})
        elif a_in is not None and b_in is None:
            a_in.add(b)
        elif b_in is not None and a_in is None:
            b_in.add(a)
        elif a_in != b_in:
            circuits.remove(a_in)
            circuits.remove(b_in)
            circuits.append(a_in | b_in)

        print(circuits)

    total = prod(sorted(map(len, circuits))[-3:])

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
