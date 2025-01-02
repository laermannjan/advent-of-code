import sys

from utils import input


def main():
    house, instructions = list(input.sections())

    house = list(map(list, house))
    instructions = "".join(instructions)

    for r, row in enumerate(house):
        for c, ele in enumerate(row):
            if ele == "@":
                rr, rc = r, c
                break
    print("initial state")
    for _ in house:
        print("".join(_))

    for inst in instructions:
        if inst == "<":
            dr, dc = 0, -1
        elif inst == ">":
            dr, dc = 0, 1
        elif inst == "^":
            dr, dc = -1, 0
        elif inst == "v":
            dr, dc = 1, 0
        else:
            raise ValueError("invalid instruction", inst)

        boxes = []
        nr, nc = rr + dr, rc + dc
        while house[nr][nc] == "O":
            boxes.append((nr, nc))
            nr += dr
            nc += dc

        if house[nr][nc] != "#":
            if boxes:
                house[boxes[-1][0] + dr][boxes[-1][1] + dc] = "O"
            house[rr + dr][rc + dc] = "@"
            house[rr][rc] = "."
            rr, rc = rr + dr, rc + dc

        print("-" * 20)
        print("move", inst)
        for _ in house:
            print("".join(_))

    s = 0
    for r, row in enumerate(house):
        for c, ele in enumerate(row):
            if ele == "O":
                s += 100 * r + c

    print(s, file=sys.stderr)


if __name__ == "__main__":
    main()
