from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
    pass


def part_two(input: Input) -> int | None:
    pass


if __name__ == "__main__":
    assert (
        p1 := part_one(Input.from_file_relpath("example.txt"))
    ) == None, f"Part one failed, {p1=}"
    assert (
        p2 := part_two(Input.from_file_relpath("example.txt"))
    ) == None, f"Part two failed, {p2=}"
    Day(part_one, part_two).run()
