import sys

import numpy as np

from utils import input


def parse_input(filename: str) -> list[list[int]]:
    with open(filename, "r") as f:
        data = f.read().splitlines()
    return [[int(digit) for digit in row] for row in data]


def line2bits(line: str) -> list[int]:
    return [int(x) for x in line]


def get_bit_fractions(lines: list[str]) -> np.ndarray:
    total = np.array(line2bits(lines[0]))
    for line in lines[1:]:
        total += np.array(line2bits(line))
    return total / len(lines)


numbers = [[int(digit) for digit in row] for row in input.lines()]
freqs = np.array(numbers).mean(axis=0)
most_common_bits = list(freqs >= 0.5)

gamma = int("".join([str(int(bit)) for bit in most_common_bits]), 2)
epsilon = int("".join([str(int(not bit)) for bit in most_common_bits]), 2)

print(gamma * epsilon, file=sys.stderr)
