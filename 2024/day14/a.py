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
    seconds = 100
    q = [0, 0, 0, 0]
    for line in input.lines():
        px, py, vx, vy = map(int, re.findall(r"\-?\d+", line))

        x = (px + seconds * vx) % WIDTH
        y = (py + seconds * vy) % HEIGHT

        if x < WMID and y < HMID:
            q[0] += 1
            print(x, y, "q0")
        elif x > WMID and y < HMID:
            q[1] += 1
            print(x, y, "q1")
        elif x < WMID and y > HMID:
            q[2] += 1
            print(x, y, "q2")
        elif x > WMID and y > HMID:
            q[3] += 1
            print(x, y, "q3")
        else:
            print(x, y, "on edge")

    print(f"{q}")
    safety = 1
    for v in q:
        safety *= v
    print(safety, file=sys.stderr)


if __name__ == "__main__":
    main()
