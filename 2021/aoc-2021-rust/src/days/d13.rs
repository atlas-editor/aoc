use std::collections::HashSet;

use itertools::Itertools;
use parse_display::FromStr;

pub fn p1(raw_input: &str) -> usize {
    let (nums, folds) = parse_input(raw_input);
    _p1(&nums, &folds[0])
}

pub fn p2(raw_input: &str) -> String {
    let (nums, folds) = parse_input(raw_input);
    _p2(nums, &folds)
}

fn parse_input(raw_input: &str) -> (Vec<(i32, i32)>, Vec<Fold>) {
    let input = raw_input.split_once("\n\n").unwrap();
    let nums = input
        .0
        .lines()
        .map(|x| {
            x.split_once(',')
                .map(|y| (y.0.parse::<i32>().unwrap(), y.1.parse::<i32>().unwrap()))
                .unwrap()
        })
        .collect_vec();
    let folds = input
        .1
        .lines()
        .map(|x| x.parse::<Fold>().unwrap())
        .collect_vec();
    (nums, folds)
}

#[derive(FromStr, Debug)]
#[display(style = "snake_case")]
enum FoldType {
    X,
    Y,
}

#[derive(FromStr, Debug)]
#[display("fold along {type}={val}")]
struct Fold {
    r#type: FoldType,
    val: i32,
}

impl Fold {
    fn apply(&self, pair: (i32, i32)) -> (i32, i32) {
        match self.r#type {
            FoldType::X => {
                if pair.0 > self.val {
                    (2 * self.val - pair.0, pair.1)
                } else {
                    pair
                }
            }
            FoldType::Y => {
                if pair.1 > self.val {
                    (pair.0, 2 * self.val - pair.1)
                } else {
                    pair
                }
            }
        }
    }
}

fn _p1(nums: &[(i32, i32)], fold: &Fold) -> usize {
    nums.iter()
        .map(|x| fold.apply(*x))
        .collect::<HashSet<_>>()
        .len()
}

fn _p2(mut nums: Vec<(i32, i32)>, folds: &[Fold]) -> String {
    for fold in folds {
        nums = nums.iter().map(|x| fold.apply(*x)).collect()
    }
    let mut code = String::new();
    for r in 0..6 {
        for c in 0..40 {
            if nums.contains(&(c, r)) {
                code.push('#');
            } else {
                code.push(' ');
            }
        }
        code.push('\n');
    }
    code
}
