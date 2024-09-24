INPUT = """FBFBBFFRLR"""  # test input

LINES = INPUT.split("\n")


def get_seat_ids():
    seat_ids = []
    for line in LINES:
        row_code = line[:7]
        row = 0
        for ch in row_code:
            if ch == "B":
                row = (row << 1) | 1
            else:
                row <<= 1

        column_code = line[-3:]
        column = 0
        for ch in column_code:
            if ch == "R":
                column = (column << 1) | 1
            else:
                column <<= 1

        seat_ids.append(row * 8 + column)

    return seat_ids


def p1():
    return max(get_seat_ids())


def p2():
    sorted_seat_ids = sorted(get_seat_ids())
    for i in range(len(sorted_seat_ids) - 1):
        if sorted_seat_ids[i] + 1 != sorted_seat_ids[i + 1]:
            return sorted_seat_ids[i] + 1


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
