import sys

from utils import input


def main():
    dial = 50
    zeros = 0
    rotations = [int(line.replace("L", "-").replace("R", "")) for line in input.lines()]

    print(rotations)
    for rot in rotations:
        print(dial, rot)

        if rot < 0:
            if dial == 0:
                clicks = -rot // 100
            else:
                clicks = (100 - dial - rot) // 100
        elif rot > 0:
            clicks = (dial + rot) // 100
        print(f" ..... {clicks} clicks")
        zeros += clicks
        dial = (dial + rot) % 100

    print(dial)
    print(zeros, file=sys.stderr)


if __name__ == "__main__":
    main()
