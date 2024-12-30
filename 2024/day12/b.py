import sys

from utils import input


def region(
    grid: dict[tuple[int, int], str],
    start: tuple[int, int],
):
    area = set()
    stack = [start]
    while stack:
        row, col = stack.pop()
        plant = grid[(row, col)]
        area.add((row, col))

        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = row + dr, col + dc
            if (nr, nc) not in grid:
                continue
            if (nr, nc) in area:
                continue
            if grid[(nr, nc)] == plant:
                stack.append((nr, nc))
    return area


def sides(area):
    # instead of counting edges, and finding line segments as below,
    # we can also count corners of the polygon that describes the area
    # each coord of a plant cell is integer based, so we represent corners of
    # such a cell as the (+/-0.5, +/-0.5) to that cell
    # applying the same delta to a cell corner, gives us the cell coords of
    # the four neighboring cells that this corner touches
    # if a cell corner has 1 or 3 neighbor cells, that are part of the area
    # it's a corner of the polygon / area
    # if it has 2 such neighbors and they are diagnoal to each other,
    # the cell corner point represent two polygon corners:
    # "X" is a cell that is part of the area, "+" is the current cell corner
    #
    # X|        X|        X|X
    # -+-       -+-       -+-
    #  |         |X       X|
    #

    deltas = [(-0.5, -0.5), (-0.5, 0.5), (0.5, 0.5), (0.5, -0.5)]
    cell_corners = {(r + dr, c + dc) for r, c in area for dr, dc in deltas}

    polygon_corners = 0
    for r, c in cell_corners:
        neighbors = [(r + dr, c + dc) in area for dr, dc in deltas]
        n_neighbors = sum(neighbors)
        if n_neighbors in (1, 3):
            print(neighbors)
            polygon_corners += 1
        elif n_neighbors == 2 and neighbors in (
            [True, False, True, False],
            [False, True, False, True],
        ):
            print(neighbors, "double")
            polygon_corners += 2
    return polygon_corners


def main():
    grid = dict(input.coords())
    covered = set()

    price = 0

    for (row, col), plant in grid.items():
        if (row, col) in covered:
            continue
        area = region(grid, (row, col))
        covered |= area

        s = sides(area)
        print(f"{plant=}, {area=}, {s=}")
        price += len(area) * s

    print(price, file=sys.stderr)


# def main():
#     grid = dict(input.coords())
#     covered = set()
#
#     price = 0
#
#     for (row, col), plant in grid.items():
#         if (row, col) in covered:
#             continue
#         area = region(grid, (row, col))
#         covered |= area
#
#         perimiter = {}
#
#         # get the plant coords on the perimiter
#         # store them with the direction towards which the edge is
#         # ... kind of a normal vector, because the same plant can have multiple edges, one to either side
#         for arow, acol in area:
#             for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
#                 nr, nc = arow + dr, acol + dc
#                 if (nr, nc) not in area:
#                     if (dr, dc) not in perimiter:
#                         perimiter[(dr, dc)] = []
#                     perimiter[(dr, dc)].append((arow, acol))
#
#         print(f"found region {plant=}, {area=}, {perimiter=}")
#         total_sides = 0
#         for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
#             normal_sides = 0
#             # iterate over edge-coords with the same normal vector
#             # a "side" only has one degree of freedom in the axis that is 0 in the normal vector
#             # basically, sort first by the direction of the normal vector, then by the remaining one
#             # also swap the order of row/col such that the normal direction is first, and the other one second
#             # (I know I suck at explaining this to myself)
#             #
#             # We know a perimiter node cannot be on the same "side" as the previous one
#             # if the coord on the axis the normal vector is pointing at has changed (think: horizontal edge (left to right) but different height value)
#             # the other coord must always be +1 to the last, otherwise there is a gap and it must be a new side
#             # to make this logic work for all normal vectors, we swap the row/col order in the tuple on the next line
#             perim_plants = sorted(
#                 [(c, r) if dr == 0 else (r, c) for r, c in perimiter[(dr, dc)]]
#             )
#             print(f"iterating perimiter: {perim_plants}")
#             if not perim_plants:
#                 continue
#             last_x, last_y = perim_plants[0]
#             normal_sides += 1
#             for x, y in perim_plants[1:]:
#                 if x != last_x or y != last_y + 1:
#                     normal_sides += 1
#                 last_x, last_y = x, y
#             print(f"{normal_sides=}")
#             total_sides += normal_sides
#
#         print(f"{total_sides=}")
#         price += len(area) * total_sides
#
#     print(price, file=sys.stderr)
#

if __name__ == "__main__":
    main()
