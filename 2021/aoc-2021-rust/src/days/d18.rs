use regex::Regex;

pub fn p1(raw_input: &str) -> i32 {
    let reprs = parse_input(raw_input);
    magnitude(sum(reprs))
}

pub fn p2(raw_input: &str) -> i32 {
    let reprs = parse_input(raw_input);
    reprs
        .iter()
        .enumerate()
        .map(|(i, ri)| {
            reprs
                .iter()
                .enumerate()
                .map(|(j, rj)| {
                    if i != j {
                        magnitude(sum(vec![ri.clone(), rj.clone()]))
                    } else {
                        -1
                    }
                })
                .max()
                .unwrap()
        })
        .max()
        .unwrap()
}

fn parse_input(raw_input: &str) -> Vec<String> {
    raw_input.lines().map(|x| x.to_string()).collect()
}

fn sum(reprs: Vec<String>) -> String {
    reprs[1..]
        .iter()
        .fold(reprs[0].clone(), |acc, x| reduce(add(acc, x.clone())))
}

fn reduce(mut repr: String) -> String {
    loop {
        if let Some(e) = explode(repr.clone()) {
            repr = e;
            continue;
        }

        if let Some(s) = split(repr.clone()) {
            repr = s;
            continue;
        }
        break;
    }
    repr
}

fn split_pair(repr: String) -> Option<(String, String)> {
    let mut stack = 0;
    for (idx, &e) in repr[1..].as_bytes().iter().enumerate() {
        if e == b'[' {
            stack += 1;
        } else if e == b']' {
            stack -= 1;
        } else if stack == 0 && e == b',' {
            return Some((
                repr[1..idx + 1].to_string(),
                repr[idx + 2..repr.len() - 1].to_string(),
            ));
        }
    }
    None
}

fn magnitude(repr: String) -> i32 {
    if let Ok(a) = repr.parse::<i32>() {
        a
    } else {
        let (left, right) = split_pair(repr).unwrap();
        3 * magnitude(left) + 2 * magnitude(right)
    }
}

fn add(n0: String, n1: String) -> String {
    format!("[{n0},{n1}]")
}

fn explode(n: String) -> Option<String> {
    let mut stack = 0;
    for (i, &e) in n.as_bytes().iter().enumerate() {
        if e == b'[' {
            stack += 1;
        } else if e == b']' {
            stack -= 1;
        }

        if stack == 5 {
            let end = n[i..].find(|x| x == ']').unwrap() + i;

            let pair = n[i + 1..end].split_once(',').unwrap();
            let (a0, a1) = (
                pair.0.parse::<i32>().unwrap(),
                pair.1.parse::<i32>().unwrap(),
            );

            let mut left = n[..i].to_string();
            let mut right = n[end + 1..].to_string();

            let re = Regex::new(r"\d+").unwrap();

            if let Some(m) = re.find_iter(left.as_str()).last() {
                let val = m.as_str().parse::<i32>().unwrap() + a0;
                left = format!("{}{val}{}", &left[..m.start()], &left[m.end()..]);
            }

            if let Some(m) = re.find(right.as_str()) {
                let val = m.as_str().parse::<i32>().unwrap() + a1;
                right = format!("{}{val}{}", &right[..m.start()], &right[m.end()..]);
            }
            return Some(format!("{}0{}", left, right));
        }
    }
    None
}

fn split(n: String) -> Option<String> {
    let re = Regex::new(r"\d+").unwrap();
    if let Some(m) = re
        .clone()
        .find_iter(n.as_str())
        .find(|m| m.as_str().parse::<i32>().unwrap() >= 10)
    {
        let n0 = m.as_str().parse::<i32>().unwrap() / 2;
        let n1 = (m.as_str().parse::<i32>().unwrap() as f64 / 2.0).ceil() as i32;
        Some(format!("{}[{n0},{n1}]{}", &n[..m.start()], &n[m.end()..]))
    } else {
        None
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 4140);
        assert_eq!(p2(raw_input()), 3993);
    }

    fn raw_input<'a>() -> &'a str {
        "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"
    }
}
