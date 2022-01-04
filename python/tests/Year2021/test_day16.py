import pytest
from ..common import validate_problem_test


@pytest.mark.parametrize("part, test", [
    ("a", 1),
    ("a", 2),
    ("a", 3),
    ("a", 4),
    ("b", 5),
    ("b", 6),
    ("b", 7),
    ("b", 8),
    ("b", 9),
    ("b", 10),
    ("b", 11),
    ("b", 12),
])
def test_day16(part, test):
    result, expected = validate_problem_test("2021", "16", part, str(test))
    assert result == expected
