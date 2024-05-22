use std::collections::VecDeque;

use regex::Regex;

fn main() {
    let inp = include_str!("/Users/david/x/aoc/2021/inputs/04.test")
        .lines()
        .next()
        .map(ints)
        .unwrap();
    println!("{}", p1(&inp));
    println!("{}", p2(&inp));
}

fn ints(input: &str) -> Vec<usize> {
    let re = Regex::new(r"-?\d+").unwrap();
    re.find_iter(input)
        .map(|m| m.as_str().parse().unwrap())
        .collect()
}

fn p1(nums: &Vec<usize>) -> u64 {
    let mut c = nums.to_owned();
    for _ in 0..80 {
        let mut ch = 0;
        for (i, e) in c.clone().iter().enumerate() {
            if *e == 0 {
                c[i] = 6;
                ch += 1;
            } else {
                c[i] -= 1;
            }
        }
        c.append(&mut vec![8; ch]);
    }
    c.len() as u64
}

fn p2(nums: &Vec<usize>) -> u64 {
    let mut state = VecDeque::from([0; 9]);
    for i in nums {
        state[*i] += 1;
    }
    for _ in 0..256 {
        state.rotate_left(1);
        state[6] += state[8];
    }
    state.iter().sum()
}
