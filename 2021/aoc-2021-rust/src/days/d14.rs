use std::collections::HashMap;

pub fn p1(raw_input: &str) -> usize {
    let (template, rules) = parse_input(raw_input);
    _p(template, &rules, 10)
}

pub fn p2(raw_input: &str) -> usize {
    let (template, rules) = parse_input(raw_input);
    _p(template, &rules, 40)
}

fn parse_input(raw_input: &str) -> (Vec<usize>, HashMap<(usize, usize), usize>) {
    let chars = "BCFHKNOPSV";
    let input = raw_input.split_once("\n\n").unwrap();
    let template = input.0.chars().map(|x| chars.find(x).unwrap()).collect();
    let rules = input
        .1
        .lines()
        .map(|x| x.split_once(" -> ").unwrap())
        .map(|x| {
            (
                (
                    chars.find(x.0.chars().next().unwrap()).unwrap(),
                    chars.find(x.0.chars().nth(1).unwrap()).unwrap(),
                ),
                chars.find(x.1).unwrap(),
            )
        })
        .collect();
    (template, rules)
}

fn add_vectors(vectors: &[&Vec<usize>]) -> Vec<usize> {
    let mut result = vectors[0].clone();
    for i in 1..vectors.len() {
        for j in 0..result.len() {
            result[j] += vectors[i][j];
        }
    }
    result
}

fn dp(rules: &HashMap<(usize, usize), usize>, d: usize) -> Vec<Vec<Vec<usize>>> {
    let zero = vec![0; 10];
    let mut table = vec![vec![vec![zero.clone(); 10]; 10]; d + 1];
    for i in 1..d + 1 {
        for j0 in 0..10 {
            for j1 in 0..10 {
                if !rules.contains_key(&(j0, j1)) {
                    table[i][j0][j1] = zero.clone();
                    continue;
                }
                let mid = rules.get(&(j0, j1)).unwrap();
                let c1 = table[i - 1][j0][*mid].clone();
                let c2 = table[i - 1][*mid][j1].clone();
                let mut c = zero.clone();
                c[*mid] += 1;
                table[i][j0][j1] = add_vectors(&[&c, &c1, &c2]);
            }
        }
    }
    table[d].clone()
}

fn _p(template: Vec<usize>, rules: &HashMap<(usize, usize), usize>, d: usize) -> usize {
    let table = dp(rules, d);
    let mut result = vec![0; 10];
    for digit in template.clone() {
        result[digit] += 1;
    }
    for i in 0..template.len() - 1 {
        let v = table[template[i]][template[i + 1]].clone();
        result = add_vectors(&[&result, &v]);
    }
    let max_ = result.iter().max().unwrap();
    let min_ = result.iter().filter(|&&x| x != 0).min().unwrap();
    max_ - min_
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 1588);
        assert_eq!(p2(raw_input()), 2188189693529);
    }

    fn raw_input<'a>() -> &'a str {
        "NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"
    }
}
