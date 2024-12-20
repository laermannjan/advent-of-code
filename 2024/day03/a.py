import re
import sys

from utils import input

matches = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", "".join(input.lines()))
res = [int(x) * int(y) for x, y in matches]
print(sum(res), file=sys.stderr)
