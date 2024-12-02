import time
from argparse import ArgumentParser
from collections.abc import Callable

from utils.input import Input


class Day:
    def __init__(
        self,
        part_one: Callable[[Input], int | None],
        part_two: Callable[[Input], int | None],
    ) -> None:
        self.part_one = part_one
        self.part_two = part_two

    def run(self) -> None:
        parser = ArgumentParser()
        parser.add_argument("-i", "--input", type=str)
        parser.add_argument("-p", "--part", type=str)
        parser.add_argument("-v", "--verbose", action="store_true")
        args = parser.parse_args()

        if args.part not in ("one", "two"):
            ValueError(f"invalid part {args.part}")

        fn = self.part_one if args.part == "one" else self.part_two

        start = time.perf_counter()
        result = fn(Input(args.input))
        elapsed = time.perf_counter() - start
        print(f"Part {args.part}: {result} (took: {elapsed:.2f}s)")
