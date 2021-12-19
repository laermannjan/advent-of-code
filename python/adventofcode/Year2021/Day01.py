def day01a(numbers: list[str]) -> int:
    this = None
    increases = 0

    for number in numbers:
        last, this = this, int(number)
        if last is None:
            continue

        if last < this:
            increases += 1
    return increases


def day01b(numbers: list[str]) -> int:
    k: int = 3
    this: list[int] = []
    increases = 0

    for number in numbers:
        last, this = this, this + [int(number)]

        if len(last) < k:
            continue

        this.pop(0)
        if sum(last) < sum(this):
            increases += 1
    return increases
