use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap().trim().to_string();

    println!("part1={}", p1(&input));
    println!("part2={}", p2(&input));
}

fn p1(input: &str) -> i64 {
    let mut sm = 0;
    for range in input.split(",") {
        let (a, b) = range.split_once('-').unwrap();

        for i in invalids(a.parse().unwrap(), b.parse().unwrap()) {
            sm += i;
        }
    }

    sm
}

fn p2(input: &str) -> i64 {
    let mut sm = 0;
    for range in input.split(",") {
        let (a, b) = range.split_once('-').unwrap();
        for i in invalids2(a.parse().unwrap(), b.parse().unwrap()) {
            sm += i;
        }
    }

    sm
}

fn invalids2(x: i64, y: i64) -> Vec<i64> {
    let mut q = Vec::new();
    'outer: for i in x..y + 1 {
        let ii = i.to_string();

        'inner: for j in 1..=ii.len() / 2 {
            if ii.len() % j != 0 {
                continue;
            }
            let mut parts = Vec::new();
            for k in 0..ii.len() / j {
                let w = &ii[k * j..(k + 1) * j];
                parts.push(w);
            }
            let a = parts[0];
            for p in parts[1..].iter() {
                if *p != a {
                    continue 'inner;
                }
            }

            q.push(i);
            continue 'outer;
        }
    }

    q
}

fn invalids(x: i64, y: i64) -> Vec<i64> {
    let mut q = Vec::new();
    for i in x..y + 1 {
        let ii = i.to_string();
        if ii.len() % 2 == 1 {
            continue;
        }

        let m = ii.len() / 2;
        let a = &ii[..m];
        let b = &ii[m..];
        if a == b {
            q.push(i);
        }
    }

    q
}
