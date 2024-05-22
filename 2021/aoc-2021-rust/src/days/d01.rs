pub fn p1(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    _p1(&input)
}

pub fn p2(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    _p2(&input)
}

fn parse_input(raw_input: &str) -> Vec<i32> {
    raw_input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect()
}

fn _p1(x: &[i32]) -> i32 {
    let mut c = 0;
    let xsize = x.len();
    let mut idx = 1;

    while idx < xsize {
        if x[idx] > x[idx - 1] {
            c += 1;
        }
        idx += 1
    }

    c
}

fn _p2(x: &[i32]) -> i32 {
    let mut c = 0;
    let xsize = x.len();
    let mut idx = 0;

    while idx < xsize - 3 {
        if x[idx] < x[idx + 3] {
            c += 1;
        }
        idx += 1
    }

    c
}
