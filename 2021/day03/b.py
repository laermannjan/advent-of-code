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


def get_sequence_by_index_freq(data: list[list[int]], most_common: bool) -> str:
    numbers = np.array(data, dtype=int)
    index = 0
    while numbers.shape[0] > 1:
        freq = numbers[:, index].mean()
        bit = freq >= 0.5 if most_common else freq < 0.5
        numbers = numbers[numbers[:, index] == bit]
        index += 1
    return "".join([str(digit) for digit in numbers.flatten()])


numbers = [[int(digit) for digit in row] for row in input.lines()]
oxygen = get_sequence_by_index_freq(numbers, most_common=True)
co2 = get_sequence_by_index_freq(numbers, most_common=False)
print(int(oxygen, 2) * int(co2, 2), file=sys.stderr)
