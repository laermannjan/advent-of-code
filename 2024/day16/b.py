import heapq
import sys

from utils import input


def main():
    graph = {k: v for k, v in dict(input.coords()).items()}
    for (r, c), tile in graph.items():
        if tile == "S":
            sr, sc = r, c
            break

    pq = [(0, sr, sc, 0, 1)]  # cost, start_row, start_col, move_dir_row, move_dir_col
    lowest_dists = {}
    backtrack = {}
    best_end_dist = float("inf")
    end_states = set()

    while pq:
        # NOTE: a list is not nearly fast enough. without the heap it would take ages
        # dist, r, c, dr, dc = pq.pop(pq.index(min(pq)))
        dist, r, c, dr, dc = heapq.heappop(pq)
        # print(f"{r=}, {c=}, {dr=}, {dc=}, {lowest_dists.get((r, c, dr, dc))=}")
        if dist > lowest_dists.get((r, c, dr, dc), float("inf")):
            continue
        if graph[(r, c)] == "E":
            if dist > best_end_dist:
                # there can't be any other better paths
                break
            best_end_dist = dist
            end_states.add((r, c, dr, dc))
        for ndist, nr, nc, ndr, ndc in [
            (dist + 1, r + dr, c + dc, dr, dc),
            (dist + 1001, r, c, -dc, dr),
            (dist + 1001, r, c, dc, -dr),
        ]:
            if graph[(nr, nc)] == "#":
                continue
            lowest_dist = lowest_dists.get((nr, nc, ndr, ndc), float("inf"))
            if ndist > lowest_dist:
                continue
            if ndist < lowest_dist:
                backtrack[(nr, nc, ndr, ndc)] = set()
                lowest_dists[(nr, nc, ndr, ndc)] = ndist
            backtrack[(nr, nc, ndr, ndc)].add((r, c, dr, dc))
            # pq.append((ndist, nr, nc, ndr, ndc))
            heapq.heappush(pq, (ndist, nr, nc, ndr, ndc))

    q = list(end_states)
    seen = set(end_states)
    while q:
        state = q.pop()
        for bt in backtrack.get(state, []):
            if bt in seen:
                continue
            seen.add(bt)
            q.append(bt)
    print(len({(r, c) for r, c, _, _ in seen}), file=sys.stderr)

    grid = [
        [graph[(r, c)] for c in range(max(graph)[1] + 1)]
        for r in range(max(graph)[0] + 1)
    ]
    for pr, pc, _, _ in seen:
        grid[pr][pc] = "O"

    for row in grid:
        print("".join(row))


if __name__ == "__main__":
    main()
