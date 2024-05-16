use std::collections::HashSet;

fn main() {
    let inp = include_str!("tmp.in").lines().collect::<Vec<_>>();
    println!("{}", p1(&inp));
    println!("{}", p2(&inp));
}

fn p1(entries: &[&str]) -> usize {
    entries
        .iter()
        .map(|e| {
            e.split('|')
                .last()
                .unwrap()
                .split_ascii_whitespace()
                .filter(|x| [2, 4, 3, 7].contains(&x.len()))
                .count()
        })
        .sum()
}

fn unique_by_len<'a>(codes: &'a [HashSet<&'a u8>], l: usize) -> HashSet<&'a u8> {
    codes.iter().find(|x| x.len() == l).unwrap().clone()
}

fn filter_by_predicate<'a, F: Fn(&'a HashSet<&'a u8>) -> bool>(
    codes: &'a [HashSet<&'a u8>],
    pred: F,
) -> HashSet<&'a u8> {
    codes.iter().find(|x| pred(x)).unwrap().clone()
}

fn set(s: &str) -> HashSet<&u8> {
    s.as_bytes().iter().collect()
}

fn p2(entries: &[&str]) -> usize {
    let mut res = 0;
    for e in entries {
        let codes = e
            .split('|')
            .next()
            .unwrap()
            .split_ascii_whitespace()
            .map(set)
            .collect::<Vec<_>>();

        let one = unique_by_len(&codes, 2);
        let four = unique_by_len(&codes, 4);
        let seven = unique_by_len(&codes, 3);
        let eight = unique_by_len(&codes, 7);

        let nine = filter_by_predicate(&codes, |x| x.len() == 6 && four.is_subset(x));
        let zero = filter_by_predicate(&codes, |x| x.len() == 6 && *x != nine && one.is_subset(x));
        let six = filter_by_predicate(&codes, |x| x.len() == 6 && *x != nine && *x != zero);
        let three = filter_by_predicate(&codes, |x| {
            x.len() == 5 && seven.is_subset(x) && nine.is_superset(x)
        });
        let five = filter_by_predicate(&codes, |x| {
            x.len() == 5 && *x != three && nine.is_superset(x)
        });
        let two = filter_by_predicate(&codes, |x| x.len() == 5 && *x != three && *x != five);

        let decoded = [zero, one, two, three, four, five, six, seven, eight, nine];

        res += e
            .split('|')
            .last()
            .unwrap()
            .split_ascii_whitespace()
            .map(|x| {
                decoded
                    .iter()
                    .position(|y| *y == set(x))
                    .unwrap()
                    .to_string()
            })
            .collect::<String>()
            .parse::<usize>()
            .unwrap();
    }

    res
}
