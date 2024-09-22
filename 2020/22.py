import sys
from collections import deque

sys.setrecursionlimit(10 ** 6)

INPUT = """Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10"""  # test input

PARTS = INPUT.split("\n\n")


def p(game):
    a, b = PARTS
    player_a = deque([int(x) for x in a.split("\n")[1:]])
    player_b = deque([int(x) for x in b.split("\n")[1:]])

    result = game(player_a, player_b)

    score = 0
    for i, e in enumerate(reversed(result), start=1):
        score += i * e

    return score


def p1():
    def game(player_a, player_b):

        while player_a and player_b:
            curr_a = player_a.popleft()
            curr_b = player_b.popleft()

            if curr_a > curr_b:
                player_a.extend([curr_a, curr_b])
            else:
                player_b.extend([curr_b, curr_a])

        return player_a + player_b

    return p(game)


def p2():
    def game(player_a, player_b):
        player_a = tuple(player_a)
        player_b = tuple(player_b)

        SEEN = set()

        def recursive_combat(pa, pb, d):
            if (pa, pb, d) in SEEN:
                return 0, None
            SEEN.add((pa, pb, d))

            if len(pa) == 0:
                return 1, pb
            if len(pb) == 0:
                return 0, pa

            curr_a = pa[0]
            curr_b = pb[0]
            pa = pa[1:]
            pb = pb[1:]

            if len(pa) < curr_a or len(pb) < curr_b:
                if curr_a > curr_b:
                    pa += (curr_a, curr_b)
                else:
                    pb += (curr_b, curr_a)
                return recursive_combat(pa, pb, d)
            else:
                sub_game, _ = recursive_combat(pa[:curr_a], pb[:curr_b], d + 1)
                if sub_game == 0:
                    pa += (curr_a, curr_b)
                else:
                    pb += (curr_b, curr_a)
                return recursive_combat(pa, pb, d)

        _, result = recursive_combat(player_a, player_b, 0)
        return result

    return p(game)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
