import math
import pathlib
from collections import Counter, defaultdict
from functools import cache

from utils import ints, flatten, rotate, fliph, flipv, vec

INPUT = """Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###..."""  # test input
INPUT = open(f"inputs/{pathlib.Path(__file__).stem}.in").read().strip()

PARTS = INPUT.split("\n\n")

DIR_MAP = {0: (0, 1), 1: (-1, 0), 2: (0, -1), 3: (1, 0)}
EDGE_CORRESPONDENCE = {0: 2, 1: 3, 2: 0, 3: 1}


def p1():
    tiletoedges = {}
    for desc in PARTS:
        tile_id = ints(desc.split("\n")[0]).pop()
        tile = [[c for c in row] for row in desc.split("\n")[1:]]
        tiletoedges[tile_id] = set(flatten(map(lambda e: [tuple(e), tuple(e)[::-1]], edges(tile))))

    degree = Counter()
    for id0 in tiletoedges.keys():
        for id1 in tiletoedges.keys():
            if id0 <= id1:
                continue
            if tiletoedges[id0] & tiletoedges[id1]:
                degree[id0] += 1
                degree[id1] += 1

    return math.prod([k for k, v in degree.items() if v == 2])


@cache
def monster_coords():
    monster = """                  # 
#    ##    ##    ###
 #  #  #  #  #  #   """
    m = [[c for c in row] for row in monster.split("\n")]
    coords = []
    for r in range(len(m)):
        for c in range(len(m[0])):
            if m[r][c] == "#":
                coords.append((r, c))

    return coords


def edges(matrix):
    return [[row[-1] for row in matrix], matrix[0], [row[0] for row in matrix], matrix[-1]]


def all_symmetries(matrix):
    for d in range(4):
        yield rotate(matrix, d * 90)
        yield fliph(rotate(matrix, d * 90))


def edge_match(old, new):
    e = edges(old)
    f = edges(new)
    for i in range(4):
        for j in range(4):
            if e[i] in [f[j], f[j][::-1]]:
                return i, j


def match(old, new):
    i, j = edge_match(old, new)

    d = (EDGE_CORRESPONDENCE[i] - j) % 4
    new = rotate(new, d * 90)

    if edges(new)[EDGE_CORRESPONDENCE[i]] != edges(old)[i]:
        if i % 2 == 0:
            new = flipv(new)
        else:
            new = fliph(new)

    return new, vec(DIR_MAP[i])


def trim_edges(matrix):
    return [row[1:-1] for row in matrix[1:-1]]


def p2():
    tiletohash = {}
    tilesmap = {}
    for desc in PARTS:
        tile_id = ints(desc.split("\n")[0]).pop()
        tile = [[c for c in row] for row in desc.split("\n")[1:]]
        tilesmap[tile_id] = tile

        e0 = tile[0]
        e1 = tile[-1]
        e2 = [row[0] for row in tile]
        e3 = [row[-1] for row in tile]

        tiletohash[tile_id] = set(flatten(map(lambda e: [tuple(e), tuple(e)[::-1]], [e0, e1, e2, e3])))
    g = defaultdict(list)
    for id0 in tiletohash.keys():
        for id1 in tiletohash.keys():
            if id0 <= id1:
                continue
            if tiletohash[id0] & tiletohash[id1]:
                g[id0].append(id1)
                g[id1].append(id0)

    corner = [k for k, v in g.items() if len(v) == 2].pop()

    stack = [(corner, vec((11, 0)))]
    mapa = [[None for _ in range(int(math.sqrt(len(tilesmap))))] for _ in range(int(math.sqrt(len(tilesmap))))]
    mapa[11][0] = trim_edges(tilesmap[corner])
    seen = {corner}
    while stack:
        curr, pt = stack.pop()
        for nbr in g[curr]:
            if nbr in seen:
                continue
            nbrrotated, mv = match(tilesmap[curr], tilesmap[nbr])
            tilesmap[nbr] = nbrrotated
            newpt = pt + mv
            mapa[newpt[0]][newpt[1]] = trim_edges(nbrrotated)
            stack.append((nbr, newpt))
            seen.add(nbr)

    bigmapa = []
    for row in mapa:
        for i in range(8):
            bigmapa.append(sum([matrix[i] for matrix in row], start=[]))

    for d in range(4):
        bigmapa = rotate(bigmapa, d * 90)
        R = len(bigmapa)
        C = len(bigmapa[0])
        monsterpts = set()
        monsters = 0
        for r in range(R):
            for c in range(C):
                coordstoadd = set()
                ismonster = True
                for dr, dc in monster_coords():
                    if r + dr < R and c + dc < C and bigmapa[r + dr][c + dc] == "#":
                        coordstoadd.add((r + dr, c + dc))
                    else:
                        ismonster = False
                        break
                if not ismonster:
                    continue
                monsterpts |= coordstoadd
                monsters += 1
        pass

    bigmapa = fliph(bigmapa)
    for d in range(4):
        bigmapa = rotate(bigmapa, d * 90)
        R = len(bigmapa)
        C = len(bigmapa[0])
        monsterpts = set()
        monsters = 0
        for r in range(R):
            for c in range(C):
                coordstoadd = set()
                ismonster = True
                for dr, dc in monster_coords():
                    if r + dr < R and c + dc < C and bigmapa[r + dr][c + dc] == "#":
                        coordstoadd.add((r + dr, c + dc))
                    else:
                        ismonster = False
                        break
                if not ismonster:
                    continue
                monsterpts |= coordstoadd
                monsters += 1
        pass

    res = 0
    for r in range(R):
        for c in range(C):
            if bigmapa[r][c] == "#" and (r, c) not in monsterpts:
                res += 1

    return res


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
