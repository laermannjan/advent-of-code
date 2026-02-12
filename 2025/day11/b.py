import queue
import sys
from collections import deque
from functools import cache
from itertools import chain

from utils import input


def main():
    conns = {
        fro.strip(): tuple(to.strip() for to in tos.split())
        for line in input.lines()
        for fro, tos in [line.split(":")]
    }

    print(conns)

    @cache
    def paths(start, end):
        if start == end:
            return 1
        return sum(paths(neighbor, end) for neighbor in conns.get(start, []))

    # Inspired by the code above by HyperNeutrino; doesn't work, however, since each cache entry
    # would have to store all paths, not just the count. And the number of paths can be VERY large.
    # @cache
    # def paths(start, end):
    #     if start == end:
    #         return {(end,)}
    #     return {
    #         (start, *path)
    #         for neighbor in conns.get(start, [])
    #         for path in paths(neighbor, end)
    #     }

    print(
        paths("svr", "dac") * paths("dac", "fft") * paths("fft", "out")
        + paths("svr", "fft") * paths("fft", "dac") * paths("dac", "out"),
        file=sys.stderr,
    )


if __name__ == "__main__":
    main()
