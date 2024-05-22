use std::{
    cmp::{max, min},
    collections::HashMap,
};

use regex::Regex;

pub fn run() {
    let inp = include_str!("inputs/05.in")
        .lines()
        .map(ints)
        .collect::<Vec<_>>();

    println!("{}", p(&inp, false));
    println!("{}", p(&inp, true));
}

fn ints(input: &str) -> Vec<i32> {
    let re = Regex::new(r"-?\d+").unwrap();
    re.find_iter(input)
        .map(|m| m.as_str().parse().unwrap())
        .collect()
}

fn range(x1: i32, y1: i32, x2: i32, y2: i32, diag: bool) -> Vec<[i32; 2]> {
    if x1 == x2 {
        let lb = min(y1, y2);
        let ub = max(y1, y2);
        return (lb..ub + 1).map(|i| [x1, i]).collect();
    } else if y1 == y2 {
        let lb = min(x1, x2);
        let ub = max(x1, x2);
        return (lb..ub + 1).map(|i| [i, y1]).collect();
    } else if diag {
        if x1 < x2 {
            let diff = x2 - x1;
            if y1 < y2 {
                return (0..diff + 1).map(|i| [x1 + i, y1 + i]).collect();
            }
            return (0..diff + 1).map(|i| [x1 + i, y1 - i]).collect();
        }
        let diff = x1 - x2;
        if y1 < y2 {
            return (0..diff + 1).map(|i| [x1 - i, y1 + i]).collect();
        }
        return (0..diff + 1).map(|i| [x1 - i, y1 - i]).collect();
    }
    vec![]
}

fn p(nums: &[Vec<i32>], p2: bool) -> usize {
    let mut m = HashMap::new();
    for p in nums {
        let (x1, y1, x2, y2) = (p[0], p[1], p[2], p[3]);
        for q in range(x1, y1, x2, y2, p2) {
            *m.entry(q).or_insert(0) += 1;
        }
    }
    m.values().filter(|v| **v >= 2).collect::<Vec<_>>().len()
}
