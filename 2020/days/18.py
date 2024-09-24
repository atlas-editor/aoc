import math
import re

INPUT = """2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"""  # test input

LINES = INPUT.split("\n")

INNER_RE = re.compile(r"\([^()]+\)")
NUM_RE = re.compile(r"\d+")


def simple(expr):
    expr = expr.replace(" ", "")
    res = int(NUM_RE.search(expr).group())

    i = len(str(res))
    while i < len(expr) - 1:
        op = expr[i]
        other = int(NUM_RE.search(expr[i:]).group())
        if op == "+":
            res += other
        else:
            res *= other
        i += 1 + len(str(other))
    return res


def simple2(expr):
    expr = expr.replace(" ", "")

    tokens = []
    i = 0
    while i < len(expr):
        if (m := NUM_RE.match(expr[i:])) is not None:
            tokens.append(int(m.group()))
            i += len(m.group())
        else:
            tokens.append(expr[i])
            i += 1

    while 1:
        for i in range(len(tokens)):
            if tokens[i] == "+":
                tokens[i - 1:i + 2] = [tokens[i - 1] + tokens[i + 1]]
                break
        else:
            break

    return math.prod(t for t in tokens if isinstance(t, int))


def p(simple_func):
    res = 0

    for line in LINES:
        while (m := INNER_RE.search(line)) is not None:
            line = line[:m.start()] + str(simple_func(line[m.start(): m.end()][1:-1])) + line[m.end():]
        res += simple_func(line)

    return res

def p1():
    return p(simple)

def p2():
    return p(simple2)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
