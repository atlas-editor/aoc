from utils import vec

INPUT = """F10
N3
F7
R90
F11"""  # test input

LINES = INPUT.split("\n")


def p1():
    return p(vec((1, 0)), False)


def p2():
    return p(vec((10, 1)), True)


def p(dir_, waypoint):
    directions = {"N": vec((0, 1)), "S": vec((0, -1)), "E": vec((1, 0)), "W": vec((-1, 0))}
    pos = vec((0, 0))
    for instruction in LINES:
        arg = int(instruction[1:])
        match instruction[0]:
            case "F":
                pos += dir_ * arg
            case "L":
                dir_ = dir_.rotate2d(arg)
            case "R":
                dir_ = dir_.rotate2d(-arg)
            case d:
                if waypoint:
                    dir_ += directions[d] * arg
                else:
                    pos += directions[d] * arg
    return abs(pos[0]) + abs(pos[1])


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
