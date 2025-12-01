import sys

sys.setrecursionlimit(10**6)

INPUT = """"""  # test input
INPUT = open("input.txt").read().strip()


def p1():
    lines = INPUT.split("\n")

    s = 50
    r = 0
    for l in lines:
        o = l[0]
        n = int(l[1:])
        if o == "L":
            s = (s - n) % 100
        else:
            s = (s + n) % 100
        if s == 0:
            r += 1

    return r


def p2():
    lines = INPUT.split("\n")

    s = 50
    r = 0
    for l in lines:
        o = l[0]
        n = int(l[1:])
        if o == "L":
            for i in range(n):
                s = (s - 1) % 100
                if s == 0:
                    r += 1
        else:
            for i in range(n):
                s = (s + 1) % 100
                if s == 0:
                    r += 1

        # if s == 0:
        #     r += 1

    return r


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
