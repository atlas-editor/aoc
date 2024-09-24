from collections import Counter, defaultdict
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

    for pt, v in list(hexgrid.items()):
        if v == 1:
            for nbr_pt in nbrs6(*pt):
                hexgrid.setdefault(nbr_pt, 0)

    for i in range(100):
        hexgrid2 = deepcopy(hexgrid)
        for k, v in list(hexgrid.items()):
            black = 0
            for pt in nbrs6(*k):
                if hexgrid[pt] == 1:
                    black += 1

            if v == 0 and black == 2:
                for pt in nbrs6(*k):
                    hexgrid2.setdefault(pt, 0)
                hexgrid2[k] = 1
            if v == 1 and (black == 0 or black > 2):
                hexgrid2[k] = 0

        hexgrid = hexgrid2

    return hexgrid.total()


def nbrs6(x, y):
    for d in DIRS.values():
        yield d + vec((x, y))


if __name__ == '__main__':
    assert p1() == 10
    assert p2() == 2208
    print("day 24 ok")
