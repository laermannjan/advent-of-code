import sys
from functools import cache

from utils import input


def main():
    towels, designs = input.sections()

    towels = set(towels[0].split(", "))

    @cache
    def possible(design: str):
        if design == "":
            return 1

        counts = 0
        for t in towels:
            if design.startswith(t):
                counts += possible(design[len(t) :])
        return counts

    # for d, design in enumerate(designs):
    #     print(f"{design}", d, "/", len(designs))
    #     print(f"{design} possible? {possible(design)}", file=sys.stderr)
    #     # break
    print(sum(possible(d) for d in designs), file=sys.stderr)


if __name__ == "__main__":
    main()
