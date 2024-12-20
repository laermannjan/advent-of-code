import sys

import numpy as np
from utils import input


def sum2n(n: int) -> int:
    return int((n + 1) * (n / 2))


crabs = [int(x) for x in input.lines()[0].split(",")]
floored_mean = np.floor(np.mean(crabs))
ceiled_mean = np.ceil(np.mean(crabs))

floored_costs = [sum2n(abs(crab - int(floored_mean))) for crab in crabs]
ceiled_costs = [sum2n(abs(crab - int(ceiled_mean))) for crab in crabs]
cost = min(sum(floored_costs), sum(ceiled_costs))
print(int(cost), file=sys.stderr)
