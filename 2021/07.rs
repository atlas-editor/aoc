use std::collections::BTreeMap;

use itertools::Itertools;
use regex::Regex;

fn main() {
    let mut inp = include_str!("tmp.in").lines().next().map(ints).unwrap();
    inp.sort();
    println!("{}", p1(&inp));
    println!("{}", p2(&inp));
}

fn ints(input: &str) -> Vec<i32> {
    let re = Regex::new(r"-?\d+").unwrap();
    re.find_iter(input)
        .map(|m| m.as_str().parse().unwrap())
        .collect()
}

fn p1(nums: &Vec<i32>) -> i32 {
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

fn p2(nums: &Vec<i32>) -> i32 {
    let nums_min = *nums.iter().min().unwrap();
    let nums_max = *nums.iter().max().unwrap();
    let c = nums.iter().counts();
    let mut min_ = i32::MAX;

    for i in nums_min..nums_max + 1 {
        let mut d = 0;
        for (val, mul) in c.iter() {
            let t = (*val - i).abs();
            d += ((t * (t + 1)) / 2) * (*mul as i32);
        }
        if d < min_ {
            min_ = d;
        } else {
            break;
        }
    }

    min_
}
