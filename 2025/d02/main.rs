use std::{collections::HashSet, fs};

fn main() {
    let input = fs::read_to_string("input.txt").unwrap().trim().to_string();

    println!("part1={}", p1(&input));
    println!("part2={}", p2(&input));
}

fn p1(input: &str) -> i128 {
    let mut sm = 0;
    for range in input.split(",") {
        let (a, b) = range.split_once('-').unwrap();

        let inva = generate_invld(a.parse().unwrap(), b.parse().unwrap(), 2);
        for i in inva {
            sm += i;
        }
    }

    sm
}

fn p2(input: &str) -> i128 {
    let mut sm = 0;
    let mut seen = HashSet::new();
    for range in input.split(",") {
        let (a, b) = range.split_once('-').unwrap();
        for i in 2..b.to_string().len() + 1 {
            let inva = generate_invld(a.parse().unwrap(), b.parse().unwrap(), i as i64);
            for i in inva {
                if seen.contains(&i) {
                    continue;
                }
                seen.insert(i);
                sm += i;
            }
        }
    }

    sm
}

fn generate_invld(a: i64, b: i64, n: i64) -> Vec<i128> {
    let mut invlds = Vec::new();
    let astr = a.to_string();
    let ai = &astr[..astr.len() / n as usize];

    let lb = if ai.is_empty() {
        1
    } else {
        ai.parse::<usize>().unwrap()
    };
    for i in lb.. {
        let mut s = String::new();
        for _ in 0..n {
            s.push_str(&i.to_string());
        }
        let snum = s.parse::<i128>().unwrap();
        if snum < a as i128 {
            continue;
        }
        if snum <= b as i128 {
            invlds.push(snum);
        } else {
            break;
        }
    }

    invlds
}
