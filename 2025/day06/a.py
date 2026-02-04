from math import prod
import sys

from utils import input


def main():
    grid = [
        [ele if ele in "+*" else int(ele) for ele in line.split()]
        for line in input.lines()
    ]

    total = 0
    for col in range(len(grid[0])):
        prob = [grid[row][col] for row in range(len(grid))]
        print(prob)
        nums, op = map(int, prob[:-1]), prob[-1]
        result = sum(nums) if op == "+" else prod(nums)
        total += result

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
