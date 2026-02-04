from math import prod
import sys

from utils import input


def main():
    grid = {k: v for k, v in input.coords()}
    maxrow, maxcol = max(grid.keys())

    op = ""
    results = []
    nums = []
    for col in range(maxcol + 1):
        if grid[(maxrow, col)].strip():
            op = grid[(maxrow, col)]

        num = "".join([grid[row, col] for row in range(maxrow)]).strip()
        print(f"{op=}, {num=}")
        if num == "":
            if op == "+":
                sub = sum(nums)
            else:
                sub = prod(nums)
            print(f"{sub=}")
            results.append(sub)
            nums = []
        else:
            nums.append(int(num))
            print(f"append(int({num}))", nums)

    if op == "+":
        sub = sum(nums)
    else:
        sub = prod(nums)
    print(f"{sub=}")
    results.append(sub)
    nums = []

    # total = 0
    # for col in range(len(grid[0])):
    #     prob = [grid[row][col] for row in range(len(grid))]
    #
    #     print(prob)
    #     nums, op = map(int, prob[:-1]), prob[-1]
    #
    #     lens = map(len, nums)
    #
    #     result = sum(nums) if op == "+" else prod(nums)
    #     total += result
    #
    print(sum(results), file=sys.stderr)


if __name__ == "__main__":
    main()
