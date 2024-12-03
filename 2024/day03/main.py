import re

from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
    matches = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", "".join(input.lines()))
    res = [int(x) * int(y) for x, y in matches]
    return sum(res)


def part_two(input: Input) -> int | None:
    matches = re.findall(
        r"(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))", "".join(input.lines())
    )

    include = True
    result = 0
    print(matches)
    for inst, x, y in matches:
        if inst == "do()":
            include = True
        elif inst == "don't()":
            include = False
        elif include:
            result += int(x) * int(y)

    return result


if __name__ == "__main__":
    Day(part_one, part_two).run()
