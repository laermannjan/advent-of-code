import sys
from itertools import combinations

from utils import input

rules_, updates = (s for s in input.sections())

rules = {}
for rule in rules_:
    x, y = map(int, rule.split("|"))
    rules[(x, y)] = True
    rules[(y, x)] = False


def is_ordered(update: list[int]) -> bool:
    for i, j in combinations(update, 2):
        if not rules.get((i, j)):
            return False
    return True


result = 0
for update in updates:
    update = list(map(int, update.split(",")))
    if is_ordered(update):
        result += update[len(update) // 2]

print(result, file=sys.stderr)
