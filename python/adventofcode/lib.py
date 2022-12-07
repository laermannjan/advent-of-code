from typing import Callable
from math import sqrt


class memoized_property:
    def __init__(self, factory: Callable):
        self._factory = factory
        self._value = None
        self._cached_instance_hash = None

    def __get__(self, instance: object, owner: type):
        instance_hash = instance.__hash__()
        if instance_hash != self._cached_instance_hash:
            self._cached_instance_hash = instance_hash
            self._value = self._factory(instance)
        return self._value


def sum_to(n: int) -> float:
    if n < 0:
        raise ValueError(f"Cannot sum from 1 up to n, for {n=}")
    return (n + 1) * (n / 2)


def sum_to_inverse(s: int) -> float:
    """
    Inverse to the sum_to function.

    Given a sum S=1+2+...+n, returns n.
    If S was not produced by such a series, use floor/ceil to find closest such sums.

    Derivation:
    S = 1+2+...+n = (n+1)(n/2)
    2*S = n(n+1) = n**2 + n = (n + 1/2)**2 - 1/4   [from binomial formula; (a+b)**2 = a**2 + 2ab + b**2]
    2*S + 1/4 = (n + 1/2)**2
    sprt(2*S + 1/4) - 1/2 = n
    """
    return sqrt(2 * s - 0.25) - 0.5


# k = (n + 1) * (n / 2) =
#
# 2k = z = n**2 + n
#
# 0 = n**2 + n - z
#
# n = -0.5 +/- sqrt(0.25 - z)
#
# 0 = (-0.5 +/- sqrt(0.25 - z)) ** 2 + (-0.5 +/- sqrt(0.25 - z)) - z
#
# 0 =
