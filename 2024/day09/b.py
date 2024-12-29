import sys

from utils import input

diskmap = input.lines()[0]

files = {}
blanks = []

pos = 0
for i, char in enumerate(diskmap):
    size = int(char)
    if i % 2:
        if size:
            blanks.append((pos, size))
    else:
        assert size, "file length should be greater than 0"
        files[i // 2] = (pos, size)
    pos += size


print(files)
print(blanks)

for fid, (fpos, fsize) in reversed(files.items()):
    for bid, (bpos, bsize) in enumerate(blanks):
        print(f"{fid=}, {bid=}")
        if bpos >= fpos:
            break
        if fsize == bsize:
            files[fid] = (bpos, fsize)
            blanks.pop(bid)
            print(f"{fid=} and {bid=} match, swapping")
            break
        if fsize < bsize:
            files[fid] = (bpos, fsize)
            blanks[bid] = (bpos + fsize, bsize - fsize)
            print(f"{fid=} < {bid=}, moving")
            break

checksum = 0
for fid, (fpos, fsize) in files.items():
    for pos in range(fpos, fpos + fsize):
        checksum += fid * pos

print(checksum, file=sys.stderr)
