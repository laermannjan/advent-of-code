from math import prod
import sys

from utils import input


def main():
    cols = list(zip(*[line.split() for line in input.lines()]))
    print(cols)

    total = 0
    for *nums, op_col in cols:
        if op_col == "+":
            op = sum
        else:
            op = prod

        total += op(map(int, nums))

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
