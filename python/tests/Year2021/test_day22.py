import pytest
from ..common import validate_problem_test


@pytest.mark.parametrize("part, test", [("a", 1), ("a", 2), ("b", 2)])
def test_day22(part, test):
    result, expected = validate_problem_test("2021", "22", part, str(test))
    assert result == expected
