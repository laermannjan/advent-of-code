import sys

from utils import input


def main():
    fresh, available = input.sections()
    total = 0
    for av in available:
        for start, end in [list(map(int, f.split("-"))) for f in fresh]:
            if start <= int(av) <= end:
                total += 1
                print(f"{av} fresh")
                break
    print()
    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
