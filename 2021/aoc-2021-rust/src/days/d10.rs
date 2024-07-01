use itertools::Itertools;
use maplit::hashmap;

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
    Incomplete(Vec<u8>),
    Corrupted(u8),
}

trait Typer<T> {
    fn r#type(&self) -> T;
}

impl Typer<LineType> for &str {
    fn r#type(&self) -> LineType {
        let mut stack = vec![];
        for b in self.as_bytes() {
            if [b')', b']', b'}', b'>'].contains(&b) {
                if let Some(pair) = stack.pop() {
                    if !bracket_match(pair, *b) {
                        return LineType::Corrupted(*b);
                    }
                } else {
                    return LineType::Corrupted(*b);
                }
            } else {
                stack.push(*b);
            }
        }
        LineType::Incomplete(stack)
    }
}

fn bracket_match(b0: u8, b1: u8) -> bool {
    match b0 {
        b'(' => b1 == b')',
        b'[' => b1 == b']',
        b'{' => b1 == b'}',
        b'<' => b1 == b'>',
        _ => {
            panic!("unexpected char")
        }
    }
}

fn _p1(lines: &[&str]) -> u32 {
    let vals = hashmap! {b')' => 3, b']' => 57, b'}' => 1197, b'>' => 25137};
    lines
        .iter()
        .map(|line| match line.r#type() {
            LineType::Incomplete(_) => 0,
            LineType::Corrupted(ch) => vals[&ch],
        })
        .sum()
}

fn _p2(lines: &[&str]) -> u64 {
    let vals = hashmap! {b'(' => 1, b'[' => 2, b'{' => 3, b'<' => 4};
    let scores = lines
        .iter()
        .filter_map(|line| match line.r#type() {
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
