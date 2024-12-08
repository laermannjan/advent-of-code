from itertools import combinations

from utils.execute import Day
from utils.input import Input


def is_ordered(update: list[int], rules: dict[tuple[int, int], bool]) -> bool:
    for i, j in combinations(update, 2):
        if not rules.get((i, j)):
            return False
    return True


def part_one(input: Input) -> int | None:
    rules, updates = (s for s in input.sections())

    rule_cache = {}
    for rule in rules:
        x, y = map(int, rule.split("|"))
        rule_cache[(x, y)] = True
        rule_cache[(y, x)] = False

    result = 0
    for update in updates:
        update = list(map(int, update.split(",")))
        if is_ordered(update, rule_cache):
            result += update[len(update) // 2]

    return result


def part_two(input: Input) -> int | None:
    pass


if __name__ == "__main__":
    assert (
        p1 := part_one(Input.from_file_relpath("example.txt"))
    ) == 143, f"Part one failed, {p1=}"
    assert (
        p2 := part_two(Input.from_file_relpath("example.txt"))
    ) == None, f"Part two failed, {p2=}"
    Day(part_one, part_two).run()
