import sys

import numpy as np
from utils import input


def sum2n(n: int) -> int:
    return int((n + 1) * (n / 2))


crabs = [int(x) for x in input.lines()[0].split(",")]
median = np.median(crabs)
costs = [abs(crab - median) for crab in crabs]
cost = sum(costs)
print(int(cost), file=sys.stderr)
