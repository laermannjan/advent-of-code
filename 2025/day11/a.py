import sys
from collections import deque

from utils import input


def main():
    conns = {
        fro.strip(): tuple(to.strip() for to in tos.split())
        for line in input.lines()
        for fro, tos in [line.split(":")]
    }

    print(conns)

    valid = set()
    q = deque([["you"]])

    while q:
        path = q.pop()
        print(f"{path=}")
        for next in conns[path[-1]]:
            if next == "out":
                valid.add(tuple(path + [next]))
            elif next not in path:
                # print("appending ", tuple(path + [next]))
                q.append(path + [next])

    print(len(valid))

    print("not implemented", file=sys.stderr)


if __name__ == "__main__":
    main()
