import sys

from utils import input


def main():
    fresh = [list(map(int, f.split("-"))) for f in next(input.sections())]
    total = 0

    last = None
    for start, end in sorted(fresh):
        if last is None:
            last = end
            total += end - start + 1
            print(f"+ {start}-{end}")
        elif end > last:
            total += end - max(start, last + 1) + 1
            print(f"+ ({last+1=}|{start=})-{end}")
            last = end

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
