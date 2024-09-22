import itertools

INPUT = """35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576"""  # test input

LINES = INPUT.split("\n")
INTS = [int(i) for i in INPUT.split("\n")]


def p1():
    window_size = 25 if len(LINES) > 20 else 5  # differentiates between test and real input
    for i in range(window_size, len(INTS) - window_size):
        prev = INTS[i - window_size: i]
        sums = {i + j for i, j in itertools.combinations(prev, 2)}
        num = INTS[i]
        if num not in sums:
            return num


def p2():
    goal = p1()
    for i in range(len(INTS)):
        curr = INTS[i]
        for j in range(i + 1, len(INTS)):
            curr += INTS[j]
            if curr == goal:
                window = INTS[i:j + 1]
                return max(window) + min(window)
            if curr > goal:
                break


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
