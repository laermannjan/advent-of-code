import sys
from functools import cache
from itertools import pairwise, product

from utils import input

numbers = """789\n456\n123\n 0A"""
arrows = " ^A\n<v>"


@cache
def compute_options(keypad: str):
    # computes all possible shortest sequences of button presses between any two buttons on the keypad
    keys = {
        (r, c): key
        for r, row in enumerate(keypad.splitlines())
        for c, key in enumerate(row)
        if key != " "
    }
    coords = {
        key: (r, c)
        for r, row in enumerate(keypad.splitlines())
        for c, key in enumerate(row)
        if key != " "
    }

    options = {}
    for start, end in product(coords, repeat=2):
        sr, sc = coords[start]
        er, ec = coords[end]
        length = abs(sr - er) + abs(sc - ec)

        # since manhattan distance, shortest paths are all combinations
        # of moves that don't move back and reach end
        seqs = set()
        moves = []
        if sr > er:
            moves.append((-1, 0, "^"))
        elif sr < er:
            moves.append((1, 0, "v"))
        if sc > ec:
            moves.append((0, -1, "<"))
        elif sc < ec:
            moves.append((0, 1, ">"))

        for move in product(moves, repeat=length):
            r, c = sr, sc
            seq = []
            for dr, dc, moves in move:
                r += dr
                c += dc
                if (r, c) not in keys:
                    break
                seq.append(moves)
            else:
                if coords[end] == (r, c):
                    seqs.add("".join(seq) + "A")
        options[(start, end)] = seqs

    return options


def compute_seqs(buttons: str, keypad: str) -> list[str]:
    move_to_seqs = compute_options(keypad)
    # all shortest seqs to get from one button to the next
    # a single seq contains a list of button presses on the keypad
    # e.g. v>vA  to get from 7 to 2
    seqs_per_move = [move_to_seqs[move] for move in pairwise("A" + buttons)]
    # since there might be multiple options to get from any button to the next
    # get the product of sub-sequences
    all_seqs = ["".join(seq) for seq in product(*seqs_per_move)]
    return all_seqs


def main():
    keypads = [numbers, arrows, arrows]

    total = 0
    for line in input.lines():
        seqs = [line]
        for keypad in keypads:
            seqs = [
                next_level_buttons
                for buttons in seqs
                for next_level_buttons in compute_seqs(buttons, keypad)
            ]
            minlen = min(map(len, seqs))
            seqs = [s for s in seqs if len(s) == minlen]

        total += minlen * int(line[:-1])

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
