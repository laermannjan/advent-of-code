import sys
from itertools import combinations

from utils import input


def main():
    conns = {}
    for line in input.lines():
        a, b = line.split("-")
        conns.setdefault(a, set()).add(b)
        conns.setdefault(b, set()).add(a)

    largest_network = set()
    for comp in conns:
        # check all connection groups which are greater than our current largest network
        for groupsize in range(max(len(largest_network) + 1, 2), len(conns[comp]) + 1):
            for others in combinations(conns[comp], r=groupsize):
                # if all of them are connected amongst each other, its a larger network
                if all(b in conns[a] for a, b in combinations(others, r=2)):
                    largest_network = set([comp, *others])

    print(",".join(sorted(largest_network)), file=sys.stderr)


if __name__ == "__main__":
    main()
