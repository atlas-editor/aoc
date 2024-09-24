from collections import Counter
from functools import cache

INPUT = """28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3"""  # test input

LINES = INPUT.split("\n")
INTS = [int(i) for i in INPUT.split("\n")]


def p1():
    jolts = [0] + sorted(INTS) + [max(INTS) + 3]
    diffs = Counter()
    for i in range(len(jolts) - 1):
        a, b = jolts[i], jolts[i + 1]
        diffs[b - a] += 1
    return diffs[1] * diffs[3]


def p2():
    jolts = [0] + sorted(INTS) + [max(INTS) + 3]

    @cache
    def dp(i):
        if i == len(jolts) - 1:
            return 1
        curr = jolts[i]
        res = 0
        for j in range(i + 1, len(jolts)):
            if 1 <= jolts[j] - curr <= 3:
                res += dp(j)
        return res

    return dp(0)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
