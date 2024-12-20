import sys

import numpy as np
from utils import input


def parse_input() -> tuple[list[int], list[np.ndarray]]:
    data = input.lines()
    draws = [int(draw) for draw in data[0].split(",")]
    boards = [
        [[int(digit) for digit in row.split()] for row in data[i : i + 5]]
        for i in range(2, len(data), 6)
    ]
    return draws, [np.array(b) for b in boards]


draws, boards = parse_input()

for draw in draws:
    for board in boards:
        board[board == draw] = 0
        if 0 in board.sum(axis=0) or 0 in board.sum(axis=1):
            print(board.sum() * draw, file=sys.stderr)
            exit(0)
