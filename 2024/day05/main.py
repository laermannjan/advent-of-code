from functools import cmp_to_key
from itertools import combinations

from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
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

    return result


def part_two(input: Input) -> int | None:
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
    return result


if __name__ == "__main__":
    assert (
        p1 := part_one(Input.from_file_relpath("example.txt"))
    ) == 143, f"Part one failed, {p1=}"
    assert (
        p2 := part_two(Input.from_file_relpath("example.txt"))
    ) == 123, f"Part two failed, {p2=}"
    Day(part_one, part_two).run()
