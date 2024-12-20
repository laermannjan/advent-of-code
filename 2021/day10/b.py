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


points = {")": 1, "]": 2, "}": 3, ">": 4}
scores = []

for line in input.lines():
    stack = []
    corrupted = False
    for char in line:
        if char in ["(", "[", "<", "{"]:
            stack.append(char)
        elif are_opposing_brackets(stack[-1], char):
            stack.pop()
        else:
            corrupted = True
            break
    if corrupted:
        continue

    reverse_stack = []
    fixes = []
    for char in reversed(stack):
        if char in [")", "]", "}", ">"]:
            reverse_stack.append(char)
        else:
            if reverse_stack and are_opposing_brackets(char, reverse_stack[-1]):
                reverse_stack.pop()
            else:
                if char == "(":
                    fixes.append(")")
                elif char == "[":
                    fixes.append("]")
                elif char == "{":
                    fixes.append("}")
                elif char == "<":
                    fixes.append(">")

    score = 0
    for fix in fixes:
        score *= 5
        score += points[fix]
    scores.append(score)

scores.sort()
print(scores[len(scores) // 2], file=sys.stderr)
