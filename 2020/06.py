INPUT = """abc

a
b
c

ab
ac

a
a
a
a

b"""  # test input

LINES = INPUT.split("\n")
PARTS = INPUT.split("\n\n")

LETTERS = set("abcdefghijklmnopqrstuvwxyz")

def p1():
    counts = 0
    for group in PARTS:
        yes = set(group) & LETTERS
        counts += len(yes)
    return counts


def p2():
    counts = 0
    for group in PARTS:
        letters = LETTERS.copy()
        for person in group.split("\n"):
            letters &= set(person)
        counts += len(letters)
    return counts


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
