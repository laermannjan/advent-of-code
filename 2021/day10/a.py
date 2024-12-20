import sys

from utils import input


def are_opposing_brackets(a, b):
    if (
        (a == "(" and b == ")")
        or (a == "[" and b == "]")
        or (a == "{" and b == "}")
        or (a == "<" and b == ">")
    ):
        return True
    return False


points = {")": 3, "]": 57, "}": 1197, ">": 25137}
corruptions = {")": 0, "]": 0, "}": 0, ">": 0}

for line in input.lines():
    stack = []
    for char in line:
        if char in ["(", "[", "<", "{"]:
            stack.append(char)
        elif are_opposing_brackets(stack[-1], char):
            stack.pop()
        else:
            corruptions[char] += 1
            break
print(sum([corruptions[x] * points[x] for x in points]), file=sys.stderr)
