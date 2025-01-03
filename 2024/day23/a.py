import sys
from itertools import combinations

from utils import input


def main():
    conns = {}
    for line in input.lines():
        a, b = line.split("-")
        conns.setdefault(a, set()).add(b)
        conns.setdefault(b, set()).add(a)

    networks = set()
    for this in conns:
        if len(conns[this]) < 2:
            continue
        for a, b in combinations(conns[this], r=2):
            if b in conns[a] and any(c.startswith("t") for c in [this, a, b]):
                networks.add(frozenset({this, a, b}))

    print(len(networks), file=sys.stderr)


if __name__ == "__main__":
    main()
