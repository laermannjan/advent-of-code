import sys

this = []
nums_increased = 0

for i, line in enumerate(sys.stdin):
    last, this = this, this + [int(line.rstrip())]
    
    if len(last) < 3:
        continue
    
    this.pop(0)
    if sum(last) < sum(this):
        nums_increased += 1

print(nums_increased)
    
