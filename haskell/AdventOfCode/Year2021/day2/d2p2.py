import sys

horizontal = 0
depth = 0
aim = 0

for line in sys.stdin:
    inst, val = line.rstrip().split(' ')
    val = int(val)

    if inst == 'forward':
        horizontal += val
        depth += aim * val
    else:
        aim += val if inst == 'down' else -val

print(horizontal * depth)