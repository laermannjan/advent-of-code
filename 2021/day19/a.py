import sys
from collections import Counter
from itertools import combinations

from utils import input


class ScannedSpace:
    def __init__(self, scanner_reports: list[list[tuple[int, int, int]]]):
        self.unaligned_scans = scanner_reports[1:]
        self.scanners = {(0, 0, 0)}
        self.beacons: set[tuple[int, int, int]] = set(scanner_reports[0])

    @classmethod
    def beacon_rotations(cls, x: int, y: int, z: int) -> list[tuple[int, int, int]]:
        return [
            rotation
            for rotation in [
                (x, y, z),
                (-y, x, z),
                (-x, -y, z),
                (y, -x, z),
                (x, -z, y),
                (z, x, y),
                (-x, z, y),
                (-z, -x, y),
                (x, -y, -z),
                (y, x, -z),
                (-x, y, -z),
                (-y, -x, -z),
                (x, z, -y),
                (-z, x, -y),
                (-x, -z, -y),
                (z, -x, -y),
                (-z, y, x),
                (-y, -z, x),
                (z, -y, x),
                (y, z, x),
                (-y, z, -x),
                (-z, -y, -x),
                (y, -z, -x),
                (z, y, -x),
            ]
        ]

    @classmethod
    def scan_rotations(
        cls, scan: list[tuple[int, int, int]]
    ) -> list[list[tuple[int, int, int]]]:
        # get all rotations per beacon, then zip to group by scanner_report
        return [*zip(*[cls.beacon_rotations(*beacon) for beacon in scan])]

    @classmethod
    def shift_scan(cls, scan, dx, dy, dz):
        return [(x + dx, y + dy, z + dz) for (x, y, z) in scan]

    @property
    def max_scanner_distance(self):
        return max(
            sum(abs(c1 - c2) for c1, c2 in zip(s1, s2))
            for s1, s2 in combinations(self.scanners, 2)
        )

    def find_scanner(
        self, scan: list[tuple[int, int, int]]
    ) -> tuple[tuple[int, int, int], list[tuple[int, int, int]]] | tuple[None, None]:
        """
        Tries to find the scanner by matching a scan with the set of known beacons.
        Returns scanner position and adjusted beacon scans if successful; None, None otherwise.

        If more than 12 beacons have the same shift to an entry in the scan report,
        this must also be the shift of the unknown scanner w.r.t. (0, 0, 0).
        """
        for rotated_scan in self.scan_rotations(scan):
            shift, count = Counter(
                (x2 - x1, y2 - y1, z2 - z1)
                for x1, y1, z1 in rotated_scan
                for x2, y2, z2 in self.beacons
            ).most_common(1)[0]
            if count >= 12:
                return shift, self.shift_scan(rotated_scan, *shift)

        return None, None

    def align_scanners(self):
        while self.unaligned_scans:
            scan = self.unaligned_scans.pop(0)
            scanner, aligned_beacons = self.find_scanner(scan)
            if scanner:
                self.scanners.add(scanner)
                self.beacons |= set(aligned_beacons)
            else:
                # maybe we can figure this one out later with more fixed beacons
                self.unaligned_scans.append(scan)

    @classmethod
    def from_file(cls):
        scanner_reports = input.stdin().read().split("\n\n")
        return cls(
            [
                [
                    tuple(map(int, beacon.split(",")))
                    for beacon in scanner_report.splitlines()[1:]
                ]
                for scanner_report in scanner_reports
            ]
        )


scanner_space = ScannedSpace.from_file()
scanner_space.align_scanners()
print(len(scanner_space.beacons), file=sys.stderr)
