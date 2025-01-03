import re
import sys

from utils import input


def main():
    a, b, c, *program = map(int, re.findall(r"\d+", input.stdin().read()))

    pointer = 0
    output = []

    print(a, b, c, program)

    def combo(operand):
        if 0 <= operand <= 3:
            return operand
        if operand == 4:
            return a
        if operand == 5:
            return b
        if operand == 6:
            return c
        raise ValueError("invalid combo operand", operand)

    while pointer < len(program):
        inst = program[pointer]
        operand = program[pointer + 1]
        print(f"{inst=} {operand=}")

        if inst == 0:  # adv
            a = a >> combo(operand)  # `>> x` is equal to // 2**x
        elif inst == 1:  # bxl
            b = b ^ operand
        elif inst == 2:  # bst
            b = combo(operand) % 8
        elif inst == 3:  # jnz
            if a != 0:
                pointer = operand
                continue
        elif inst == 4:  # bxc
            b = b ^ c
        elif inst == 5:  # out
            output.append(combo(operand) % 8)
        elif inst == 6:  # bdv
            b = a >> combo(operand)
        elif inst == 7:  # cdv
            c = a >> combo(operand)

        pointer += 2

    print(",".join(map(str, output)), file=sys.stderr)


if __name__ == "__main__":
    main()
