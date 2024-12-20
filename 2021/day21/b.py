import sys
from functools import cache
from itertools import product

from utils import input


def parse_input() -> list[int]:
    data = input.lines()
    players = [int(line.split(":")[1].strip()) for line in data]
    return players


@cache
def count_wins(position_1, position_2, score_1, score_2):
    wins_1 = 0
    wins_2 = 0
    for roll_1, roll_2, roll_3 in product((1, 2, 3), repeat=3):
        new_position_1 = (position_1 - 1 + roll_1 + roll_2 + roll_3) % 10 + 1
        new_score_1 = score_1 + new_position_1
        if new_score_1 >= 21:
            wins_1 += 1
        else:
            new_wins_2, new_wins_1 = count_wins(
                position_2, new_position_1, score_2, new_score_1
            )
            wins_1 += new_wins_1
            wins_2 += new_wins_2
    return wins_1, wins_2


pos1, pos2 = parse_input()
wins = count_wins(pos1, pos2, 0, 0)
print(max(wins), file=sys.stderr)
