from functools import cache

INPUT = open("input.txt").read().strip()


def p1():
    m = INPUT.splitlines()
    R, C = len(m), len(m[0])
    S = -1
    for i in range(C):
        if m[0][i] == "S":
            S = i
            break

    beampos = {}
    beampos[0] = {S}

    splt = 0
    for i in range(1, R):
        curr = set()
        for p in beampos[i - 1]:
            if m[i][p] == "^":
                splt += 1
                curr.add(p - 1)
                curr.add(p + 1)
            else:
                curr.add(p)

        beampos[i] = curr

    return splt


def p2():
    m = INPUT.splitlines()
    R, C = len(m), len(m[0])
    S = -1
    for i in range(C):
        if m[0][i] == "S":
            S = i
            break

    @cache
    def f(r, c):
        if r == R:
            return 1

        if m[r][c] == "^":
            return f(r + 1, c - 1) + f(r + 1, c + 1)
        else:
            return f(r + 1, c)

    return f(1, S)


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part1={p2()}")
