use std::collections::VecDeque;

use crate::utils::ints;

pub fn p1(raw_input: &str) -> u64 {
    let input = parse_input(raw_input);
    p(&input, 80)
}

pub fn p2(raw_input: &str) -> u64 {
    let input = parse_input(raw_input);
    p(&input, 256)
}

fn parse_input(raw_input: &str) -> Vec<usize> {
    raw_input.lines().next().map(ints).unwrap()
}

fn p(nums: &Vec<usize>, loops: u16) -> u64 {
    let mut state = VecDeque::from([0; 9]);
    for i in nums {
        state[*i] += 1;
    }
    for _ in 0..loops {
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
