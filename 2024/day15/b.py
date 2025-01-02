import sys

from utils import input


def main():
    house, instructions = list(input.sections())

    scaled = []
    for row in house:
        scaled.append([])
        for ele in row:
            if ele == "#":
                scaled[-1].extend(list("##"))
            elif ele == "O":
                scaled[-1].extend(list("[]"))
            elif ele == ".":
                scaled[-1].extend(list(".."))
            elif ele == "@":
                scaled[-1].extend(list("@."))
    house = scaled
    instructions = "".join(instructions)

    for r, row in enumerate(house):
        for c, ele in enumerate(row):
            print(ele)
            if ele == "@":
                rr, rc = r, c
                print("found")
                break
    print("initial state")
    for _ in house:
        print("".join(_))

    for inst in instructions:
        print("-" * 20)
        if inst in "<>":
            # left/right still can only happen in a single row
            dc = 1 if inst == ">" else -1

            boxes = []
            nr, nc = rr, rc + dc
            while house[nr][nc] in "[]":
                boxes.append((nr, nc))
                nc += dc + dc

            if house[nr][nc] != "#":
                house[nr].pop(nc)
                house[rr].insert(rc, ".")
                rr, rc = rr, rc + dc

        elif inst in ["^", "v"]:
            dr = -1 if inst == "^" else 1

            # look up/down, if there is a box half, add it and the other part to the stack
            # for each box-half on the stack see if it's a left or a right half of a box
            # if it is a box-half, add it and its counterpart to the box set
            # also, add the two coordinates above/below it to the stack
            # i.e. a "detected" box will always add the two coords directly above/below it
            # and because a box half automatically adds the counterpart half, we can expand
            if house[rr + dr][rc] not in ".#":
                boxes = set()
                stack = {(rr + dr, rc)}
                if house[rr + dr][rc] == "[":
                    stack.add((rr + dr, rc + 1))
                elif house[rr + dr][rc] == "]":
                    stack.add((rr + dr, rc - 1))

                while stack:
                    r, c = stack.pop()
                    if house[r][c] == "[":
                        boxes.add((r, c))
                        boxes.add((r, c + 1))
                        stack.add((r + dr, c))
                        stack.add((r + dr, c + 1))
                    elif house[r][c] == "]":
                        boxes.add((r, c))
                        boxes.add((r, c - 1))
                        stack.add((r + dr, c - 1))
                        stack.add((r + dr, c))
                print(boxes)

                for box in boxes:
                    # we need to check for every box if there would be a collision
                    # because we could have "gaps", e.g. a U-shaped pile of boxes infront of us
                    r, c = box
                    if house[r + dr][c] == "#":
                        break
                else:
                    print("no obstacle, moving boxes")
                    new_house = [[e for e in row] for row in house]
                    for box in sorted(boxes, key=lambda x: x[0] * dr * -1):
                        r, c = box
                        print(f"moving {box=} from ({r},{c}) to ({r+dr},{c})")
                        new_house[r + dr][c] = house[r][c]
                        new_house[r][c] = "."
                    house = new_house
            if house[rr + dr][rc] == ".":
                house[rr + dr][rc] = "@"
                house[rr][rc] = "."
                rr, rc = rr + dr, rc

        else:
            raise ValueError("invalid instruction", inst)

        print("move", inst)
        for _ in house:
            print("".join(_))
        print("-" * 20)
    s = 0
    for r, row in enumerate(house):
        for c, ele in enumerate(row):
            if ele == "[":
                s += 100 * r + c

    print(s, file=sys.stderr)


if __name__ == "__main__":
    main()
