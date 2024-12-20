import sys

from utils import input

lines = [list(map(int, line.split())) for line in input.lines()]
lists = list(map(list, zip(*lines)))

for lst in lists:
    lst.sort()

print(sum([abs(x - y) for x, y in zip(*lists)]), file=sys.stderr)
