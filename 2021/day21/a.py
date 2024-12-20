import sys

from utils import input
from utils.misc import sum_to


def parse_input() -> list[int]:
    data = input.lines()
    players = [int(line.split(":")[1].strip()) for line in data]
    return players


positions = parse_input()
scores = [0, 0]
times_rolled = 0
while max(scores) < 1000:
    roll = 3 * times_rolled + sum_to(3)
    times_rolled += 3
    positions[0] = int((positions[0] - 1 + roll) % 10 + 1)
    scores[0] += positions[0]
    positions = list(reversed(positions))
    scores = list(reversed(scores))
print(scores[0] * times_rolled, file=sys.stderr)
