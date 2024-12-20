import re
import sys

from utils import input

matches = re.findall(
    r"(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))", "".join(input.lines())
)
include = True
result = 0
for inst, x, y in matches:
    if inst == "do()":
        include = True
    elif inst == "don't()":
        include = False
    elif include:
        result += int(x) * int(y)

print(result, file=sys.stderr)
