use std::collections::{HashMap, HashSet};

use itertools::Itertools;

pub fn p1() -> i32 {
    let (nums, matrices) = parse_input();
    p(&nums, &matrices, false)
}

pub fn p2() -> i32 {
    let (nums, matrices) = parse_input();
    p(&nums, &matrices, true)
}

fn parse_input() -> (Vec<i32>, Vec<HashMap<i32, (i32, i32)>>) {
    let input = include_str!("../inputs/04.in").split("\n\n").collect_vec();
    let nums = input[0]
        .split(',')
        .map(|n| n.parse::<i32>().unwrap())
        .collect_vec();
    let matrices = input[1..].iter().map(|m| parse_matrix(m)).collect_vec();
    (nums, matrices)
}

fn parse_matrix(s: &str) -> HashMap<i32, (i32, i32)> {
    let mut matrix = HashMap::new();
    for (row_idx, row) in s.lines().enumerate() {
        for (col_idx, e) in row
            .split_ascii_whitespace()
            .map(|n| n.parse().unwrap())
            .enumerate()
        {
            matrix.insert(e, (row_idx as i32, col_idx as i32));
        }
    }
    matrix
}

fn p(nums: &[i32], matrices: &[HashMap<i32, (i32, i32)>], p2: bool) -> i32 {
    let mut bingo = HashMap::new();
    let mut won = HashSet::new();

    for (i, n) in nums.iter().enumerate() {
        for (idx, m) in matrices.iter().enumerate() {
            if let Some((r, c)) = m.get(n) {
                *bingo.entry((idx, r, 0)).or_insert(0) += 1;
                *bingo.entry((idx, c, 1)).or_insert(0) += 1;

                if bingo[&(idx, r, 0)] >= 5 || bingo[&(idx, c, 1)] >= 5 {
                    if !p2 {
                        // part1
                        return m.keys().filter(|n| !nums[..i + 1].contains(n)).sum::<i32>() * n;
                    }
                    // for part2 only
                    won.insert(idx);
                    if won.len() == matrices.len() {
                        return m.keys().filter(|n| !nums[..i + 1].contains(n)).sum::<i32>() * n;
                    }
                }
            }
        }
    }
    -1
}
