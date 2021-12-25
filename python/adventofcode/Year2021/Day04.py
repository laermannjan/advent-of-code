"""
--- Day 4: Giant Squid ---
You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.
Maybe it wants to play bingo?
Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)
The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6
14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):
22 13 17  X  0         3 15  0  2 22        14 21 17 24  X
 8  2 23  X 24         X 18 13 17  X        10 16 15  X 19
21  X 14 16  X        19  8  X 25 23        18  8 23 26 20
 6 10  3 18  X        20  X 10 24  X        22  X 13  6  X
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  X
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:
22 13  X  X  0         3 15  0  X 22         X  X  X 24  X
 8  X  X  X 24         X 18 13  X  X        10 16 15  X 19
 X  X  X 16  X        19  8  X 25  X        18  8  X 26 20
 6 10  3 18  X        20  X 10 24  X        22  X 13  6  X
 1 12 20 15 19         X  X 16 12  6         X  0 12  3  X
Finally, 24 is drawn:
22 13  X  X  0         3 15  0  X 22         X  X  X  X  X
 8  X  X  X  X         X 18 13  X  X        10 16 15  X 19
 X  X  X 16  X        19  8  X 25  X        18  8  X 26 20
 6 10  3 18  X        20  X 10  X  X        22  X 13  6  X
 1 12 20 15 19         X  X 16 12  6         X  0 12  3  X
At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).
The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.
To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?
"""
import numpy as np


def parse_input(filename: str) -> tuple[list[int], list[np.ndarray]]:
    with open(filename, 'r') as f:
        data = f.read().splitlines()
    draws = [int(draw) for draw in data[0].split(",")]
    boards = [
        [
            [int(digit) for digit in row.split()]
            for row in data[i:i+5]]
        for i in range(2, len(data), 6)
    ]
    return draws, [np.array(b) for b in boards]


def day04a(filename: str):
    draws, boards = parse_input(filename)

    for draw in draws:
        for board in boards:
            board[board == draw] = 0
            if 0 in board.sum(axis=0) or 0 in board.sum(axis=1):
                return board.sum() * draw


def day04b(filename: str):
    draws, boards = parse_input(filename)
    for draw in draws:
        winning_boards = []
        for i, board in enumerate(boards):
            board[board == draw] = 0
            if 0 in board.sum(axis=0) or 0 in board.sum(axis=1):
                if len(boards) == 1:
                    return board.sum() * draw
                # boards.pop(i)
                winning_boards.append(i)
        boards = [board for i, board in enumerate(boards) if i not in winning_boards]
