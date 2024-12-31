import re
import sys

from utils import input


def main():
    total_cost = 0
    for g, game in enumerate(input.sections()):
        match = re.match(r".*X([+-]\d+), Y([+-]\d+)", game[0])
        Ax, Ay = int(match.group(1)), int(match.group(2))
        match = re.match(r".*X([+-]\d+), Y([+-]\d+)", game[1])
        Bx, By = int(match.group(1)), int(match.group(2))
        match = re.match(r".*X=(\d+), Y=(\d+)", game[2])
        Px = int(match.group(1)) + 10000000000000
        Py = int(match.group(2)) + 10000000000000

        print(f"game {g} - {Px} - {Py}")

        # We have two equations
        # Px = Ax * a + Bx * b  (Prize x coord, Button A x-delta, Button A presses, Button B x-delta, Button B presses)
        # Py = Ay * a + By * b  (Prize y coord, Button A y-delta, Button A presses, Button B y-delta, Button B presses)
        # we need to find `a` and `b`, that satify these equations, but both must be integers in [0, 100]
        # convert both formulas to slope form
        # b = (-Ax * a + Px) / Bx
        # b = (-Ay * a + Py) / By
        # set equal
        # (-Ax * a + Px) / Bx = (-Ay * a + Py) / By
        # solve for a
        # -(Ax/Bx) * a + Px/Bx = -(Ay/By) * a + Py/By
        # (Ay/By - Ax/Bx) * a = Py/By - Px/Bx
        # a = (Py/By - Px/Bx) / (Ay/By - Ax/Bx)
        # put into any of the two slope formulas
        # b = (-Ax * a + Px) / Bx

        a = (Py / By - Px / Bx) / (Ay / By - Ax / Bx)
        b = (-Ax * a + Px) / Bx

        a_round = round(a)
        b_round = round(b)

        if not (abs(a - a_round) < 1e-3 and abs(b - b_round) < 1e-3):
            print("  skip, no integer solution")
            # solution is not integer
            continue

        a, b = a_round, b_round
        cost = 3 * a + b
        print(f"game {g} - {a=}, {b=}, {cost=}")
        total_cost += cost

    print(total_cost, file=sys.stderr)


if __name__ == "__main__":
    main()
