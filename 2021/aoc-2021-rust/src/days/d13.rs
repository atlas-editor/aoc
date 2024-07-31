use bstr::ByteSlice;
use itertools::Itertools;

use crate::days::utils::*;
use crate::matrix;

pub fn p1(raw_input: &[u8]) -> usize {
    let (nums, folds) = parse_input(raw_input);
    _p1(&nums, &folds)
}

pub fn p2(raw_input: &[u8]) -> String {
    let (nums, folds) = parse_input(raw_input);
    _p2(nums, &folds)
}

fn parse_input(raw_input: &[u8]) -> (Vec<(u16, u16)>, Vec<Fold>) {
    let input = raw_input.split_once_str(b"\n\n").unwrap();
    let dots = input
        .0
        .lines()
        .map(|x| {
            x.split_once_str(b",")
                .map(|y| (atopi(y.0), atopi(y.1)))
                .unwrap()
        })
        .collect_vec();
    let folds = input.1.lines().map(|x| Fold::parse(x)).collect_vec();
    (dots, folds)
}

#[derive(PartialEq)]
enum FoldType {
    X,
    Y,
}

struct Fold {
    r#type: FoldType,
    val: u16,
}

impl Fold {
    fn apply(&self, dot: (u16, u16)) -> (u16, u16) {
        match self.r#type {
            FoldType::X => {
                if dot.0 > self.val {
                    (2 * self.val - dot.0, dot.1)
                } else {
                    dot
                }
            }
            FoldType::Y => {
                if dot.1 > self.val {
                    (dot.0, 2 * self.val - dot.1)
                } else {
                    dot
                }
            }
        }
    }

    fn parse(repr: &[u8]) -> Self {
        let val = atopi(&repr[13..]);
        if repr[11] == b'x' {
            Self {
                r#type: FoldType::X,
                val,
            }
        } else {
            Self {
                r#type: FoldType::Y,
                val,
            }
        }
    }
}

fn find_first_fold_val(folds: &[Fold], r#type: FoldType) -> usize {
    folds.iter().find(|fold| fold.r#type == r#type).unwrap().val as usize
}

fn find_last_fold_val(folds: &[Fold], r#type: FoldType) -> usize {
    folds
        .iter()
        .rev()
        .find(|fold| fold.r#type == r#type)
        .unwrap()
        .val as usize
}

fn _p1(nums: &[(u16, u16)], folds: &[Fold]) -> usize {
    let (x_size, y_size): (usize, usize) = (
        2 * find_first_fold_val(folds, FoldType::X) + 1,
        2 * find_first_fold_val(folds, FoldType::Y) + 1,
    );
    let mut paper = matrix![false; y_size, x_size];

    let mut dots = 0;
    for pt in nums {
        let (x, y) = folds[0].apply(*pt);
        let (x, y) = (x as usize, y as usize);
        if !paper[(y, x)] {
            dots += 1;
            paper[(y, x)] = true;
        }
    }
    dots
}

fn _p2(mut nums: Vec<(u16, u16)>, folds: &[Fold]) -> String {
    let (x_size, y_size): (usize, usize) = (
        find_last_fold_val(folds, FoldType::X),
        find_last_fold_val(folds, FoldType::Y),
    );
    let mut paper = matrix!["."; y_size, x_size];

    for mut pt in nums {
        for fold in folds {
            pt = fold.apply(pt);
        }
        paper[(pt.1 as usize, pt.0 as usize)] = "#";
    }
    paper.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 17);
        assert_eq!(
            p2(raw_input()),
            "#####
#...#
#...#
#...#
#####
.....
....."
        );
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5"
    }
}
