from collections import Counter
from copy import deepcopy

from utils import *

INPUT = """sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew"""  # test input

LINES = INPUT.split("\n")

DIRS = {"e": vec((2, 0)), "se": vec((1, -1)), "sw": vec((-1, -1)), "w": vec((-2, 0)), "nw": vec((-1, 1)), "ne": vec((1, 1))}


def layout():
    hexgrid = Counter()
    for tile_desc in LINES:
        tile = vec((0, 0))
        while tile_desc:
            for d in DIRS.keys():
                if tile_desc.startswith(d):
                    tile += DIRS[d]
                    tile_desc = tile_desc.removeprefix(d)
                    break
        hexgrid[tile] = 1 - hexgrid[tile]

    return hexgrid


def p1():
    return layout().total()


def p2():
    hexgrid = layout()

    xs = sorted(x for x, y in hexgrid.keys())
    ys = sorted(y for x, y in hexgrid.keys())
    min_x, max_x = min(xs) - 200, max(xs) + 200
    min_y, max_y = min(ys) - 100, max(ys) + 100

    for i in range(min_x, max_x + 1, 2):
        for j in range(min_y, max_y + 1):
            if j % 2 == 0:
                hexgrid[vec((i, j))] += 0
            else:
                hexgrid[vec((i + 1, j))] += 0

    for i in range(100):
        hexgrid2 = deepcopy(hexgrid)
        for k, v in hexgrid.items():
            black = 0
            for pt in nbrs6(*k):
                if hexgrid[pt] == 1:
                    black += 1

            if v == 0 and black == 2:
                hexgrid2[k] = 1
            if v == 1 and (black == 0 or black > 2):
                hexgrid2[k] = 0

        hexgrid = hexgrid2

    return hexgrid.total()


def nbrs6(x, y):
    for d in DIRS.values():
        yield d + vec((x, y))


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
