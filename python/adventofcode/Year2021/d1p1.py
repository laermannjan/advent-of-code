import sys

this = None
nums_increased = 0

for line in sys.stdin:
    last, this = this, int(line.rstrip())
    
    if last is None:
        continue

    if last < this:
        nums_increased += 1

print(nums_increased)
    
