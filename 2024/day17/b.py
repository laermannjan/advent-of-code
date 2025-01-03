import re
import sys

from utils import input

# real input...
# Program: 2,4,1,3,7,5,1,5,0,3,4,2,5,5,3,0
#
#
# 2,4 -> b = a % 8
# 1,3 -> b = b ^ 3
# 7,5 -> c = a >> b
# 1,5 -> b = b ^ 5
# 0,3 -> a = a >> 3   # div by 8
# 4,2 -> b = b ^ c
# 5,5 -> output b % 8
# 3,0 -> loop again from start
#
#
# - last output is 0 -> last loop must start with 0 <= a < 8, because anything >= 8, would mean a >> 3 is not 0, and we would loop again


def find(prog, end_of_single_iter_a):
    if prog == []:
        return end_of_single_iter_a
    for t in range(8):
        # in the real program, a = a >> 3
        # therefore any initial a' = (a << 3) + some 0 <= t < 8
        # will lead to the same a (since the >> 3 will truncate the last 3 bits)
        a = (end_of_single_iter_a << 3) + t
        b = a % 8
        b = b ^ 3
        c = a >> b
        b = b ^ 5
        b = b ^ c
        print(f"{t=} {a=} {b=} - {prog[-1]}")
        if b % 8 == prog[-1]:
            print("found partial solution")
            sub = find(prog[:-1], a)
            if sub is None:
                print("no sub")
                continue
            return sub


def main():
    _, _, _, *program = map(int, re.findall(r"\d+", input.stdin().read()))
    print(find(program, 0), file=sys.stderr)


if __name__ == "__main__":
    main()
