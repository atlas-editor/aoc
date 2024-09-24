from copy import deepcopy

from utils import nbrs8

INPUT = """L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL"""  # test input

LINES = INPUT.split("\n")
MATRIX = [[c for c in row] for row in LINES]
R = len(MATRIX)
C = len(MATRIX[0])


def p(occupied_func, limit):
    m = deepcopy(MATRIX)
    i = 0
    while True:
        new_m = deepcopy(m)
        for r in range(R):
            for c in range(C):
                occupied = occupied_func(r, c, m)
                curr = m[r][c]
                if curr == "L" and occupied == 0:
                    new_m[r][c] = "#"
                elif curr == "#" and occupied >= limit:
                    new_m[r][c] = "L"

        if m == new_m:
            occupied = 0
            for row in new_m:
                for c in row:
                    if c == "#":
                        occupied += 1
            return occupied

        m = new_m
        i += 1


def p1():
    def occupied_func(r, c, m):
        occupied = 0
        for rr, cc in nbrs8(r, c ,R, C):
            if m[rr][cc] == "#":
                occupied += 1
        return occupied

    return p(occupied_func, 4)


def p2():
    def occupied_func(r, c, m):
        occupied = 0
        for dr, dc in [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]]:
            j = 1
            while True:
                rr = r + (j * dr)
                cc = c + (j * dc)
                if 0 <= rr < R and 0 <= cc < C:
                    if m[rr][cc] == ".":
                        j += 1
                        continue
                    elif m[rr][cc] == "#":
                        occupied += 1
                break
        return occupied

    return p(occupied_func, 5)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
