import sys
from itertools import product

from utils import input


def concat(a: int, b: int) -> int:
    return int(str(a) + str(b))


total = 0

for l, line in enumerate(inp := input.lines()):
    print(f"{l + 1} / {len(inp)}")
    res = int(line.split(": ")[0])
    nums = [int(d) for d in line.split(": ")[1].split(" ")]

    combos = list(product([int.__add__, int.__mul__, concat], repeat=len(nums) - 1))

    print(f"{res=}, {nums=}, {len(nums)=}, {len(combos)=}")
    for operators in combos:
        result = nums[0]
        for i, op in enumerate(operators):
            result = op(result, nums[i + 1])
        print(f"  {result=}")
        if result == res:
            total += res
            print("    matched")
            break
print(total, file=sys.stderr)
