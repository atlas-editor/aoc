INPUT = """389125467"""  # test input


class node:
    def __init__(self, val):
        self.val = val
        self.next = None


class ll:
    def __init__(self, ptr, map_, max_val):
        self.ptr = ptr
        self.map = map_
        self.max_val = max_val

    def move(self):
        pickup_vals = {self.ptr.next.val, self.ptr.next.next.val, self.ptr.next.next.next.val}

        for i in range(1, 10):
            dest_val = (self.ptr.val - 1 - i) % self.max_val + 1
            if dest_val not in pickup_vals:
                break

        dest = self.find(dest_val)
        dest_next = dest.next
        first = self.ptr.next
        third = self.ptr.next.next.next

        self.ptr.next = self.ptr.next.next.next.next

        dest.next = first
        third.next = dest_next

        self.ptr = self.ptr.next

    def res(self):
        curr = self.find(1)
        curr = curr.next
        order = ""
        while curr.val != 1:
            order += str(curr.val)
            curr = curr.next
        return order

    def res2(self):
        one = self.find(1)
        return one.next.val * one.next.next.val

    def find(self, val):
        return self.map[val]


def p(max_val, loops):
    cup_lst = [node(int(c)) for c in INPUT]
    cup_lst += [node(i) for i in range(10, max_val + 1)]

    map_ = {}
    for c in cup_lst:
        map_[c.val] = c

    for i in range(len(cup_lst) - 1):
        cup_lst[i].next = cup_lst[i + 1]
    cup_lst[-1].next = cup_lst[0]

    c = ll(cup_lst[0], map_, max_val)
    for i in range(loops):
        c.move()

    return c


def p1():
    return p(9, 100).res()


def p2():
    return p(10 ** 6, 10 ** 7).res2()


if __name__ == '__main__':
    assert p1() == "67384529"
    assert p2() == 149245887792
    print("day 23 ok")
