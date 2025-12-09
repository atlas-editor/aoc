import re
from bisect import bisect_left

INPUT = open("input.txt").read().strip()


def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))


def nbrs4(r, c, R, C):
    for dr, dc in [(-1, 0), (0, -1), (0, 1), (1, 0)]:
        new_r = r + dr
        new_c = c + dc
        if 0 <= new_r < R and 0 <= new_c < C:
            yield new_r, new_c


def p1():
    pts = []

    for p in INPUT.splitlines():
        pp = ints(p)
        pts.append((pp[0], pp[1]))

    r = -1
    for a in pts:
        for b in pts:
            if a == b:
                continue
            s = (abs(a[0] - b[0]) + 1) * (abs(a[1] - b[1]) + 1)
            if s > r:
                r = s

    return r


def sgn(x):
    if x < 0:
        return -1
    if x == 0:
        return 0
    return 1


def sizes(pts):
    r = []
    for i, a in enumerate(pts):
        for j, b in enumerate(pts):
            if i >= j:
                continue
            s = (abs(a[0] - b[0]) + 1) * (abs(a[1] - b[1]) + 1)
            r.append((s, i, j))
    return sorted(r, reverse=True)


def p2():
    D = 1000

    pts = []
    xs = set()
    ys = set()
    for p in INPUT.splitlines():
        x, y = ints(p)
        pts.append((x, y))
        xs.add(x)
        ys.add(y)

    xs = sorted(xs)
    ys = sorted(ys)

    def split(p, q):
        px, py = p
        qx, qy = q
        if px == qx:
            y0 = min(py, qy)
            y1 = max(py, qy)
            i = bisect_left(ys, y0)
            j = bisect_left(ys, y1)
            return j - i
        else:
            x0 = min(px, qx)
            x1 = max(px, qx)
            i = bisect_left(xs, x0)
            j = bisect_left(xs, x1)
            return j - i

    arr = [[0] * D for _ in range(D)]

    dir = []
    edges = []
    segments = []
    prev = pts[0]
    for p in pts[1:] + [pts[0]]:
        x, y = p
        px, py = prev
        d = (sgn(x - px), sgn(y - py))
        dir.append(d)
        edges.append([prev, p])
        segments.append(split(prev, p))
        prev = p

    x, y = (500, 500)
    arr[x][y] = 1
    altpts = [(x, y)]
    for d, s in zip(dir, segments):
        for _ in range(s):
            x += d[0]
            y += d[1]
            arr[x][y] = 1
        altpts.append((x, y))

    q = [(0, 0)]
    seen = {q[0]}
    while len(q) > 0:
        x, y = q.pop()
        arr[x][y] = -1

        for xx, yy in nbrs4(x, y, D, D):
            if (xx, yy) not in seen and arr[xx][yy] <= 0:
                q.append((xx, yy))
            seen.add((xx, yy))

    def contained(xmin, xmax, ymin, ymax):
        for x in range(xmin, xmax + 1):
            for y in range(ymin, ymax + 1):
                if arr[x][y] < 0:
                    return False
        return True

    for s, i, j in sizes(pts):
        x0, y0 = altpts[i]
        x1, y1 = altpts[j]
        xmin = min(x0, x1)
        ymin = min(y0, y1)
        xmax = max(x0, x1)
        ymax = max(y0, y1)

        if contained(xmin, xmax, ymin, ymax):
            return s


if __name__ == "__main__":
    print(f"part1={p1()}")
    print(f"part2={p2()}")
