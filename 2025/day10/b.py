import sys
from functools import cache
from itertools import combinations

from utils import input


def main():
    total = 0
    for lidx, line in enumerate(input.lines(), 1):
        _, *buttons, jolt = line.split()
        jolt = tuple(int(j) for j in jolt[1:-1].split(","))
        buttons = [tuple(int(b) for b in but[1:-1].split(",")) for but in buttons]
        buttons = [tuple(int(i in but) for i in range(len(jolt))) for but in buttons]
        print(f"{buttons=}, {jolt=}")

        # for each jolt_delta reachable by pressing each button at most once, holds the number of presses
        delta_costs = {}

        print(
            "Finding all jolt deltas for press combinations where each button "
            "is pressed at most once"
        )
        for num_presses in range(len(buttons) + 1):
            for buttons_pressed in combinations(buttons, num_presses):
                jolt_delta = (
                    tuple(map(sum, zip(*buttons_pressed)))
                    if num_presses > 0
                    else (0,) * len(jolt)
                )
                print(f"{buttons_pressed=}, {jolt_delta=}")
                if jolt_delta not in delta_costs:
                    delta_costs[jolt_delta] = num_presses

        @cache
        def solve(goal):
            if all(i == 0 for i in goal):
                return 0

            answer = 1e9
            # NOTE: We can see that the goal will have levels which are even and odd
            # An odd level means that the length of the sequence of buttons pressed to reach
            # this level must also be odd, and vice versa for the even levels.
            # We search through all button press sequences that result in a jolt delta
            # with the same parity as the goal, i.e. the same requirement for odd/even button presses per level.
            # After applying the delta, we know that each level must now be even.
            # We therefore know that each button in the solution sequence must be pressed an even number of times.
            # That means we can also "half" the new goal (goal - delta) and multiply the solution by 2.
            # This new goal will again have some odd and some even levels.
            # We recursively solve this until we reach the all 0s goal.
            # NOTE: credit to u/tenthmascot

            for delta, cost in delta_costs.items():
                if tuple(i % 2 for i in delta) != tuple(i % 2 for i in goal):
                    continue
                if all(d <= g for d, g in zip(delta, goal)):
                    next_goal = tuple((g - d) // 2 for d, g in zip(delta, goal))
                    next_answer = 2 * solve(next_goal) + cost
                    if answer is None or next_answer < answer:
                        answer = next_answer
            return answer

        answer = solve(jolt)
        print(f"Machine {lidx}: {answer=}")
        total += answer

    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
