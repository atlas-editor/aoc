import math
import re
import sys

import numpy as np

sys.setrecursionlimit(10**6)

# INPUT = """"""  # test input
INPUT = open(sys.argv[1]).read().strip()
PARTS = INPUT.split("\n\n")

def p1():
    ...

def p2():
    res = 0
    for p in PARTS:
        lines = p.split("\n")
        A = ints(lines[0])
        B = ints(lines[1])
        prize = ints(lines[2])

        a0, a1, b0, b1, pr0, pr1 = A[0], A[1], B[0], B[1], prize[0] + 10000000000000, prize[1] + 10000000000000

        a = np.array([[a0, b0], [a1, b1]])
        b = np.array([pr0, pr1])

        s = np.linalg.solve(a, b)
        n, m = s[0], s[1]

        if 0 <= n and math.isclose(n, round(n), abs_tol=0.0001, rel_tol=1e-30) and 0 <= m and math.isclose(m, round(m), abs_tol=0.0001, rel_tol=1e-30):
            res += round(n)*3 + round(m)

    return res

def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))

if __name__ == '__main__':
    print(p1())
    print(p2())