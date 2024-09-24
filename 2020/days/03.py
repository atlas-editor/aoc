INPUT = """..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#"""  # test input

LINES = INPUT.split("\n")
MATRIX = [[c for c in row] for row in LINES]
R = len(MATRIX)
C = len(MATRIX[0])


def p(slopes):
    trees = 1
    for slope in slopes:
        curr = 0
        pos = [0, 0]
        while pos[0] < R:
            if MATRIX[pos[0]][pos[1]] == "#":
                curr += 1
            pos[0] += slope[0]
            pos[1] = (pos[1] + slope[1]) % C
        trees *= curr

    return trees


def p1():
    return p([(1, 3)])


def p2():
    return p([(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)])


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
