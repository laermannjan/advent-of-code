import re
import sys

from utils import input

WIDTH = 101
HEIGHT = 103

# WIDTH = 11
# HEIGHT = 7

WMID = WIDTH // 2
HMID = HEIGHT // 2


def main():
    config = [list(map(int, re.findall(r"\-?\d+", line))) for line in input.lines()]
    seconds = 10000  # just some large number

    min_safety, best_s = float("inf"), 0
    for s in range(seconds):
        q = [0, 0, 0, 0]
        print("SECONDS", s)

        for px, py, vx, vy in config:
            x = (px + s * vx) % WIDTH
            y = (py + s * vy) % HEIGHT
            if x < WMID and y < HMID:
                q[0] += 1
            elif x > WMID and y < HMID:
                q[1] += 1
            elif x < WMID and y > HMID:
                q[2] += 1
            elif x > WMID and y > HMID:
                q[3] += 1

        # The idea is that the "christmas tree" appears in one of the quadrants
        # and therefore most of the robots concentrate in a single quadrant
        # that means the product of the quadrant counts will be lowest, when
        # the christmas tree appears
        safety = q[0] * q[1] * q[2] * q[3]
        if safety < min_safety:
            min_safety = safety
            best_s = s

    print(best_s, file=sys.stderr)


if __name__ == "__main__":
    main()
