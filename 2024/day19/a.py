import sys
from functools import cache

from utils import input


def main():
    towels, designs = input.sections()

    towels = set(towels[0].split(", "))

    @cache
    def possible(design: str):
        if design == "":
            return True

        for t in towels:
            if design.startswith(t):
                if possible(design[len(t) :]):
                    return True
        return False

    # for d, design in enumerate(designs):
    #     print(f"{design}", d, "/", len(designs))
    #     print(f"{design} possible? {possible(design)}", file=sys.stderr)
    #     # break
    print(sum(possible(d) for d in designs), file=sys.stderr)


if __name__ == "__main__":
    main()
