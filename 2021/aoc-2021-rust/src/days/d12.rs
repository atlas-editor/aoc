use std::collections::{HashMap, HashSet};

use itertools::Itertools;

pub fn p1(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    p(&input, false)
}

pub fn p2(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    p(&input, true)
}

fn parse_input(raw_input: &str) -> Vec<(&str, &str)> {
    raw_input
        .lines()
        .map(|x| x.split_once('-').unwrap())
        .collect()
}

fn parse_graph<'a>(input: &'a [(&'a str, &'a str)]) -> HashMap<&'a str, Vec<&'a str>> {
    let mut g = HashMap::new();
    for (u, v) in input {
        g.entry(*u).or_insert(vec![]).push(*v);
        g.entry(*v).or_insert(vec![]).push(*u);
    }
    g
}

fn has_lowercase_duplicate(path: &[&str]) -> bool {
    let lc = path
        .iter()
        .filter(|x| x.chars().all(|y| y.is_ascii_lowercase()))
        .collect_vec();
    lc.len() != lc.into_iter().collect::<HashSet<_>>().len()
}

fn p(input: &[(&str, &str)], p2: bool) -> i32 {
    let g = parse_graph(input);
    let mut res = 0;

    let mut path = vec!["start"];
    let mut stack = vec![g["start"].clone()];

    while let Some(nbrs) = stack.pop() {
        if nbrs.is_empty() {
            path.pop();
            continue;
        }
        stack.push(nbrs[1..].to_vec());

        let curr = nbrs[0];
        // part1 condition
        if !p2 && curr.chars().all(|x| x.is_ascii_lowercase()) && path.contains(&curr) {
            continue;
        }
        // part2 condition
        if p2
            && (curr == "start"
                || path.contains(&curr)
                    && curr.chars().all(|x| x.is_ascii_lowercase())
                    && has_lowercase_duplicate(&path))
        {
            continue;
        }

        path.push(curr);

        if curr == "end" {
            res += 1;
            path.pop();
        } else {
            stack.push(g[&curr].clone());
        }
    }
    res
}
