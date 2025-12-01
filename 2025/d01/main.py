INPUT = """"""  # test input
INPUT = open("input.txt").read().strip()
LINES = INPUT.split("\n")


def p1():
    s = 50
    r = 0
    for l in LINES:
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
    s = 50
    r = 0
    for l in LINES:
        o = l[0]
        n = int(l[1:])
        if o == "L":
            [a, b] = divmod(s - n, 100)
            r += abs(a)
            if s == 0 and b > 0:
                r -= 1
            elif s > 0 and b == 0:
                r += 1

            s = b
        else:
            [a, b] = divmod(s + n, 100)
            r += a
            s = b

    return r


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
