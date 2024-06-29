use std::collections::HashMap;

pub fn p1(raw_input: &str) -> i32 {
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

fn _p1(lines: &[&str]) -> i32 {
    let vals = [(')', 3), (']', 57), ('}', 1197), ('>', 25137)]
        .into_iter()
        .collect::<HashMap<_, _>>();
    let mut res = 0;
    for line in lines {
        let mut stack = vec![];
        for ch in line.chars() {
            if [')', ']', '}', '>'].contains(&ch) {
                if let Some(pair) = stack.pop() {
                    let t = format!("{pair}{ch}");
                    if !["()", "[]", "{}", "<>"].contains(&t.as_str()) {
                        res += vals[&ch];
                        break;
                    }
                } else {
                    res += vals[&ch];
                    break;
                }
            } else {
                stack.push(ch);
            }
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
                if let Some(pair) = stack.pop() {
                    let t = format!("{pair}{ch}");
                    if !["()", "[]", "{}", "<>"].contains(&t.as_str()) {
                        continue 'outer;
                    }
                } else {
                    continue 'outer;
                }
            } else {
                stack.push(ch);
            }
        }
        res.push(stack.iter().rev().fold(0, |acc, ch| acc * 5 + vals[ch]));
    }
    res.sort();
    res[(res.len() - 1) / 2]
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
