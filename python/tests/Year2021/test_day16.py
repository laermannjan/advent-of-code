import pytest
from ..common import validate_problem_test


@pytest.mark.skip
@pytest.mark.parametrize("part", ["a", "b"])
@pytest.mark.parametrize("test", list(range(1, 13)))
def test_day16(part, test):
    assert validate_problem_test("2021", "16", part, str(test))
