import sys

from utils import input


def main():
    total = 0

    for bank in input.lines():
        bats = list(map(int, bank))
        indices = list(range(len(bats)))
        sorted_indices = sorted(indices, key=lambda i: -bats[i])

        bats_left = 2
        min_index = -1

        for i in reversed(range(bats_left)):
            # digit = max(bats[min_index + 1 : len(bats) - i])
            max_index = next(
                (x for x in sorted_indices if min_index < x < len(bats) - i)
            )
            digit = bats[max_index]
            min_index = max_index
            total += digit * (10**i)

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
