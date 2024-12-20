import os
import sys
import time
from functools import total_ordering

from colored import attr, bg, fg
from utils import input
from utils.misc import memoized_property


@total_ordering
class Tile:
    def __init__(self, row: int, col: int, risk: int):
        self.row = row
        self.col = col
        self.risk = risk

    def __lt__(self, other):
        if self.risk == float("inf") and other.risk == float("inf"):
            return self.row**2 + self.col**2 < other.row**2 + other.col**2
        return self.risk < other.risk

    def __eq__(self, other):
        return self is other

    def __repr__(self):
        return f"Tile({self.row},{self.col})[{self.risk}]"


class RiskLevelMap:
    def __init__(self, risk_level_map: list[list[Tile]]):
        self._map = risk_level_map
        self._cost = [
            [Tile(row, col, float("inf")) for col in range(self.width)]
            for row in range(self.height)
        ]

    @property
    def height(self):
        return len(self._map)

    @property
    def width(self):
        return len(self._map[0])

    def __getitem__(self, pos: tuple[int, int]) -> Tile:
        return self._map[pos[0]][pos[1]]

    def __hash__(self):
        return hash(
            [tile.risk for row in self._map for tile in row]
            + [tile.risk for row in self._cost for tile in row]
        )

    def get_neighbors(self, tile: Tile, use_cost=False):
        map = self._cost if use_cost else self._map
        neighbors = []
        row, col = tile.row, tile.col
        if row > 0:
            neighbors.append(map[row - 1][col])
        if row < self.height - 1:
            neighbors.append(map[row + 1][col])
        if col > 0:
            neighbors.append(map[row][col - 1])
        if col < self.width - 1:
            neighbors.append(map[row][col + 1])
        return neighbors

    def get_min_cost(self, start: Tile, end: Tile):
        self._cost[start.row][start.col].risk = 0
        if int(os.getenv("ADVENTOFCODE_VISUALIZE", 0)):
            os.system("cls" if os.name == "nt" else "clear")
        # BFS
        queue = [start]
        while queue:
            current = queue.pop(0)
            for neighbor in self.get_neighbors(current):
                neighbor_cost = self._cost[neighbor.row][neighbor.col].risk
                cost_to_neighbor = (
                    self._cost[current.row][current.col].risk + neighbor.risk
                )
                if neighbor_cost > cost_to_neighbor:
                    self._cost[neighbor.row][neighbor.col].risk = cost_to_neighbor
                    queue.append(neighbor)
                    if int(os.getenv("ADVENTOFCODE_VISUALIZE", 0)):
                        print(f"\033[{0};{0}H", end="")
                        print(self.heatmap())
                        time.sleep(0.1)
        return self._cost[end.row][end.col].risk

    def heatmap(self) -> str:
        """Return a colored heatmap of a matrix, lowest numbers being blue, highest being red"""
        output = ""
        # colors representing medium blue, sky blue, teal, green, yellow, yellow-orange, orange, red-orange, red
        foregrounds = [25, 39, 6, 119, 3, 214, 166, 124]
        min_num = min(min(row) for row in self._cost).risk
        max_num = max(
            max(num.risk if num.risk < float("inf") else min_num for num in row)
            for row in self._cost
        )

        for row in range(self.height):
            for col in range(self.width):
                tile = self._cost[row][col]
                background = 255 if tile in self.get_shortest_path else 0
                if tile.risk == float("inf"):
                    output += "  ∞"
                    continue
                color = foregrounds[
                    int((tile.risk - min_num) / (max_num - min_num) * 7)
                ]
                output += f"{fg(color)}{bg(background)}{tile.risk:>3}{attr('reset')}"
            output += "\n"
        return output

    @memoized_property
    def get_shortest_path(self):
        tile = self._cost[-1][-1]
        path = [tile]
        while tile is not self._cost[0][0]:
            neighbors = self.get_neighbors(tile, use_cost=True)
            tile = min(neighbors)
            path.append(tile)
        return path

    @classmethod
    def from_file(cls, scale_factor: int = 1) -> "RiskLevelMap":
        data = input.lines()
        map = [
            [Tile(row, col, int(risk)) for col, risk in enumerate(row_data)]
            for row, row_data in enumerate(data)
        ]

        if scale_factor > 1:
            height = len(map)
            width = len(map[0])

            scaled_map = [
                [Tile(row, col, None) for col in range(width * scale_factor)]
                for row in range(height * scale_factor)
            ]
            for row in range(height * scale_factor):
                for col in range(width * scale_factor):
                    cell_to_copy = map[row % height][col % width]
                    offset = row // height + col // width
                    # risk loops around to 1 if it is greater than 9
                    risk = (cell_to_copy.risk + offset - 1) % 9 + 1
                    scaled_map[row][col].risk = risk
            map = scaled_map
        return cls(map)


rlm = RiskLevelMap.from_file()
print(int(rlm.get_min_cost(rlm[0, 0], rlm[-1, -1])), file=sys.stderr)
