from functools import cache

INPUT = open("input.txt").read().strip()
M = INPUT.splitlines()
R, C = len(M), len(M[0])


def spos():
    for i in range(C):
        if M[0][i] == "S":
            return i


def p1():
    beampos = {}
    beampos[0] = {spos()}

    splt = 0
    for i in range(1, R):
        curr = set()
        for p in beampos[i - 1]:
            if M[i][p] == "^":
                splt += 1
                curr.add(p - 1)
                curr.add(p + 1)
            else:
                curr.add(p)

        beampos[i] = curr

    return splt


def p2():
    @cache
    def f(r, c):
        if r == R:
            return 1

        if M[r][c] == "^":
            return f(r + 1, c - 1) + f(r + 1, c + 1)
        else:
            return f(r + 1, c)

    return f(1, spos())


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part1={p2()}")
