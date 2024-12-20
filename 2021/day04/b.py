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
    winning_boards = []
    for i, board in enumerate(boards):
        board[board == draw] = 0
        if 0 in board.sum(axis=0) or 0 in board.sum(axis=1):
            if len(boards) == 1:
                print(board.sum() * draw, file=sys.stderr)
                exit(0)
            # boards.pop(i)
            winning_boards.append(i)
    boards = [board for i, board in enumerate(boards) if i not in winning_boards]
