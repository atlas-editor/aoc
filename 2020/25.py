from itertools import count

INPUT = """5764801
17807724"""  # test input

LINES = INPUT.split("\n")
INTS = [int(i) for i in LINES]


def p1():
    card_pk = INTS[0]
    door_pk = INTS[1]
    door_ls = -1
    val = 1
    for i in count():
        val *= 7
        val %= 20201227
        if val == door_pk:
            door_ls = i + 1
            break

    val = 1
    for i in range(door_ls):
        val *= card_pk
        val %= 20201227

    return val


def p2():
    return


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
