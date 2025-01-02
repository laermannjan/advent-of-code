import sys

from utils import input


def main():
    graph = {k: v for k, v in dict(input.coords()).items()}
    for (r, c), tile in graph.items():
        if tile == "S":
            sr, sc = r, c
            break

    pq = [(0, sr, sc, 0, 1)]  # cost, start_row, start_col, move_dir_row, move_dir_col
    visisted = {}

    while pq:
        # NOTE: heapq would be faster..., just started out with a list
        dist, r, c, dr, dc = pq.pop(pq.index(min(pq)))
        visisted[(r, c)] = dist
        if graph[(r, c)] == "E":
            print(dist, file=sys.stderr)
            break
        for ndist, ndr, ndc in [
            (dist + 1, dr, dc),
            (dist + 1001, -dc, dr),
            (dist + 1001, dc, -dr),
        ]:
            nr, nc = r + ndr, c + ndc
            if graph[(nr, nc)] == "#":
                continue
            if (nr, nc) in visisted:
                continue
            pq.append((ndist, nr, nc, ndr, ndc))

    # min_dist = dist
    # path = [(r, c, "E")]
    # while (r, c) != (sr, sc):
    #     for nr, nc in [(r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)]:
    #         if (dist := visisted.get((nr, nc), float("inf"))) < min_dist:
    #             min_dist = dist
    #             if graph[(nr, nc)] == "S":
    #                 tile = "S"
    #             elif nr - r == -1:
    #                 tile = "v"
    #             elif nr - r == 1:
    #                 tile = "^"
    #             elif nc - c == -1:
    #                 tile = ">"
    #             elif nc - c == 1:
    #                 tile = "<"
    #             r, c = nr, nc
    #     path.append((r, c, tile))
    #
    # grid = [
    #     [graph[(r, c)] for c in range(max(graph)[1] + 1)]
    #     for r in range(max(graph)[0] + 1)
    # ]
    # for pr, pc, tile in path:
    #     grid[pr][pc] = tile
    #
    # for row in grid:
    #     print("".join(row))


if __name__ == "__main__":
    main()
