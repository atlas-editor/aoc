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
                .split_whitespace()
                .filter(|x| [2, 4, 3, 7].contains(&x.len()))
                .count()
        })
        .sum()
}

fn p2(nums: &[&str]) -> usize {
    0
}
