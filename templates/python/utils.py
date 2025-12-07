import itertools
import re

"""
int parsing
"""


def pints(string):
    return list(map(int, re.findall(r"\d+", string)))


def ints(string):
    return list(map(int, re.findall(r"-?\d+", string)))


"""
itertool recipes
"""


def flatten(list_of_lists):
    return list(itertools.chain.from_iterable(list_of_lists))

def first_true(iterable, default=False, predicate=None):
    return next(filter(predicate, iterable), default)



"""
linear algebra
"""


class vec(tuple):
    def __new__(cls, *args):
        if len(args) == 1 and hasattr(args[0], "__iter__"):
            return super().__new__(cls, args[0])
        return super().__new__(cls, args)

    def __add__(self, other):
        return vec(a + b for a, b in zip(self, other))

    def __mul__(self, scalar):
        return vec(a * scalar for a in self)

    def __rmul__(self, scalar):
        return self * scalar

    def rotate2d(self, d):
        # counterclockwise
        a, b = self[0], self[1]
        for _ in range(d % 4):
            a, b = -b, a
        return vec(a, b)


def rotate(matrix, d):
    # counterclockwise
    for _ in range(d % 4):
        matrix = [list(row) for row in zip(*matrix)][::-1]
    return matrix


def fliph(matrix):
    return [row[::-1] for row in matrix]


def flipv(matrix):
    return matrix[::-1]

def transpose(matrix):
    return [list(r) for r in zip(*matrix, strict=True)]


"""
number theory
"""


def crt(nums, rems):
    prod = 1
    for n in nums:
        prod *= n

    result = 0
    for i in range(len(nums)):
        prod_i = prod // nums[i]
        _, inv_i, _ = _gcd_extended(prod_i, nums[i])
        result += rems[i] * prod_i * inv_i

    return result % prod


"""
graphs
"""


def nbrs4(r, c, R, C):
    for dr, dc in [(-1, 0), (0, -1), (0, 1), (1, 0)]:
        new_r = r + dr
        new_c = c + dc
        if 0 <= new_r < R and 0 <= new_c < C:
            yield new_r, new_c


def nbrs8(r, c, R, C):
    for dr, dc in itertools.product([-1, 0, 1], repeat=2):
        if dr == dc == 0:
            continue
        new_r = r + dr
        new_c = c + dc
        if 0 <= new_r < R and 0 <= new_c < C:
            yield new_r, new_c


"""
interval arithmetic
"""


class intervals:
    def __init__(self, ranges):
        # split
        sorted_r = sorted(ranges)
        split = [sorted_r[0]]
        curr_i = sorted_r[0][1]
        for a, b in sorted_r[1:]:
            if curr_i >= a:
                a = curr_i
            if a < b:
                split.append((a, b))
            curr_i = max(curr_i, b)

        # merge
        self.intervals = []
        curr_r = split[0]
        for a, b in split[1:]:
            ca, cb = curr_r
            if cb == a:
                curr_r = (ca, b)
            else:
                self.intervals.append(curr_r)
                curr_r = (a, b)

        if curr_r[0] < curr_r[1]:
            self.intervals.append(curr_r)

    def count(self):
        return len(self.intervals)

    def sum(self):
        return sum([b - a for a, b in self.intervals])


"""
private funcs
"""


def _gcd_extended(a, b):
    if a == 0:
        return b, 0, 1
    gcd, x1, y1 = _gcd_extended(b % a, a)
    x = y1 - (b // a) * x1
    y = x1
    return gcd, x, y
