import itertools
import re

from utils import pints

INPUT = """mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
"""  # test input


def p1():
    onere = re.compile("1")
    zerore = re.compile("0")
    mem = {}
    for part in INPUT.split("mask = ")[1:]:
        program = part.strip().split("\n")
        mask = program[0]
        ones = []
        zeros = []
        for m in onere.finditer(mask):
            ones.append(36 - m.start() - 1)
        for m in zerore.finditer(mask):
            zeros.append(36 - m.start() - 1)

        for locnum in program[1:]:
            loc, num = pints(locnum)
            for o in ones:
                num |= 1 << o
            for z in zeros:
                num &= ~(1 << z)
            mem[loc] = num
    return sum(mem.values())


def p2():
    onere = re.compile("1")
    xre = re.compile("X")
    mem = {}
    for part in INPUT.split("mask = ")[1:]:
        program = part.strip().split("\n")
        mask = program[0]
        ones = []
        xs = []
        for m in onere.finditer(mask):
            ones.append(36 - m.start() - 1)
        for m in xre.finditer(mask):
            xs.append(36 - m.start() - 1)

        for locnum in program[1:]:
            loc, num = pints(locnum)
            for o in ones:
                loc |= 1 << o

            for seq in itertools.product("01", repeat=len(xs)):
                curr = loc
                for b, idx in zip(seq, xs):
                    if b == "1":
                        curr |= 1 << idx
                    else:
                        curr &= ~(1 << idx)
                mem[curr] = num

    return sum(mem.values())


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
