import sys

from utils import input

horizontal = 0
depth = 0

for line in input.lines():
    inst, val_str = line.split(" ")
    val = int(val_str)

    if inst == "forward":
        horizontal += val
    else:
        depth += val if inst == "down" else -val

print(horizontal * depth, file=sys.stderr)
