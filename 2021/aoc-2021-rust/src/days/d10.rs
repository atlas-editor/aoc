use std::collections::HashMap;

pub fn p1() -> i32 {
    let input = parse_input();
    _p1(&input)
}

pub fn p2() -> u64 {
    let input = parse_input();
    _p2(&input)
}

fn parse_input<'a>() -> Vec<&'a str> {
    include_str!("../inputs/10.in").lines().collect()
}

fn _p1(lines: &[&str]) -> i32 {
    let vals = [(')', 3), (']', 57), ('}', 1197), ('>', 25137)]
        .into_iter()
        .collect::<HashMap<_, _>>();
    let mut res = 0;
    for line in lines {
        let mut stack = vec![];
        for ch in line.chars() {
            if [')', ']', '}', '>'].contains(&ch) {
                if stack.is_empty() {
                    res += vals[&ch];
                    break;
                }
                let pair = stack.pop().unwrap();
                let t = [pair, ch].iter().collect::<String>();
                if !["()", "[]", "{}", "<>"].contains(&t.as_str()) {
                    res += vals[&ch];
                    break;
                }
                continue;
            }
            stack.push(ch);
        }
    }
    res
}

fn _p2(lines: &[&str]) -> u64 {
    let vals = [('(', 1), ('[', 2), ('{', 3), ('<', 4)]
        .into_iter()
        .collect::<HashMap<_, _>>();
    let mut res = vec![];
    'outer: for line in lines {
        let mut stack = vec![];
        for ch in line.chars() {
            if [')', ']', '}', '>'].contains(&ch) {
                if stack.is_empty() {
                    continue 'outer;
                }
                let pair = stack.pop().unwrap();
                let t = [pair, ch].iter().collect::<String>();
                if !["()", "[]", "{}", "<>"].contains(&t.as_str()) {
                    continue 'outer;
                }
                continue;
            }
            stack.push(ch);
        }
        res.push(stack.iter().rev().fold(0, |acc, ch| acc * 5 + vals[ch]));
    }
    res.sort();
    res[(res.len() - 1) / 2]
}
