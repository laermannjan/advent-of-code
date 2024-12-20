import sys
from collections import Counter

from utils import input

lines = [list(map(int, line.split())) for line in input.lines()]
lists = list(map(list, zip(*lines)))

counter = Counter(lists[1])
print(sum([n * counter.get(n, 0) for n in lists[0]]), file=sys.stderr)
