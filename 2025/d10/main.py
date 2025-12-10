import re
from functools import cache

from z3 import *

INPUT = open("input.txt").read().strip()


def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))


@cache
def f(state, buttons, i):
    if i > 10:
        return 10**9
    if all(s == 0 for s in state):
        return i

    r = 10**9
    for b in buttons:
        nstate = tuple([1 - s if j in b else s for j, s in enumerate(state)])
        r = min(r, f(nstate, buttons, i + 1))

    return r


def p1():
    s = 0
    for line in INPUT.splitlines():
        data = line.split()
        diagram = tuple([0 if l == "." else 1 for l in data[0][1:-1]])
        buttons = []
        for d in data[1:-1]:
            buttons.append(tuple(ints(d)))

        r = f(diagram, tuple(buttons), 0)
        s += r
    return s


def p2():
    res = 0

    vars = [
        Int("a"),
        Int("b"),
        Int("c"),
        Int("d"),
        Int("e"),
        Int("f"),
        Int("g"),
        Int("h"),
        Int("i"),
        Int("j"),
        Int("k"),
        Int("l"),
        Int("m"),
    ]

    for line in INPUT.splitlines():
        data = line.split()
        joltage = ints(data[-1])

        buttons = []
        for dd in data[1:-1]:
            buttons.append(ints(dd))

        solver = Optimize()

        for idx, target in enumerate(joltage):
            lhs = []
            for idx2, button in enumerate(buttons):
                if idx in button:
                    lhs.append(vars[idx2])

            solver.add(sum(lhs) == target)

        for var in vars:
            solver.add(var >= 0)

        solver.minimize(sum(vars))
        solver.check()
        mm = solver.model()

        res += mm.eval(sum(vars)).as_long()

    return res


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
