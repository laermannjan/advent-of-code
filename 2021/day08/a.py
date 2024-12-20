import sys

from utils import input

counter = 0
for line in input.lines():
    outputs = line.split(" | ")[1].split(" ")
    for digit in outputs:
        if len(digit) in (2, 3, 4, 7):
            counter += 1
print(counter, file=sys.stderr)
