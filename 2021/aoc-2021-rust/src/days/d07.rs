use std::collections::BTreeMap;

use itertools::Itertools;

use crate::utils::ints;

pub fn p1() -> i32 {
    let input = parse_input();
    _p1(&input)
}

pub fn p2() -> i32 {
    let input = parse_input();
    _p2(&input)
}

fn parse_input() -> Vec<i32> {
    let mut input = include_str!("../inputs/07.in")
        .lines()
        .next()
        .map(ints)
        .unwrap();
    input.sort();
    input
}

fn _p1(nums: &Vec<i32>) -> i32 {
    let mut d: i32 = 0;

    for i in nums[1..].iter() {
        d += i - nums[0];
    }

    let c = nums.iter().counts().into_iter().collect::<BTreeMap<_, _>>();
    let uniqs = c.clone().into_keys().collect::<Vec<_>>();
    let nums_len = nums.len() as i32;
    let mut prev = 0;
    let mut min_ = d;

    for (i, el) in c.into_iter().enumerate() {
        prev += el.1 as i32;
        if i + 1 == uniqs.len() {
            break;
        }
        for _ in uniqs[i] + 1..uniqs[i + 1] + 1 {
            d = d + prev - (nums_len - prev);
            if d < min_ {
                min_ = d;
            }
        }
    }
    min_
}

fn _p2(nums: &Vec<i32>) -> i32 {
    let nums_min = *nums.iter().min().unwrap();
    let nums_max = *nums.iter().max().unwrap();
    let mut res = i32::MAX;

    for i in nums_min..nums_max + 1 {
        let mut d = 0;
        for n in nums {
            let t = (n - i).abs();
            d += (t * (t + 1)) / 2;
        }
        if d < res {
            res = d;
        } else {
            break;
        }
    }

    res
}
