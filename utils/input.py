import sys
from typing import Generator, TextIO


def stdin() -> TextIO:
    if sys.stdin.isatty():
        raise RuntimeError("no stdin")
    return sys.stdin


def lines() -> list[str]:
    return stdin().read().splitlines()


def coords() -> Generator[tuple[tuple[int, int], str], None, None]:
    for r, row in enumerate(lines()):
        for c, ele in enumerate(row):
            yield (r, c), ele


def sections(gap: int = 1) -> Generator[list[str], None, None]:
    empty_lines = 0
    section = []
    for line in lines():
        if line == "":
            empty_lines += 1
            if empty_lines == gap:
                yield section
                empty_lines = 0
                section = []
        else:
            section.append(line)
    yield section
