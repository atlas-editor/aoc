from collections import defaultdict

INPUT = """light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags."""  # test input

LINES = INPUT.split("\n")


def get_rule_maps():
    contains = defaultdict(list)
    is_contained_in = defaultdict(list)
    for rule in LINES:
        bag, bags_within = rule.split(" contain ")
        if "no other bags" in bags_within:
            continue
        for bag_within in bags_within[:-1].split(", "):
            count, bag_type = bag_within.split(" ", 1)
            if int(count) == 1:
                bag_type += "s"
            contains[bag].append((bag_type, int(count)))
            is_contained_in[bag_type].append(bag)

    return contains, is_contained_in


def p1():
    _, is_contained_in = get_rule_maps()

    stack = is_contained_in["shiny gold bags"]
    count = 0
    visited = set()
    while len(stack) > 0:
        curr_bag = stack.pop()
        count += 1
        visited.add(curr_bag)

        for bag in is_contained_in[curr_bag]:
            if bag not in visited:
                stack.append(bag)

    return count


def p2():
    contains, _ = get_rule_maps()

    stack = contains["shiny gold bags"]
    count = 0
    while len(stack) > 0:
        curr_bag, c = stack.pop()
        count += c

        for bag in contains[curr_bag]:
            stack.append((bag[0], c * bag[1]))

    return count


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
