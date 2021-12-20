import pytest
from ..common import validate_problem_test


@pytest.mark.skip
@pytest.mark.parametrize("part", ["a", "b"])
@pytest.mark.parametrize("test", [1])
def test_day18(part, test):
    assert validate_problem_test("2021", "18", part, str(test))
