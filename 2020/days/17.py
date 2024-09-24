import itertools
from collections import defaultdict
from copy import deepcopy

INPUT = """.#.
..#
###"""  # test input

LINES = INPUT.split("\n")
MATRIX = [[c for c in row] for row in LINES]
R = len(MATRIX)
C = len(MATRIX[0])
X, Y = C, R


def nbrs(*p):
    for q in itertools.product([-1, 0, 1], repeat=len(p)):
        if all(i == 0 for i in q):
            continue
        yield tuple(i + j for i, j in zip(p, q))


def p(dim):
    space = defaultdict(lambda: ".")
    for r in range(R):
        for c in range(C):
            if MATRIX[r][c] == "#":
                if dim == 3:
                    space[(c, r, 0)] = "#"
                else:
                    space[(c, r, 0, 0)] = "#"

    lb = []
    ub = []
    for i in range(dim):
        lb.append(min(p[i] for p in space.keys()))
        ub.append(max(p[i] for p in space.keys()))

    for _ in range(6):
        next_ = deepcopy(space)

        for p in itertools.product(*(range(lb[i] - 1, ub[i] + 2) for i in range(dim))):
            active = 0
            inactive = 0
            for q in nbrs(*p):
                if space[q] == ".":
                    inactive += 1
                else:
                    active += 1

            if space[p] == "#" and not (2 <= active <= 3):
                del next_[p]
            if space[p] == "." and active == 3:
                next_[p] = "#"

                for i in range(dim):
                    lb[i] = min(lb[i], p[i])
                    ub[i] = max(ub[i], p[i])
        space = next_

    return sum(v == "#" for v in space.values())


def p1():
    return p(3)


def p2():
    return p(4)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
