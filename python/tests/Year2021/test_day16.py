import pytest
from ..common import validate_problem_test


@pytest.mark.parametrize("part", ["a", "b"])
@pytest.mark.parametrize("test", range(1, 13))
def test_day16(part, test):
    result, expected = validate_problem_test("2021", "16", part, str(test))
    assert result == expected
