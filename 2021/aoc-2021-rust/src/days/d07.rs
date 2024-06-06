use itertools::Itertools;

pub fn p1(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    find_min(|t| t, &input)
}

pub fn p2(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    find_min(|t| (t * (t + 1)) / 2, &input)
}

fn parse_input(raw_input: &str) -> Vec<i32> {
    raw_input
        .lines()
        .next()
        .unwrap()
        .split(',')
        .map(|x| x.parse::<i32>().unwrap())
        .collect_vec()
}

fn find_min<T: Fn(i32) -> i32>(f: T, nums: &[i32]) -> i32 {
    let (a_ref, b_ref) = nums.iter().minmax().into_option().unwrap();
    let (mut a, mut b) = (*a_ref, *b_ref);
    let g = |i: i32| nums.iter().map(|n| f((n - i).abs())).sum::<i32>();
    while a != b {
        let mid = (a + b) / 2;
        if g((a + mid) / 2) <= g((mid + b) / 2) {
            b = mid;
        } else {
            a = mid;
        }
    }
    g(a)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 37);
        assert_eq!(p2(raw_input()), 168);
    }

    fn raw_input<'a>() -> &'a str {
        "16,1,2,0,4,2,7,1,2,14"
    }
}
