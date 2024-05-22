use itertools::Itertools;
use std::collections::HashMap;

use crate::utils::ints;

pub fn p1(raw_input: &str) -> usize {
    let input = parse_input(raw_input);
    p(&input, false)
}

pub fn p2(raw_input: &str) -> usize {
    let input = parse_input(raw_input);
    p(&input, true)
}

fn parse_input(raw_input: &str) -> Vec<Vec<i32>> {
    raw_input.lines().map(ints).collect()
}

fn range(x1: i32, y1: i32, x2: i32, y2: i32, diag: bool) -> Vec<[i32; 2]> {
    if x1 == x2 {
        let (&lb, &ub) = [y1, y2].iter().minmax().into_option().unwrap();
        return (lb..ub + 1).map(|i| [x1, i]).collect();
    } else if y1 == y2 {
        let (&lb, &ub) = [x1, x2].iter().minmax().into_option().unwrap();
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
    m.values().filter(|v| **v >= 2).collect_vec().len()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 5);
        assert_eq!(p2(raw_input()), 12);
    }

    fn raw_input<'a>() -> &'a str {
        "0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"
    }
}
