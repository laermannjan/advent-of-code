import sys

from utils import input


def rating(coords: dict[tuple[int, int], str], start: tuple[int, int]) -> int:
    seen = set()
    stack = [start]

    n_trails = 0

    while stack:
        row, col = stack.pop()
        if coords[(row, col)] == "9":
            print(f"trail end at ({row},{col})")
            n_trails += 1
            continue
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            pos = row + dr, col + dc
            if (
                pos not in seen
                and pos in coords
                and coords[pos] != "."
                and int(coords[pos]) - int(coords[(row, col)]) == 1
            ):
                stack.insert(0, pos)

    return n_trails


def main():
    coords = dict(input.coords())

    total = 0
    for (row, col), height in coords.items():
        if height != "0":
            continue
        s = rating(coords, (row, col))
        print(f"tailhead at ({row},{col}) - rating {s}")
        total += s

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
