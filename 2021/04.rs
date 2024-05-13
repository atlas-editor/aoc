use std::collections::HashMap;

fn main() {
    let inp = include_str!("/Users/david/x/aoc/2021/inputs/04.in");
    let (nums, matrices) = parse_input(inp);
    println!("{}", p1(&nums, &matrices));
    println!("{}", p2(&nums, &matrices));
}
fn parse_input(s: &str) -> (Vec<i32>, Vec<HashMap<i32, [i32; 2]>>) {
    let ss = s.split("\n\n").collect::<Vec<_>>();
    let nums = ss[0]
        .split(",")
        .map(|n| n.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    let matrices = ss[1..].iter().map(|m| parse_matrix(m)).collect::<Vec<_>>();

    return (nums, matrices);
}

fn parse_matrix(s: &str) -> HashMap<i32, [i32; 2]> {
    let mut matrix = HashMap::new();
    for (row_idx, row) in s.lines().enumerate() {
        for (col_idx, e) in row
            .split_whitespace()
            .map(|n| n.parse::<i32>().unwrap())
            .enumerate()
        {
            matrix.insert(e, [row_idx as i32, col_idx as i32]);
        }
    }

    return matrix;
}

fn p1(nums: &Vec<i32>, matrices: &Vec<HashMap<i32, [i32; 2]>>) -> i32 {
    let mut bingo = HashMap::new();

    let mut winner = 0;
    let mut wi = 0;
    'outer: for (i, n) in nums.iter().enumerate() {
        for (idx, m) in matrices.iter().enumerate() {
            if m.contains_key(&n) {
                let rc = m.get(&n).unwrap();
                let (r, c) = (rc[0], rc[1]);
                *bingo.entry((idx, r, 0)).or_insert(0) += 1;
                *bingo.entry((idx, c, 1)).or_insert(0) += 1;

                if bingo[&(idx, r, 0)] >= 5 || bingo[&(idx, c, 1)] >= 5 {
                    winner = idx;
                    wi = i;
                    break 'outer;
                }
            }
        }
    }

    let s = matrices[winner]
        .keys()
        .filter(|n| !nums[..wi + 1].contains(n))
        .sum::<i32>();

    return s * nums[wi];
}

fn p2(nums: &Vec<i32>, matrices: &Vec<HashMap<i32, [i32; 2]>>) -> i32 {
    let mut bingo = HashMap::new();
    let mut won = HashMap::new();

    let mut li = 0;
    let mut last = 0;
    'outer: for (i, n) in nums.iter().enumerate() {
        for (idx, m) in matrices.iter().enumerate() {
            if m.contains_key(&n) {
                let rc = m.get(&n).unwrap();
                let (r, c) = (rc[0], rc[1]);
                *bingo.entry((idx, r, 0)).or_insert(0) += 1;
                *bingo.entry((idx, c, 1)).or_insert(0) += 1;

                if bingo[&(idx, r, 0)] >= 5 || bingo[&(idx, c, 1)] >= 5 {
                    won.insert(idx, true);
                    li = i;
                    last = idx;
                }

                if won.len() == matrices.len() {
                    break 'outer;
                }
            }
        }
    }

    let s = matrices[last]
        .keys()
        .filter(|n| !nums[..li + 1].contains(n))
        .sum::<i32>();

    return s * nums[li];
}
