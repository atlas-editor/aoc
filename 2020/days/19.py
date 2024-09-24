from functools import cache

from utils import ints

INPUT = """0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb"""  # test input

PARTS = INPUT.split("\n\n")


def p(extra):
    first, second = PARTS

    rules = {}
    for line in first.split("\n"):
        num, r = line.split(": ")
        num = int(num)
        if "a" in r:
            rules[num] = "a"
        elif "b" in r:
            rules[num] = "b"
        elif "|" in r:
            r0, r1 = r.split(" | ")
            rules[num] = [tuple(ints(r0)), tuple(ints(r1))]
        else:
            rules[num] = [tuple(ints(r))]

    if extra:
        rules[8] = [(42,), (42, 8)]
        rules[11] = [(42, 31), (42, 11, 31)]

    @cache
    def dp(message, remaining):
        if len(message) == 0 and len(remaining) == 0:
            return True
        if {0} < {len(message), len(remaining)}:
            return False

        rule = rules[remaining[0]]
        if isinstance(rule, str):
            if message[0] == rule:
                return dp(message[1:], remaining[1:])
            else:
                return False

        for subrule in rule:
            if dp(message, subrule + remaining[1:]):
                return True

        return False

    return sum(dp(msg, (0,)) for msg in second.split("\n"))


def p1():
    return p(False)


def p2():
    return p(True)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
