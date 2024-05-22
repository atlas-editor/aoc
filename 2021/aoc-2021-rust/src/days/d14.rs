use itertools::Itertools;
use std::collections::HashMap;
use std::fmt::Write;

pub fn p1(raw_input: &str) -> i64 {
    let (template, rules) = parse_input(raw_input);
    _p1(template, &rules)
}

pub fn p2(raw_input: &str) -> i64 {
    -1
}

fn parse_input(raw_input: &str) -> (String, HashMap<String, String>) {
    let input = raw_input.split_once("\n\n").unwrap();
    let template = input.0.to_string();
    let rules = input
        .1
        .lines()
        .map(|x| x.split_once(" -> ").unwrap())
        .map(|x| (x.0.to_string(), x.1.to_string()))
        .collect();
    (template, rules)
}

fn apply_rules(template: String, rules: &HashMap<String, String>) -> String {
    template
        .chars()
        .tuple_windows::<(char, char)>()
        .fold(String::new(), |mut acc, x| {
            let _ = write!(
                acc,
                "{}{}",
                x.0,
                *rules
                    .get(format!("{}{}", x.0, x.1).as_str())
                    .unwrap_or(&String::new())
            );
            acc
        })
        + &template.chars().last().unwrap().to_string()
}

fn _p1(mut template: String, rules: &HashMap<String, String>) -> i64 {
    for _ in 0..10 {
        template = apply_rules(template, rules);
    }

    let counts = template.chars().counts();
    let mm = counts.values().minmax().into_option().unwrap();

    *mm.1 as i64 - *mm.0 as i64
}
