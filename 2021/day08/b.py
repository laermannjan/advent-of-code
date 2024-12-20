import sys

from utils import input

decodes = []
for line in input.lines():
    signals_str, outputs_str = line.split(" | ")
    signals = sorted([set(s) for s in signals_str.split(" ")], key=len)
    outputs = ["".join(sorted(s)) for s in outputs_str.split(" ")]

    digi2sig: dict[int, set[str]] = {
        1: signals[0],  # 2 segments -> shortest
        7: signals[1],  # 3 segments
        4: signals[2],  # 4 segments
        8: signals[-1],  # 7 elements -> longest
    }

    for signal in signals[3:-1]:
        residual = digi2sig[8] - signal
        if len(residual) == 1:
            # 0, 6 or 9
            if len(residual - digi2sig[4]) == 1:
                digi2sig[9] = signal
            elif len(residual - digi2sig[1]) == 1:
                digi2sig[0] = signal
            else:
                digi2sig[6] = signal
        if len(residual) == 2:
            # 2, 3 or 5
            if len(residual - digi2sig[1]) == 2:
                digi2sig[3] = signal
            elif len(residual - digi2sig[4]) == 1:
                digi2sig[5] = signal
            else:
                digi2sig[2] = signal

    sig2digi = {
        # e.g.  {"fg": 1, "efg": 7}
        "".join(sorted(v)): k
        for k, v in digi2sig.items()
    }

    decodes.append(int("".join([str(sig2digi[o]) for o in outputs])))
print(decodes, file=sys.stderr)
