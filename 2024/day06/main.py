from utils.execute import Day
from utils.input import Input


def part_one(input: Input) -> int | None:
    grid = list(input.lines())

    def in_bounds(row, col):
        return 0 <= row < len(grid) and 0 <= col < len(grid[0])

    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if grid[r][c] == "^":
                break
        else:
            continue
        break
    else:
        RuntimeError("guard not found")

    positions = set()
    dr, dc = -1, 0

    while True:
        positions.add((r, c))

        nr, nc = r + dr, c + dc
        print(f"looking at ({nr}, {nc})")

        if not in_bounds(nr, nc):
            print(f"({nr}, {nc}) is out of bounds")
            break

        if grid[nr][nc] == "#":
            print(f"({nr}, {nc}) is #")
            dr, dc = dc, -dr
            continue

        r = r + dr
        c = c + dc

        # sleep(0.5)

    return len(positions)


def part_two(input: Input) -> int | None:
    grid = [list(line) for line in input.lines()]

    def in_bounds(row, col):
        return 0 <= row < len(grid) and 0 <= col < len(grid[0])

    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if grid[r][c] == "^":
                break
        else:
            continue
        break
    else:
        RuntimeError("guard not found")

    def walk(grid, r, c):
        positions = set()
        dr, dc = -1, 0

        while True:
            positions.add((r, c, dr, dc))

            nr, nc = r + dr, c + dc
            # print(f"looking at ({nr}, {nc})")

            if not in_bounds(nr, nc):
                # print(f"({nr}, {nc}) is out of bounds")
                return positions, False

            if grid[nr][nc] == "#":
                # print(f"({nr}, {nc}) is #")
                dr, dc = dc, -dr
            else:
                r = r + dr
                c = c + dc

            if (r, c, dr, dc) in positions:
                return positions, True

    # get the guards path
    positions, _ = walk(grid, r, c)

    # place obstactle at every step of the guard's original path and see if it loops
    loops = set()
    for rr, cc, _, _ in positions:
        if grid[rr][cc] == ".":
            grid[rr][cc] = "#"
            _, looped = walk(grid, r, c)
            if looped:
                print(f"found loop with obstactle at ({rr},{cc})")
                loops.add((rr, cc))
            grid[rr][cc] = "."

    return len(loops)


if __name__ == "__main__":
    assert (
        p1 := part_one(Input.from_file_relpath("example.txt"))
    ) == 41, f"Part one failed, {p1=}"
    assert (
        p2 := part_two(Input.from_file_relpath("example.txt"))
    ) == 6, f"Part two failed, {p2=}"
    Day(part_one, part_two).run()
