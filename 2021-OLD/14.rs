use std::collections::HashMap;

use itertools::Itertools;

fn main() {
    let input = include_str!("tmp.test").split_once("\n\n").unwrap();
    let template = input.0.to_string();
    let rules = input
        .1
        .lines()
        .map(|x| x.split_once(" -> ").unwrap())
        .map(|x| (x.0.to_string(), x.1.to_string()))
        .collect();
    println!("{}", p1(template.clone(), &rules));
    println!("{}", p2(template.clone(), &rules));
}

fn apply_rules(template: String, rules: &HashMap<String, String>) -> String {
    template
        .chars()
        .tuple_windows::<(char, char)>()
        .map(|x| {
            format!(
                "{}{}",
                x.0,
                *rules
                    .get(format!("{}{}", x.0, x.1).as_str())
                    .unwrap_or(&String::new())
            )
        })
        .collect::<String>()
        + &(template.chars().last().unwrap()).to_string()
}

fn apply_rules2(template: String, rules: &HashMap<String, String>) -> i64 {
    let mut xx = HashMap::new();
    for (k, v) in rules {
        for (ch, s) in v.chars().counts() {
            *xx.entry(k.clone())
                .or_insert(HashMap::new())
                .entry(ch)
                .or_insert(0) += s;
        }
    }

    template
        .chars()
        .tuple_windows::<(char, char)>()
        .map(|x| {
            format!(
                "{}{}",
                x.0,
                *rules
                    .get(format!("{}{}", x.0, x.1).as_str())
                    .unwrap_or(&String::new())
            )
        })
        .collect::<String>()
        + &(template.as_bytes()[template.len() - 1] as char).to_string();

    -1
}

fn p1(mut template: String, rules: &HashMap<String, String>) -> i64 {
    for _ in 0..10 {
        template = apply_rules(template, &rules);
    }

    let counts = template.chars().counts();
    let mm = counts.values().minmax().into_option().unwrap();

    *mm.1 as i64 - *mm.0 as i64
}

fn p2(mut template: String, rules: &HashMap<String, String>) -> i64 {
    let mut extended_rules: HashMap<String, String> = HashMap::new();
    for k in rules.keys() {
        let mut curr = k.to_string();
        for _ in 0..20 {
            curr = apply_rules(curr, &rules);
        }
        extended_rules.insert(k.to_string(), curr);
    }
    println!("first step done..");
    template = apply_rules(template, &extended_rules);
    // template = apply_rules(template, &extended_rules);

    let counts = template.chars().counts();
    let mm = counts.values().minmax().into_option().unwrap();

    *mm.1 as i64 - *mm.0 as i64
}
