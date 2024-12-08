import inspect
import os
from pathlib import Path
from typing import Generator, Self


class Input:
    def __init__(self, path: Path | str) -> None:
        if isinstance(path, str):
            path = Path(path)
        self.path = path

    @classmethod
    def from_example(cls, year: int, day: int, part: int | None = None) -> Self:
        return cls(
            f"{os.getenv('AOC_DATA_ROOT')}/{year}/examples/{day:02d}{part if part else ''}.txt"
        )

    @classmethod
    def from_input(cls, year: int, day: int) -> Self:
        return cls(f"{os.getenv('AOC_DATA_ROOT')}/{year}/inputs/{day:02d}.txt")

    @classmethod
    def from_file_relpath(cls, path: str | Path) -> Self:
        """Interpret `path` as relative to `__file__` of the caller."""
        if isinstance(path, str):
            path = Path(path)
        path = Path(inspect.stack()[1].filename).resolve().parent / path
        return cls(path)

    def coords(self) -> Generator[tuple[tuple[int, int], str]]:
        for r, row in enumerate(self.lines()):
            for c, col in enumerate(row):
                yield (r, c), col

    def elements(self, sep: str | None = None) -> Generator[str]:
        """Yield elements stratified over lines.
        If sep is None, elements are at each index; "a  o,c" -> ['a', ' ', ' ', 'o', ',', 'c']
        If sep is '' (empty str), all whitespace is stripped; "a  o,c" -> ['a', 'o,c']
        If sep is str, that sep is used to split; (e.g. sep="o,") -> ['a  ', 'c']
        """
        with self.path.open() as f:
            while line := f.readline().rstrip("\n"):
                chars = list(line) if sep == "" else line.split(sep)
                for char in chars:
                    yield char

    def lines(self) -> Generator[str]:
        with self.path.open() as f:
            while line := f.readline().rstrip("\n"):
                yield line

    def sections(self, gap: int = 1) -> Generator[list[str]]:
        empty_lines = 0
        with self.path.open() as f:
            section = []
            while line := f.readline().rstrip("\n"):
                if line == "":
                    empty_lines += 1
                    if empty_lines == gap:
                        yield section
                        empty_lines = 0
                        section = []
                section.append(line)
