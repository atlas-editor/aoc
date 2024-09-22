from collections import defaultdict
from copy import deepcopy

from utils import flatten

INPUT = """mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)"""  # test input

LINES = INPUT.split("\n")


def p():
    all_ingredients = set()
    all_allergens = set()
    foods = []

    for desc in LINES:
        a, b = desc.split(" (contains ")
        ingredients = set(a.split())
        allergens = set(b[:-1].split(", "))

        all_ingredients |= ingredients
        all_allergens |= allergens
        foods.append((ingredients, allergens))

    allergen_map = defaultdict(lambda: deepcopy(all_ingredients))
    for allergen in all_allergens:
        for food in foods:
            if allergen in food[1]:
                allergen_map[allergen] &= food[0]

    seen = set()
    while len(seen) != len(allergen_map):
        for v in allergen_map.values():
            if len(v) == 1 and next(iter(v)) not in seen:
                n = next(iter(v))
                seen.add(n)

                for k in allergen_map:
                    if allergen_map[k] != {n}:
                        allergen_map[k] -= {n}

                break

    return foods, allergen_map


def p1():
    foods, allergen_map = p()

    allergic_ingredients = {*flatten(allergen_map.values())}

    res = 0
    for food in foods:
        res += len(food[0] - allergic_ingredients)

    return res


def p2():
    foods, allergen_map = p()

    sorted_allergens = sorted(allergen_map.keys())
    res_allergens = []
    for k in sorted_allergens:
        res_allergens.append(allergen_map[k].pop())

    return ",".join(res_allergens)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
