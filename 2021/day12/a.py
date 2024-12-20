import sys
from collections import defaultdict

from utils import input


class Graph:
    def __init__(self) -> None:
        self._caves: defaultdict[str, list[str]] = defaultdict(list)
        for line in input.lines():
            cave_a, cave_b = line.split("-")
            self.add_connection(cave_a, cave_b)

    def add_connection(self, cave_a: str, cave_b: str):
        self._caves[cave_a].append(cave_b)
        self._caves[cave_b].append(cave_a)

    def is_big(self, cave: str) -> bool:
        return cave.isupper()

    def get_connected_caves(self, cave: str) -> list[str]:
        return self._caves[cave]

    def count_paths(self) -> int:
        return self._count_paths(current="start", visited={"start"})

    def _count_paths(self, current: str, visited: set) -> int:
        raise NotImplementedError


class GraphA(Graph):
    def _count_paths(self, current: str, visited: set) -> int:
        if current == "end":
            return 1
        return sum(
            self._count_paths(connected_cave, visited | {connected_cave})
            for connected_cave in self.get_connected_caves(current)
            if self.is_big(connected_cave) or connected_cave not in visited
        )


class GraphB(Graph):
    def _count_paths(
        self, current: str, visited: set, can_visit_twice: bool = True
    ) -> int:
        if current == "end":
            return 1

        count = 0
        for connected_cave in self.get_connected_caves(current):
            if self.is_big(connected_cave) or connected_cave not in visited:
                count += self._count_paths(
                    connected_cave, visited | {connected_cave}, can_visit_twice
                )
            elif can_visit_twice and connected_cave not in ("start", "end"):
                count += self._count_paths(
                    connected_cave, visited | {connected_cave}, False
                )
        return count


graph = GraphA()
print(graph.count_paths(), file=sys.stderr)
