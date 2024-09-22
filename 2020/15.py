from utils import ints

INPUT = """0,3,6"""  # test input


def p(rounds):
    nums = ints(INPUT)

    spoken = {n: i+1 for i, n in list(enumerate(nums))[:-1]}
    last = nums[-1]
    r = len(nums)
    while 1:
        match spoken.get(last):
            case None:
                spoken[last] = r
                last = 0
            case i:
                spoken[last] = r
                last = r - i

        if r == rounds-1:
            return last

        r += 1


def p1():
    return p(2020)


def p2():
    return p(30000000)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
