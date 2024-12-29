import sys

from utils import input


def score(coords: dict[tuple[int, int], str], start: tuple[int, int]) -> int:
    seen = set()
    stack = [start]

    endings = set()

    while stack:
        row, col = stack.pop()
        if coords[(row, col)] == "9":
            endings.add((row, col))
            continue
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            pos = row + dr, col + dc
            if (
                pos not in seen
                and pos in coords
                and int(coords[pos]) - int(coords[(row, col)]) == 1
            ):
                stack.append(pos)

    return len(endings)


coords = dict(input.coords())


total = 0
for (row, col), height in coords.items():
    if height != "0":
        continue
    s = score(coords, (row, col))
    print(f"tailhead at ({row},{col}) - score {s}")
    total += s

print(total, file=sys.stderr)
