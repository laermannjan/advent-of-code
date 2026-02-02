import sys

from utils import input


def main():
    dial = 50
    zeros = 0
    rotations = [int(line.replace("L", "-").replace("R", "")) for line in input.lines()]

    print(rotations)
    for rot in rotations:
        dial = (dial + rot) % 100
        print(dial)
        if dial == 0:
            zeros += 1
    print(zeros, file=sys.stderr)


if __name__ == "__main__":
    main()
