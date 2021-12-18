from adventofcode import AVAILABLE_SOLVERS as solvers
import os
import argparse

PROJECT_PATH = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("year", choices=("2021",), help="Year")
    parser.add_argument(
        "problem",
        choices=[f"{day}{part}" for part in "ab" for day in range(1, 26)],
        help="Problem",
    )
    args = parser.parse_args()

    input_fname = f"{args.year}/Day{args.problem[:-1].zfill(2)}.txt"

    with open(f"{PROJECT_PATH}/../input/{input_fname}", "r") as file:
        print(solvers[args.year][args.problem](file.readlines()))
