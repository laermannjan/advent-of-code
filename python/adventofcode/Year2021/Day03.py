import numpy as np


def line2bits(line: str) -> list[int]:
    return [int(x) for x in line]


def get_bit_fractions(lines: list[str]) -> np.ndarray:
    total = np.array(line2bits(lines[0]))
    for line in lines[1:]:
        total += np.array(line2bits(line))
    return total / len(lines)


def day03a(lines: list[str]) -> int:
    total = np.array(line2bits(lines[0]))
    for line in lines[1:]:
        total += np.array(line2bits(line))

    fractions = total / len(lines)
    most_common_bits = list(fractions >= 0.5)

    gamma = int("".join([str(int(bit)) for bit in most_common_bits]), 2)
    epsilon = int("".join([str(int(not bit)) for bit in most_common_bits]), 2)

    return gamma * epsilon


def get_sequence_by_element_freq(numbers: list[str], most_common: bool) -> str:
    index = 0
    while len(numbers) > 1:
        bits_at_index = [int(number[index]) for number in numbers]
        fractions_at_index = sum(bits_at_index) / len(numbers)
        bit = fractions_at_index >= 0.5 if most_common else fractions_at_index < 0.5
        numbers = [number for number in numbers if number[index] == str(int(bit))]
        index += 1
    return numbers[0]


def day03b(lines: list[str]) -> int:
    numbers = [line for line in lines]
    oxygen = get_sequence_by_element_freq(numbers, most_common=True)
    co2 = get_sequence_by_element_freq(numbers, most_common=False)
    return int(oxygen, 2) * int(co2, 2)
