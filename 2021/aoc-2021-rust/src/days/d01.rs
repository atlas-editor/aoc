pub fn p1() -> i32 {
    let input = parse_input();
    _p1(&input)
}

pub fn p2() -> i32 {
    let input = parse_input();
    _p2(&input)
}

fn parse_input() -> Vec<i32> {
    include_str!("../inputs/01.in")
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<_>>()
}

fn _p1(x: &Vec<i32>) -> i32 {
    let mut c = 0;
    let xsize = x.len();
    let mut idx = 1;

    while idx < xsize {
        if x[idx] > x[idx - 1] {
            c += 1;
        }
        idx += 1
    }

    return c;
}

fn _p2(x: &Vec<i32>) -> i32 {
    let mut c = 0;
    let xsize = x.len();
    let mut idx = 0;

    while idx < xsize - 3 {
        if x[idx] < x[idx + 3] {
            c += 1;
        }
        idx += 1
    }

    return c;
}
