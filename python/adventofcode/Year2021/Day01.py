def day01a(lines: list[str]) -> int:
    this = None
    increases = 0

    for line in lines:
        last, this = this, int(line)
        if last is None:
            continue

        if last < this:
            increases += 1
    return increases


def day01b(lines: list[str]) -> int:
    this: list[int] = []
    increases = 0
    for line in lines:
        last, this = this, this + [int(line)]

        if len(last) < 3:
            continue

        this.pop(0)
        if sum(last) < sum(this):
            increases += 1
    return increases
