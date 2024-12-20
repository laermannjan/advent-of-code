import sys

from utils import input


def step(tracker, delay, offset):
    tracker = {key - 1: value for key, value in tracker.items()}
    new_offspring = tracker.pop(-1)
    tracker[delay - 1] += new_offspring
    tracker[delay - 1 + offset] = new_offspring
    return tracker


def count_fish_after_days(timings: list[int], days: int):
    delay = 7
    offset = 2

    tracker = {t: 0 for t in range(delay + offset)}
    for t in timings:
        tracker[int(t)] += 1

    for day in range(days):
        tracker = step(tracker, delay, offset)

    return sum(tracker.values())


timings = [int(x) for x in input.lines()[0].split(",")]
print(count_fish_after_days(timings, 80), file=sys.stderr)
