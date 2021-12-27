from adventofcode import AVAILABLE_SOLVERS as solvers
import argparse
import time
from adventofcode import PROJECT_PATH
from memory_profiler import profile
import os


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("year", choices=("2021",), help="Year")
    parser.add_argument(
        "problem",
        choices=[f"{day}{part}" for part in "ab" for day in range(1, 26)],
        help="Problem",
    )
    parser.add_argument("--profile", action="store_true")
    parser.add_argument("--viz", action="store_true", help="Visualize solution if possible")
    return parser.parse_args()


def solve(year, day, part, input_fname):
    start = time.perf_counter()
    # with open(input_fname, "r") as f:
    #     answer = solvers[year][f"{day}{part}"](f.read().splitlines())
    answer = solvers[year][f"{day}{part}"](input_fname)
    end = time.perf_counter()
    return answer, end - start


if __name__ == "__main__":
    args = parse_args()

    os.environ["ADVENTOFCODE_VISUALIZE"] = '1' if args.viz else '0'
    day = args.problem[:-1].zfill(2)
    part = args.problem[-1]

    fname = f"day{day}/day{day}_input_puzzle.txt"
    rel_path = f"../data/{args.year}"
    input_fname = f"{PROJECT_PATH}/{rel_path}/{fname}"

    if args.profile:
        solve = profile(solve)
    answer, delta = solve(args.year, day, part, input_fname)
    print(answer)
    print(f"Took {delta:.4f}s")
