from adventofcode import AVAILABLE_SOLVERS as solvers
import os
import argparse
import time
from memory_profiler import profile

PROJECT_PATH = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("year", choices=("2021",), help="Year")
    parser.add_argument(
        "problem",
        choices=[f"{day}{part}" for part in "ab" for day in range(1, 26)],
        help="Problem",
    )
    parser.add_argument("--profile", action="store_true")
    return parser.parse_args()


def solve(year, problem, input_fname):
    start = time.perf_counter()
    with open(input_fname, "r") as f:
        answer = solvers[year][problem](f.read().splitlines())
    end = time.perf_counter()
    return answer, end - start


if __name__ == "__main__":
    args = parse_args()
    fname = f"Day{args.problem[:-1].zfill(2)}.txt"
    rel_path = f"../input/{args.year}"
    input_fname = f"{PROJECT_PATH}/{rel_path}/{fname}"

    if args.profile:
        solve = profile(solve)
    answer, delta = solve(args.year, args.problem, input_fname)
    print(answer)
    print(f"Took {delta:.4f}s")
