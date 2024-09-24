import math

from utils import ints, pints, flatten

INPUT = """class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9"""  # test input

PARTS = INPUT.split("\n\n")


def parse():
    first, your, nearby = PARTS

    rules = {}
    for line in first.split("\n"):
        a0, a1, b0, b1 = pints(line)
        field, _ = line.split(":")
        rules[field] = (a0, a1, b0, b1)

    return rules, ints(your), [ints(t) for t in nearby.split("\n")[1:]]


def p1():
    rules, _, nearby = parse()

    invalid = 0
    for num in flatten(nearby):
        for a0, a1, b0, b1 in rules.values():
            if a0 <= num <= a1 or b0 <= num <= b1:
                break
        else:
            invalid += num

    return invalid


def p2():
    rules, your, nearby = parse()
    ticket_size = len(your)

    valid = set()
    for ticket in nearby:
        is_valid = True
        for num in ticket:
            for a0, a1, b0, b1 in rules.values():
                if a0 <= num <= a1 or b0 <= num <= b1:
                    break
            else:
                is_valid = False
                break
        if is_valid:
            valid.add(tuple(ticket))

    possible = [set(rules.keys()) for _ in range(ticket_size)]
    for ticket in valid:
        restriction = [set() for _ in range(ticket_size)]
        for i, num in enumerate(ticket):
            for (name, r) in rules.items():
                a0, a1, b0, b1 = r
                if a0 <= num <= a1 or b0 <= num <= b1:
                    restriction[i].add(name)
        possible = [a & b for a, b in zip(possible, restriction)]


    matched = set()
    while len(matched) != ticket_size:
        idx = 0
        for i, p in enumerate(possible):
            if i not in matched and len(p) == 1:
                idx = i
                break

        for i in range(len(possible)):
            if i != idx:
                possible[i] -= possible[idx]

        matched |= {idx}

    return math.prod(num for i, num in enumerate(your) if possible[i].pop().startswith("departure"))


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
