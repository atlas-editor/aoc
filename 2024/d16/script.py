import math
import pathlib
import sys
from collections import defaultdict
from heapq import heappop, heappush

sys.setrecursionlimit(10 ** 6)

INPUT = """#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################"""  # test input
INPUT = open(f"input.txt").read().strip()

LINES = INPUT.split("\n")
# INTS = [int(i) for i in INPUT.split("\n")]
PARTS = INPUT.split("\n\n")
MATRIX = [[c for c in row] for row in LINES]
R = len(MATRIX)
C = len(MATRIX[0])


class vec(tuple):
    def __add__(self, other):
        return vec(a + b for a, b in zip(self, other))


def turn(v):
    a, b = v[0], v[1]
    return [vec((-b, a)), vec((b, -a))]


def p1():
    start = (-1, -1)
    end = (-1, -1)
    for r in range(R):
        for c in range(C):
            if MATRIX[r][c] == 'S':
                start = vec((r, c))
            if MATRIX[r][c] == 'E':
                end = vec((r, c))

    start_state = (start, vec((0, 1)))
    dist = defaultdict(lambda: math.inf, {start_state: 0})
    prev = {start_state: None}
    seen = set()
    h = [(0, start_state)]
    best = math.inf
    best_spots = set()
    while len(h) > 0:


        d, pd = heappop(h)
        if ((d, pd)) in seen:
            continue
        seen.add((d, pd))

        pos, dir = pd
        if pos == end:
            best = d
            print("best found")

            curr = frozenset({pd})
            def f(aa):
                if aa is None:
                    return

                for i in aa:
                    best_spots.add(i[0])
                    f(prev.get(i))

            f(frozenset({pd}))

            pathpts = []
            while curr is not None:

                for i in curr:
                    best_spots.add(i[0])
                    pass
                p, dd = set(curr).pop()
                pathpts.append(p)
                best_spots.add(p)
                curr = prev.get((p, dd))

            for r in range(R):
                for c in range(C):
                    if MATRIX[r][c] != '.':
                        print(MATRIX[r][c], end='')
                    else:
                        if vec((r,c)) in pathpts:
                            print('!', end='')
                        else:
                            print('.', end='')
                print()
            continue
        if d > best:
            continue

        forward = pos + dir
        r, c = forward[0], forward[1]
        if 0 <= r < R and 0 <= c < C and MATRIX[r][c] != '#':
            nbr = (forward, dir)
            alt = d + 1
            if alt < dist[nbr]:
                dist[nbr] = alt
                prev[nbr] = frozenset({pd})
                # prev[nbr] = frozenset(set(prev.get(nbr, set())) | {pd})
                heappush(h, (alt, nbr))
                seen.add(nbr)

            if alt == dist[nbr]:
                dist[nbr] = alt
                prev[nbr] = frozenset(set(prev[nbr]) | {pd})
                heappush(h, (alt, nbr))
                seen.add(nbr)

        for t in turn(dir):
            nbr = (pos, t)

            alt = d + 1000
            if alt < dist[nbr]:
                dist[nbr] = alt
                prev[nbr] = frozenset({pd})
                # prev[nbr] = frozenset(set(prev.get(nbr, set())) | {pd})
                heappush(h, (alt, nbr))
                seen.add(nbr)
            if alt == dist[nbr]:
                dist[nbr] = alt
                prev[nbr] = frozenset(set(prev.get(nbr, set())) | {pd})
                heappush(h, (alt, nbr))
                seen.add(nbr)


    # pts = set()
    # def dfs(u, dir, seen, path, distance):
    #     if u == end:
    #         for p in path:
    #             pts.add(p)
    #         return
    #
    #     seen[u] = True
    #     path.append(u)
    #
    #     if


    return len(best_spots)

def p2():
    return


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
