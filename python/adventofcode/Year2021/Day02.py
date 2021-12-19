def day02a(lines: list[str]) -> int:
    horizontal = 0
    depth = 0

    for line in lines:
        inst, val_str = line.split(" ")
        val = int(val_str)

        if inst == "forward":
            horizontal += val
        else:
            depth += val if inst == "down" else -val

    return horizontal * depth


def day02b(lines: list[str]) -> int:
    horizontal = 0
    depth = 0
    aim = 0

    for line in lines:
        inst, val_str = line.split(" ")
        val = int(val_str)

        if inst == "forward":
            horizontal += val
            depth += aim * val
        else:
            aim += val if inst == "down" else -val
    return horizontal * depth
