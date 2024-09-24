import itertools
import math

INPUT = """1721
979
366
299
675
1456"""  # test input


LINES = INPUT.split("\n")

def p(n):
    for i in itertools.combinations(LINES, n):
        ii = list(map(int, i))
        if sum(ii) == 2020:
            return math.prod(ii)

def p1():
    return p(2)

def p2():
    return p(3)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
