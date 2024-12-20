from utils import input

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

print(len(positions))
