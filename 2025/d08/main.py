import math
import re

import networkx as nx

INPUT = open("input.txt").read().strip()
LINES = INPUT.splitlines()


def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))


def g_dists():
    pts = []
    g = nx.Graph()
    for line in LINES:
        p = tuple(ints(line))
        pts.append(p)
        g.add_node(p)

    dists = {}

    for i, v0 in enumerate(pts):
        for j, v1 in enumerate(pts):
            if i <= j:
                continue
            dists[(v0, v1)] = int(
                math.sqrt(
                    abs(v0[0] - v1[0]) ** 2
                    + abs(v0[1] - v1[1]) ** 2
                    + abs(v0[2] - v1[2]) ** 2
                )
            )

    return g, sorted(dists.items(), key=lambda x: x[1])


def p1():
    g, dists = g_dists()

    for k, _ in dists:
        g.add_edge(k[0], k[1])

        if len(g.edges) == 1000:
            break

    return math.prod(
        sorted([len(c) for c in list(nx.connected_components(g))], reverse=True)[:3]
    )


def p2():
    g, dists = g_dists()

    for k, _ in dists:
        g.add_edge(k[0], k[1])

        if nx.is_connected(g):
            return k[0][0] * k[1][0]


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
