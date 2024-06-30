use itertools::Itertools;
use std::collections::HashMap;

pub fn p1(raw_input: &str) -> u32 {
    let input = parse_input(raw_input);
    _p1(&input)
}

pub fn p2(raw_input: &str) -> u64 {
    let input = parse_input(raw_input);
    _p2(&input)
}

fn parse_input(raw_input: &str) -> Vec<&str> {
    raw_input.lines().collect()
}

enum LineType {
    Incomplete(Vec<char>),
    Corrupted(char),
}

impl LineType {
    fn from_str(line: &str) -> Self {
        let mut stack = vec![];
        for ch in line.chars() {
            if [')', ']', '}', '>'].contains(&ch) {
                if let Some(pair) = stack.pop() {
                    let t = format!("{pair}{ch}");
                    if !["()", "[]", "{}", "<>"].contains(&t.as_str()) {
                        return LineType::Corrupted(ch);
                    }
                } else {
                    return LineType::Corrupted(ch);
                }
            } else {
                stack.push(ch);
            }
        }
        LineType::Incomplete(stack)
    }
}

fn _p1(lines: &[&str]) -> u32 {
    let vals = [(')', 3), (']', 57), ('}', 1197), ('>', 25137)]
        .into_iter()
        .collect::<HashMap<_, _>>();
    lines
        .iter()
        .map(|line| match LineType::from_str(line) {
            LineType::Incomplete(_) => 0,
            LineType::Corrupted(ch) => vals[&ch],
        })
        .sum()
}

fn _p2(lines: &[&str]) -> u64 {
    let vals = [('(', 1), ('[', 2), ('{', 3), ('<', 4)]
        .into_iter()
        .collect::<HashMap<_, _>>();
    let scores = lines
        .iter()
        .filter_map(|line| match LineType::from_str(line) {
            LineType::Incomplete(stack) => {
                Some(stack.iter().rev().fold(0, |acc, ch| acc * 5 + vals[ch]))
            }
            LineType::Corrupted(_) => None,
        })
        .sorted()
        .collect_vec();
    scores[(scores.len() - 1) / 2]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 26397);
        assert_eq!(p2(raw_input()), 288957);
    }

    fn raw_input<'a>() -> &'a str {
        "[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]"
    }
}
