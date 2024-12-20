import re
import sys
from functools import total_ordering
from math import ceil, floor

from utils import input


class NoExplosion(Exception):
    pass


class NoSplitting(Exception):
    pass


@total_ordering
class SnailfishNumber:
    def __init__(self, number: str):
        self._number = number
        self._reduce()

    @property
    def magnitude(self):
        s = self._number
        while s.count("["):
            pairs = re.findall(r"\[(\d+),(\d+)\]", s)
            for a, b in pairs:
                s = s.replace(f"[{a},{b}]", str(int(a) * 3 + int(b) * 2))
        return int(s)

    def _reduce(self):
        reduced = False
        while not reduced:
            try:
                self._number = self._explode()
            except NoExplosion:
                try:
                    self._number = self._split()
                except NoSplitting:
                    reduced = True

    def _explode(self) -> str:
        depth = 0
        for i, v in enumerate(self._number):
            if v.isnumeric() and depth > 4:
                closing_bracket_distance = self._number[i:].index("]")
                slice_before = self._number[: i - 1]
                slice_after = self._number[i + closing_bracket_distance + 1 :]
                pair = [
                    *map(int, self._number[i : i + closing_bracket_distance].split(","))
                ]

                if len(regulars_before := re.findall(r"\d+", slice_before)):
                    last_regular = regulars_before[-1]
                    last_index = slice_before.rindex(last_regular)
                    slice_before = (
                        slice_before[:last_index]
                        + str(int(last_regular) + pair[0])
                        + slice_before[last_index + len(last_regular) :]
                    )

                if len(regulars_after := re.findall(r"\d+", slice_after)):
                    next_regular = regulars_after[0]
                    next_index = slice_after.index(next_regular)
                    slice_after = (
                        slice_after[:next_index]
                        + str(int(next_regular) + pair[1])
                        + slice_after[next_index + len(next_regular) :]
                    )
                return slice_before + "0" + slice_after
            else:
                if v == "[":
                    depth += 1
                elif v == "]":
                    depth -= 1
        raise NoExplosion

    def _split(self) -> str:
        s = self._number
        large_regulars = [i for i in re.findall(r"\d+", self._number) if int(i) > 9]
        if len(large_regulars):
            regular = large_regulars[0]
            regular_index = self._number.index(regular)

            slice_before = self._number[:regular_index]
            slice_after = self._number[regular_index + len(regular) :]
            return (
                slice_before
                + f"[{floor(int(regular) / 2)},{ceil(int(regular) / 2)}]"
                + slice_after
            )

        else:
            raise NoSplitting

    def __lt__(self, other: "SnailfishNumber"):
        if not isinstance(other, SnailfishNumber):
            return NotImplemented
        return self.magnitude < other.magnitude

    def __eq__(self, other):
        if not isinstance(other, SnailfishNumber):
            return NotImplemented
        return self.magnitude == other.magnitude

    def __add__(self, other: "SnailfishNumber") -> "SnailfishNumber":
        if isinstance(other, str):
            other = self.__class__(other)
        return self.__class__(f"[{self._number},{other._number}]")

    def __radd__(self, other) -> "SnailfishNumber":
        return self if other == 0 else self.__add__(other)


total = sum(SnailfishNumber(num) for num in input.lines())
print(total.magnitude, file=sys.stderr)
