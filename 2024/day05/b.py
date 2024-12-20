import sys
from functools import cmp_to_key
from itertools import combinations

from utils import input

rules_, updates = (s for s in input.sections())

rules = {}
for rule in rules_:
    x, y = map(int, rule.split("|"))
    rules[(x, y)] = -1
    rules[(y, x)] = 1


def is_ordered(update: list[int]) -> bool:
    for i, j in combinations(update, 2):
        if rules.get((i, j)) == 1:
            return False
    return True


result = 0


def cmp(x, y):
    return rules.get((x, y), 0)


for update in updates:
    update = list(map(int, update.split(",")))
    if not is_ordered(update):
        update.sort(key=cmp_to_key(cmp))
        result += update[len(update) // 2]
print(result, file=sys.stderr)
