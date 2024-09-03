use crate::days::utils::atoi;

pub fn p1(raw_input: &[u8]) -> i32 {
    // works when there is an integer `k` s.t. `x0 <= (k*(k+1))/2 <= x1` AND y0 is negative
    let v = -parse_input(raw_input).2 - 1;
    (v * (v + 1)) / 2
}

pub fn p2(raw_input: &[u8]) -> i32 {
    let nums = parse_input(raw_input);
    _p2(nums)
}

type Target = (i32, i32, i32, i32);

fn parse_input(raw_input: &[u8]) -> Target {
    let mut target: Vec<i32> = Vec::with_capacity(4);
    let mut i = 0;
    while i < raw_input.len() {
        let start = i;
        while i < raw_input.len() && (raw_input[i] == b'-' || raw_input[i].is_ascii_digit()) {
            i += 1;
        }

        if start < i {
            target.push(atoi(&raw_input[start..i]));
        }

        i += 1;
    }
    (target[0], target[1], target[2], target[3])
}

fn simulate(mut vx: i32, mut vy: i32, target: Target) -> bool {
    let (x0, x1, y0, y1) = target;
    let mut curr = (0, 0);
    while curr.0 <= x1 && curr.1 >= y0 {
        if curr.0 >= x0 && curr.0 <= x1 && curr.1 >= y0 && curr.1 <= y1 {
            return true;
        }
        curr = (curr.0 + vx, curr.1 + vy);
        if vx > 0 {
            vx -= 1;
        }
        vy -= 1;
    }
    false
}

fn _p2(nums: Target) -> i32 {
    let (_, x1, y0, _) = nums;
    let mut res = 0;
    for vx in 0..x1 + 1 {
        for vy in y0..-y0 + 1 {
            if simulate(vx, vy, nums) {
                res += 1
            }
        }
    }
    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 45);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), 112);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"target area: x=20..30, y=-10..-5"
    }
}
