use std::collections::HashMap;

pub fn p1(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    _p1(&input)
}

pub fn p2(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    _p2(&input)
}

fn parse_input(raw_input: &str) -> Vec<String> {
    raw_input.lines().map(|l| l.to_string()).collect::<Vec<_>>()
}

fn _p1(x: &Vec<String>) -> i32 {
    let xsize = x.len();
    let numsize = x[0].len();
    let mut counter = HashMap::new();

    for e in x {
        for (idx, ch) in e.chars().enumerate() {
            if ch == '1' {
                *counter.entry(idx).or_insert(0) += 1;
            }
        }
    }

    let mut gamma = String::new();
    let mut epsilon = String::new();
    for i in 0..numsize {
        if counter[&i] > xsize / 2 {
            gamma.push('1');
            epsilon.push('0');
        } else {
            gamma.push('0');
            epsilon.push('1');
        }
    }

    i32::from_str_radix(&gamma, 2).unwrap() * i32::from_str_radix(&epsilon, 2).unwrap()
}

fn _p2(x: &[String]) -> i32 {
    let numsize = x[0].len();

    let mut o2 = x.iter().collect::<Vec<_>>();
    let mut co2 = x.iter().collect::<Vec<_>>();

    for idx in 0..numsize {
        o2 = match ones_zeros(o2.clone(), idx) {
            (ones, zeros) if ones.len() >= zeros.len() => ones,
            (_, zeros) => zeros,
        };

        co2 = match ones_zeros(co2.clone(), idx) {
            (ones, zeros) if !ones.is_empty() && (ones.len() < zeros.len() || zeros.is_empty()) => {
                ones
            }
            (_, zeros) => zeros,
        };
    }

    i32::from_str_radix(o2[0], 2).unwrap() * i32::from_str_radix(co2[0], 2).unwrap()
}

fn ones_zeros(x: Vec<&String>, idx: usize) -> (Vec<&String>, Vec<&String>) {
    return x
        .into_iter()
        .partition(|&s| s.chars().nth(idx).unwrap() == '1');
}
