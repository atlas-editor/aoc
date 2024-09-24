from collections import Counter

from utils import pints

INPUT = """1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc"""  # test input

LINES = INPUT.split("\n")


def p(is_valid):
    valid = 0
    for line in LINES:
        a, b = pints(line)
        first, word = line.split(": ")
        char = first[-1]
        if is_valid(a, b, char, word):
            valid += 1

    return valid


def p1():
    def is_valid(a, b, char, word):
        c = Counter(word)
        return a <= c[char] <= b

    return p(is_valid)


def p2():
    def is_valid(a, b, char, word):
        inplace = 0
        if word[a - 1] == char:
            inplace += 1
        if word[b - 1] == char:
            inplace += 1
        return inplace == 1

    return p(is_valid)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
