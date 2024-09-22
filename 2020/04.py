INPUT = """ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in"""  # test input


PARTS = INPUT.split("\n\n")


def p(is_valid):
    valid = 0
    for part in PARTS:
        pairs = part.split()
        passport = {}
        for pair in pairs:
            key, val = pair.split(":")
            passport[key] = val
        if is_valid(passport):
            valid += 1

    return valid


def p1():
    def is_valid(passport):
        return set(passport.keys()) >= {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

    return p(is_valid)


def p2():
    def is_valid(passport):
        byr = passport.get("byr", "")
        if not (byr.isdigit() and 1920 <= int(byr) <= 2002):
            return False

        iyr = passport.get("iyr", "")
        if not (iyr.isdigit() and 2010 <= int(iyr) <= 2020):
            return False

        eyr = passport.get("eyr", "")
        if not (eyr.isdigit() and 2020 <= int(eyr) <= 2030):
            return False

        hgt = passport.get("hgt", "")
        num = hgt[:-2]
        unit = hgt[-2:]
        if not num.isdigit():
            return False
        if unit == "cm":
            if not (150 <= int(num) <= 193):
                return False
        elif unit == "in":
            if not (59 <= int(num) <= 76):
                return False
        else:
            return False

        hcl = passport.get("hcl", "")
        if not (hcl and hcl[0] == "#" and all(("0" <= char <= "9") or ("a" <= char <= "f") for char in hcl[1:])):
            return False

        ecl = passport.get("ecl", "")
        if not (ecl in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}):
            return False

        pid = passport.get("pid", "")
        if not (pid.isdigit() and len(pid) == 9):
            return False

        return True

    return p(is_valid)


if __name__ == '__main__':
    print(f"part1={p1()}")
    print(f"part2={p2()}")
