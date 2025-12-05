import re

INPUT = open("input.txt").read().strip()
PARTS = INPUT.split("\n\n")


def pints(string):
    return list(map(int, re.findall(r"\d+", string)))


def p1():
    ranges, ids = PARTS
    franges = []
    for r in ranges.split("\n"):
        a, b = pints(r)
        franges.append((a, b))

    r = 0
    for i in ids.split("\n"):
        ii = int(i)
        for fr in franges:
            if fr[0] <= ii <= fr[1]:
                r += 1
                break

    return r


def p2():
    ranges, _ = PARTS
    franges = []
    pts = set()
    for r in ranges.split("\n"):
        a, b = pints(r)
        pts.add(a)
        pts.add(b + 1)
        franges.append((a, b + 1))

    newr = set()
    pts = sorted(pts)
    for r in franges:
        a, b = r
        preva = a
        for p in pts:
            if a < p < b:
                rr = (preva, p)
                newr.add(rr)
                preva = p
        if preva < b:
            newr.add((preva, b))

    r = 0
    for a, b in newr:
        r += b - a
    return r


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
