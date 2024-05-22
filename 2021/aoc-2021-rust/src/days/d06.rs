use std::collections::VecDeque;

use crate::utils::ints;

pub fn p1(raw_input: &str) -> u64 {
    let input = parse_input(raw_input);
    _p1(&input)
}

pub fn p2(raw_input: &str) -> u64 {
    let input = parse_input(raw_input);
    _p2(&input)
}

fn parse_input(raw_input: &str) -> Vec<usize> {
    raw_input.lines().next().map(ints).unwrap()
}

fn _p1(nums: &Vec<usize>) -> u64 {
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

fn _p2(nums: &Vec<usize>) -> u64 {
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 5934);
        assert_eq!(p2(raw_input()), 26984457539);
    }

    fn raw_input<'a>() -> &'a str {
        "3,4,3,1,2"
    }
}
