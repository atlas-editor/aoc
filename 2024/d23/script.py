import sys

import networkx as nx

INPUT = open(sys.argv[1]).read().strip()
LINES = INPUT.split("\n")


def parse():
    g = nx.Graph()
    for line in LINES:
        u, v = line.split("-")
        g.add_edge(u, v)
    return g


def p1():
    g = parse()

    r = 0
    for clique in nx.enumerate_all_cliques(g):
        match clique:
            case [u, v, w]:
                if u[0] == 't' or v[0] == 't' or w[0] == 't':
                    r += 1
            case x if len(x) > 3:
                break

    return r


def p2():
    g = parse()

    largest = []
    for clique in nx.find_cliques(g):
        if len(clique) > len(largest):
            largest = clique

    return ",".join(sorted(largest))


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
