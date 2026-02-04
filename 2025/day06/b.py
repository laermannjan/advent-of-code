from math import prod
import sys

from utils import input


def main():
    cols = list(zip(*input.lines()))

    total = 0
    section = []
    for *num, op_col in cols:
        if op_col != " ":
            op = sum if op_col == "+" else prod
        if set(num) == {" "}:
            sub = op(section)
            print(section, sub)
            total += sub
            section = []
        else:
            section.append(int("".join(num)))

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
