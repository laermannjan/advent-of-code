import re
import sys
from dataclasses import dataclass

from utils import input


@dataclass
class Cuboid:
    is_on: bool
    x_min: int
    x_max: int
    y_min: int
    y_max: int
    z_min: int
    z_max: int

    @property
    def space(self):
        space = (
            (self.x_max - self.x_min)
            * (self.y_max - self.y_min)
            * (self.z_max - self.z_min)
        )
        return space if self.is_on else -space

    def __sub__(self, other):
        """Intersection between two cuboids."""
        if not isinstance(other, Cuboid):
            return NotImplemented
        if (
            self.x_min < other.x_max
            and self.x_max > other.x_min
            and self.y_min < other.y_max
            and self.y_max > other.y_min
            and self.z_min < other.z_max
            and self.z_max > other.z_min
        ):
            return Cuboid(
                not self.is_on,
                self.x_min if other.x_min <= self.x_min < other.x_max else other.x_min,
                self.x_max if other.x_min <= self.x_max < other.x_max else other.x_max,
                self.y_min if other.y_min <= self.y_min < other.y_max else other.y_min,
                self.y_max if other.y_min <= self.y_max < other.y_max else other.y_max,
                self.z_min if other.z_min <= self.z_min < other.z_max else other.z_min,
                self.z_max if other.z_min <= self.z_max < other.z_max else other.z_max,
            )

    @classmethod
    def from_string(cls, string: str):
        x_min, x_max, y_min, y_max, z_min, z_max = (
            int(n) for n in re.findall(r"-?\d+", string)
        )
        return cls(
            string.startswith("on"),
            x_min,
            x_max + 1,
            y_min,
            y_max + 1,
            z_min,
            z_max + 1,
        )


def reboot_reactor(cuboids):
    """
    All "space" in the reactor is initially negative.
    Thus, we only need to consider cuboids that are turned on and the intersections of turned off cuboids
    with the entire turned-on volume.
    Additionally, we need to keep track of intersections of newly added cuboids and the volume to not
    count double.
    """
    all_cuboids = []
    for cuboid in cuboids:
        all_cuboids.extend(
            [
                intersection
                for other in all_cuboids
                if (intersection := other - cuboid) is not None
            ]
        )
        if cuboid.is_on:
            all_cuboids.append(cuboid)
    return sum(cuboid.space for cuboid in all_cuboids)


cuboids = [Cuboid.from_string(line) for line in input.lines()]
space = reboot_reactor(cuboids)
print(space, file=sys.stderr)
