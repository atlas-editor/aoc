#![allow(non_snake_case)]

use std::collections::{HashSet, VecDeque};

pub fn p1() -> i32 {
    let input = parse_input();
    _p1(&input)
}

pub fn p2() -> usize {
    let input = parse_input();
    _p2(&input)
}

fn parse_input() -> Vec<Vec<i32>> {
    include_str!("../inputs/09.in")
        .lines()
        .map(|x| {
            x.chars()
                .map(|y| y.to_string().parse::<i32>().unwrap())
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>()
}

//noinspection ALL
fn _p1(map: &[Vec<i32>]) -> i32 {
    let R = map.len();
    let C = map[0].len();
    let mut res = 0;
    for i in 0..R {
        for j in 0..C {
            let mut nbrs = vec![];
            if i > 0 {
                nbrs.push(map[i - 1][j]);
            }
            if i < R - 1 {
                nbrs.push(map[i + 1][j]);
            }
            if j > 0 {
                nbrs.push(map[i][j - 1]);
            }
            if j < C - 1 {
                nbrs.push(map[i][j + 1]);
            }
            if map[i][j] < *nbrs.iter().min().unwrap() {
                res += map[i][j] + 1;
            }
        }
    }
    res
}

fn _p2(map: &[Vec<i32>]) -> usize {
    let R = map.len();
    let C = map[0].len();
    let mut visited = HashSet::new();
    let mut res = vec![];
    for i in 0..R {
        for j in 0..C {
            if map[i][j] == 9 || visited.contains(&(i, j)) {
                continue;
            }

            let mut csize = 0;
            let mut queue = VecDeque::new();
            queue.push_back((i, j));

            while !queue.is_empty() {
                let (r, c) = queue.pop_front().unwrap();

                if visited.contains(&(r, c)) {
                    continue;
                }
                visited.insert((r, c));
                csize += 1;

                for (dr, dc) in [(0, 1), (0, -1), (1, 0), (-1, 0)] {
                    let rr = r as i32 + dr;
                    let cc = c as i32 + dc;

                    if rr >= 0
                        && rr < R as i32
                        && cc >= 0
                        && cc < C as i32
                        && !visited.contains(&(rr as usize, cc as usize))
                        && map[rr as usize][cc as usize] != 9
                    {
                        queue.push_back((rr as usize, cc as usize));
                    }
                }
            }
            res.push(csize);
        }
    }

    res.sort();
    res.reverse();
    res[..3].iter().product()
}
