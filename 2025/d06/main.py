import re
from math import prod

INPUT = open("input.txt").read().strip()


def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))


def rotate(matrix, d):
    for _ in range(d % 4):
        matrix = [list(row) for row in zip(*matrix)][::-1]
    return matrix


def p1():
    lines = INPUT.split("\n")
    nums = []
    for i in range(len(lines) - 1):
        nums.append(ints(lines[i]))
    ops = lines[len(lines) - 1].split()

    nums = rotate(nums, -1)

    r = 0
    for i, op in enumerate(ops):
        if op == "*":
            r += prod(nums[i])
        else:
            r += sum(nums[i])

    return r


def p2():
    lines = INPUT.split("\n")

    nums = lines[: len(lines) - 1]
    ops = lines[len(lines) - 1].split()

    nums = rotate(nums, -1)

    curr = []
    j = 0
    r = 0
    for nl in nums:
        if "".join(nl).strip() == "":
            if ops[j] == "*":
                r += prod(curr)
            else:
                r += sum(curr)
            curr = []
            j += 1
        else:
            curr.append(int("".join(nl)[::-1]))

    if ops[len(ops) - 1] == "*":
        r += prod(curr)
    else:
        r += sum(curr)

    return r


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
