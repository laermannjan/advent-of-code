import re
import sys
from collections import deque

from utils import input


def main():
    total = 0
    for line in input.lines():
        machine = re.findall(r"\[(.+)\] (.*) \{(.+)\}", line)[0]
        lights, buttons, joltage = machine

        nlights = len(lights)
        lights = int(lights.replace(".", "0").replace("#", "1")[::-1], 2)

        buttons = [
            tuple(map(int, button[1:-1].split(","))) for button in buttons.split()
        ]
        buttons = [sum([1 << toggle for toggle in button]) for button in buttons]
        print(
            f"{lights:0{nlights}b}",
            *[f"{button:0{nlights}b}" for button in buttons],
            joltage,
        )

        dists = {lights: 0}
        paths = deque([lights])

        while paths:
            config = paths.popleft()
            for button in buttons:
                next_config = config ^ button
                print(
                    f"{config:0{nlights}b} ^ {button:0{nlights}b} = {next_config:0{nlights}b}"
                )

                if next_config in dists:
                    print("  ... already found")
                    continue

                next_dist = dists[config] + 1

                if next_config == 0:
                    print(f"need {next_dist} button presses")
                    total += next_dist
                    break

                dists[next_config] = next_dist
                paths.append(next_config)
            else:
                print("-" * 20)
                continue
            break

        print("#" * 20)
    print(total, file=sys.stderr)


if __name__ == "__main__":
    main()
