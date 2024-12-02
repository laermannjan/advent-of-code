from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
    lines = [list(map(int, line.split())) for line in input.lines()]
    lists = list(map(list, zip(*lines)))

    for lst in lists:
        lst.sort()

    return sum([abs(x - y) for x, y in zip(*lists)])


def part_two(input: Input) -> int | None:
    pass


if __name__ == "__main__":
    Day(part_one, part_two).run()