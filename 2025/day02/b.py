import sys

from utils import input


def main():
    ranges = [list(map(int, r.split("-"))) for r in input.stdin().read().split(",")]
    numbers = sum([list(range(start, end + 1)) for start, end in ranges], [])

    invalids = 0

    for num in numbers:
        s = str(num)
        for reps in range(2, len(s) + 1):
            # PERF: actually would only need to test primes here
            if len(s) % reps == 0 and s[: len(s) // reps] * reps == s:
                print(num)
                invalids += num
                break

    print(invalids, file=sys.stderr)


if __name__ == "__main__":
    main()
