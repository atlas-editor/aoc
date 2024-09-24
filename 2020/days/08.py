from copy import deepcopy

INPUT = """nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6"""  # test input

LINES = INPUT.split("\n")


def get_instructions():
    instructions = []
    for line in LINES:
        op, arg = line.split(" ")
        instructions.append([op, int(arg)])

    return instructions


def run(instructions):
    acc = 0
    visited = set()
    i = 0
    while True:
        if i in visited:
            return acc, True
        if i >= len(instructions):
            return acc, False
        op, arg = instructions[i]
        visited.add(i)

        match op:
            case "nop":
                i += 1
            case "acc":
                acc += arg
                i += 1
            case "jmp":
                i += arg


def p1():
    instructions = get_instructions()
    return run(instructions)[0]


def p2():
    instructions = get_instructions()

    for i in range(len(instructions)):
        instructions_fixed = deepcopy(instructions)
        if instructions_fixed[i][0] == "nop":
            instructions_fixed[i][0] = "jmp"
        elif instructions_fixed[i][0] == "jmp":
            instructions_fixed[i][0] = "nop"
        else:
            continue

        acc, loop = run(instructions_fixed)
        if not loop:
            return acc


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
