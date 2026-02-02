import sys

from utils import input


def main():
    total = 0
    for bank in input.lines():
        argsorted = sorted(range(len(bank)), key=lambda i: -int(bank[i]))
        print(argsorted)
        if argsorted[0] == len(bank) - 1:
            leftover = [i for i in argsorted if i > argsorted[1]]
            total += int(bank[argsorted[1]] + bank[leftover[0]])
        else:
            leftover = [i for i in argsorted if i > argsorted[0]]
            total += int(bank[argsorted[0]] + bank[leftover[0]])
    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
