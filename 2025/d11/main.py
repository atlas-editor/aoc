from collections import defaultdict

INPUT = open("input.txt").read().strip()


def dfs(start, goal, g):
    DP = {}

    def f(path):
        curr = path[-1]
        if curr in DP:
            return DP[curr]
        if curr == goal:
            return 1

        s = 0
        for n in g[curr]:
            if n not in path:
                c = f(path + [n])
                DP[n] = c
                s += c

        return s

    return f([start])


def load_g():
    g = defaultdict(list)

    for line in INPUT.splitlines():
        u, vs = line.split(":")
        for v in vs.strip().split():
            g[u].append(v)

    return g


def p1():
    return dfs("you", "out", load_g())


def p2():
    g = load_g()
    # there are 0 paths from dac to fft
    return dfs("svr", "fft", g) * dfs("fft", "dac", g) * dfs("dac", "out", g)


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
