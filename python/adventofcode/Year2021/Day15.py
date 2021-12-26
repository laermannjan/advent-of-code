from dataclasses import dataclass

@dataclass
class Tile:
    row: int
    col: int
    risk: int

class RiskLevelMap:
    def __init__(self, risk_level_map: list[list[Tile]]):
        self._map = risk_level_map

    @property
    def height(self):
        return len(self._map)

    @property
    def width(self):
        return len(self._map[0])

    def __getitem__(self, pos: tuple[int, int]) -> Tile:
        return self._map[pos[0]][pos[1]]

    def get_neighbors(self, tile: Tile):
        neighbors = []
        row, col = tile.row, tile.col
        if row > 0:
            neighbors.append(self._map[row - 1][col])
        if row < self.height - 1:
            neighbors.append(self._map[row + 1][col])
        if col > 0:
            neighbors.append(self._map[row][col - 1])
        if col < self.width - 1:
            neighbors.append(self._map[row][col + 1])
        return neighbors

    def get_min_cost(self, start: Tile, end: Tile):
        cost = [[float("inf") for _ in range(self.width)] for _ in range(self.height)]
        cost[start.row][start.col] = 0

        # BFS
        queue = [start]
        while queue:
            current = queue.pop(0)
            for neighbor in self.get_neighbors(current):
                neighbor_cost = cost[neighbor.row][neighbor.col]
                cost_to_neighbor = cost[current.row][current.col] + neighbor.risk
                if neighbor_cost > cost_to_neighbor:
                    cost[neighbor.row][neighbor.col] = cost_to_neighbor
                    queue.append(neighbor)
        return cost[end.row][end.col]

    @classmethod
    def from_file(cls, filename: str, expand: int = 1) -> "RiskLevelMap":
        with open(filename) as f:
            data = f.read().splitlines()
        map = [
            [Tile(row, col, int(risk)) for col, risk in enumerate(row_data)]
            for row, row_data in enumerate(data)
        ]

        if expand > 1:
            height = len(map)
            width = len(map[0])

            expanded_map = [
                [None for _ in range(width * expand)]
                for _ in range(height * expand)
            ]
            for row in range(height * expand):
                for col in range(width * expand):
                    cell_to_copy = map[row % height][col % width]
                    offset = row // height + col // width
                    # risk loops around to 1 if it is greater than 9
                    risk = (cell_to_copy.risk + offset - 1) % 9 + 1
                    expanded_map[row][col] = Tile(row, col, risk)
            map = expanded_map
        return cls(map)



def day15a(filename: str) -> int:
    rlm = RiskLevelMap.from_file(filename)
    return int(rlm.get_min_cost(rlm[0, 0], rlm[-1, -1]))


def day15b(filename: str) -> int:
    rlm = RiskLevelMap.from_file(filename, expand=5)
    return int(rlm.get_min_cost(rlm[0, 0], rlm[-1, -1]))
