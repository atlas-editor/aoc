import re
from math import prod

INPUT = open("input.txt").read().strip()


def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))


def p1():
    parts = INPUT.split("\n\n")
    bdata = parts[:-1]
    cases = parts[-1]
    bsize = []
    for b in bdata:
        bsize.append(b.count("#"))

    r = 0
    for case in cases.splitlines():
        dim, counts = case.split(": ")
        c = ints(counts)
        req = 0
        for i, s in enumerate(bsize):
            req += c[i] * s
        r += int(req <= prod(ints(dim)))

    return r


def p2():
    return


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
