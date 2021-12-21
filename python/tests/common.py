from adventofcode import PROJECT_PATH
import adventofcode as aoc


def validate_problem_test(year: str, day: str, part: str, test: str) -> tuple[int, int]:
    day_dir = f"{PROJECT_PATH}/../data/{year}/day{day}"
    input_fname = f"day{day}_input_test_{test}.txt"
    result_fname = f"day{day}{part}_result_test_{test}.txt"

    with (
        open(f"{day_dir}/{input_fname}") as test_input,
        open(f"{day_dir}/{result_fname}") as test_result,
    ):
        input = test_input.read().splitlines()
        result = int(test_result.read())
        return aoc.AVAILABLE_SOLVERS[year][f"{day}{part}"](input), result
