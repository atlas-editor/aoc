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

fn set(s: &str) -> HashSet<&u8> {
    s.as_bytes().iter().collect()
}

fn find_pop<T, F>(vec: &mut Vec<T>, predicate: F) -> Option<T>
where
    F: Fn(&T) -> bool,
{
    if let Some(index) = vec.iter().position(predicate) {
        Some(vec.remove(index))
    } else {
        None
    }
}

fn p2(entries: &[&str]) -> usize {
    let mut res = 0;
    for e in entries {
        let mut codes = e
            .split('|')
            .next()
            .unwrap()
            .split_ascii_whitespace()
            .map(set)
            .collect::<Vec<_>>();

        let one = find_pop(&mut codes, |x| x.len() == 2).unwrap();
        let four = find_pop(&mut codes, |x| x.len() == 4).unwrap();
        let seven = find_pop(&mut codes, |x| x.len() == 3).unwrap();
        let eight = find_pop(&mut codes, |x| x.len() == 7).unwrap();
        let nine = find_pop(&mut codes, |x| x.len() == 6 && four.is_subset(x)).unwrap();
        let zero = find_pop(&mut codes, |x| x.len() == 6 && one.is_subset(x)).unwrap();
        let six = find_pop(&mut codes, |x| x.len() == 6).unwrap();
        let three = find_pop(&mut codes, |x| x.len() == 5 && seven.is_subset(x)).unwrap();
        let five = find_pop(&mut codes, |x| x.len() == 5 && nine.is_superset(x)).unwrap();
        let two = codes.pop().unwrap();

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