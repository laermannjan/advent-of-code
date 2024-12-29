import sys

from utils import input

diskmap = input.lines()[0]


blocks = []
for i, val in enumerate(diskmap):
    if i % 2:
        entry = "."
    else:
        entry = str(i // 2)
    for j in range(int(val)):
        blocks.append(entry)

print(f"{blocks=}")

print("".join(blocks), len("".join(blocks)))

j = len(blocks) - 1
for i in range(len(blocks)):
    if i >= j:
        break
    print(f"{i=}, {j=}")
    while i < j and (val := blocks[i]) == ".":
        if blocks[j] == ".":
            j -= 1
            print(f"{j=}")
        else:
            print(f"swap ({i=},{blocks[i]=}) <> ({j=},{blocks[j]=})")
            blocks[i], blocks[j] = blocks[j], blocks[i]


print("".join(blocks), len("".join(blocks)))

checksum = 0
for i, val in enumerate(blocks):
    if val == ".":
        break
    print(f"{i=} * {int(val)=}")
    checksum += i * int(val)

print(checksum, file=sys.stderr)
