import math

from utils import ints, crt

INPUT = """939
7,13,x,x,59,x,31,19"""  # test input

LINES = INPUT.split("\n")


def p1():
    earliest = int(LINES[0])
    bus_ids = ints(LINES[1])

    min_next_departure = math.inf
    res = 0
    for bid in bus_ids:
        r = earliest % bid
        next_departure = earliest + (bid - r)
        if next_departure < min_next_departure:
            min_next_departure = next_departure
            res = bid * (bid - r)

    return res



def p2():
    nums = []
    rems = []
    for i, bus_id in enumerate(LINES[1].split(",")):
        if bus_id == "x":
            continue
        nums.append(int(bus_id))
        rems.append(-i)

    return crt(nums, rems)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
