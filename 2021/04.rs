use std::collections::HashMap;

fn main() {
    let inp = include_str!("/Users/david/x/aoc/2021/inputs/04.in");
    let pi: (Vec<i32>, Vec<HashMap<i32, [i32; 2]>>) = parse_input(inp);
    // println!("{:#?}", p1(pi));
    println!("{:#?}", p2(pi));
}
fn parse_input(s: &str) -> (Vec<i32>, Vec<HashMap<i32, [i32; 2]>>) {
    let ss = s.split("\n\n").collect::<Vec<_>>();
    let nums = ss[0]
        .split(",")
        .map(|n| n.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    let mut matrices = Vec::new();
    for m in &ss[1..] {
        let matrix = parse_matrix(m);
        matrices.push(matrix);
    }

    return (nums, matrices);
}

fn parse_matrix(s: &str) -> HashMap<i32, [i32; 2]> {
    let mut matrix: HashMap<i32, [i32; 2]> = HashMap::new();
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

fn p1(x: (Vec<i32>, Vec<HashMap<i32, [i32; 2]>>)) -> i32 {
    let (nums, matrices) = x;
    let mut bingo_rows = HashMap::new();
    let mut bingo_cols = HashMap::new();

    let mut winner = 0;
    let mut wi = 0;
    'outer: for (i, n) in nums.clone().into_iter().enumerate() {
        for (idx, m) in matrices.clone().into_iter().enumerate() {
            if m.contains_key(&n) {
                let rc = m.get(&n).unwrap();
                let (r, c) = (rc[0], rc[1]);
                *bingo_rows.entry((idx, r)).or_insert(0) += 1;
                *bingo_cols.entry((idx, c)).or_insert(0) += 1;

                if bingo_rows[&(idx, r)] >= 5 || bingo_cols[&(idx, c)] >= 5 {
                    winner = idx;
                    wi = i;
                    break 'outer;
                }
            }
        }
    }

    let mut s = matrices[winner].keys().sum::<i32>();
    for i in nums[..wi+1].into_iter() {
        if matrices[winner].contains_key(i) {
            s -= i;
        }
    }

    return s * nums[wi];
}

fn p2(x: (Vec<i32>, Vec<HashMap<i32, [i32; 2]>>)) -> i32 {
    let (nums, matrices) = x;
    let mut bingo_rows = HashMap::new();
    let mut bingo_cols = HashMap::new();
    let mut won = HashMap::new();

    let mut wi = 0;
    let mut last = 0;
    'outer: for (i, n) in nums.clone().into_iter().enumerate() {
        for (idx, m) in matrices.clone().into_iter().enumerate() {
            if m.contains_key(&n) {
                let rc = m.get(&n).unwrap();
                let (r, c) = (rc[0], rc[1]);
                *bingo_rows.entry((idx, r)).or_insert(0) += 1;
                *bingo_cols.entry((idx, c)).or_insert(0) += 1;

                if bingo_rows[&(idx, r)] >= 5 || bingo_cols[&(idx, c)] >= 5 {
                    won.insert(idx, true);
                    wi = i;
                    last = idx;
                }

                if won.len() == matrices.len() {
                    break 'outer;
                }
            }
        }
    }

    let mut s = matrices[last].keys().sum::<i32>();
    for i in nums[..wi+1].into_iter() {
        if matrices[last].contains_key(i) {
            s -= i;
        }
    }

    return s * nums[wi];
}