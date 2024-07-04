use crate::days::utils::*;
use itertools::Itertools;

pub fn p1(raw_input: &[u8]) -> i16 {
    let mut input = parse_input(raw_input);
    _p1(&mut input)
}

pub fn p2(raw_input: &[u8]) -> usize {
    let mut input = parse_input(raw_input);
    _p2(&mut input)
}

fn parse_input(raw_input: &[u8]) -> Matrix<i16> {
    Matrix::digits(raw_input)
}

fn neighbors(
    r_size: usize,
    c_size: usize,
    r: usize,
    c: usize,
    m: &Matrix<i16>,
    flashed: i16,
) -> Vec<(usize, usize)> {
    let mut nbrs = vec![];
    for (dr, dc) in [
        (1, 0),
        (1, 1),
        (0, 1),
        (-1, 1),
        (-1, 0),
        (-1, -1),
        (0, -1),
        (1, -1),
    ] {
        let rr = r as i8 + dr;
        let cc = c as i8 + dc;

        if rr >= 0
            && rr < r_size as i8
            && cc >= 0
            && cc < c_size as i8
            && m[rr as usize][cc as usize] != flashed
        {
            nbrs.push((rr as usize, cc as usize))
        }
    }

    nbrs
}

fn step(energy: &mut Matrix<i16>, idx: i16) -> i16 {
    let (r_size, c_size) = energy.shape;
    let mut step_flashes = 0;
    let mut stack = vec![];

    let flashed = -idx - 1;
    let threshold = flashed + 9;

    for p in (0..r_size).into_iter().cartesian_product(0..c_size) {
        if energy[p] <= threshold {
            continue;
        }
        stack.push(p);
        energy[p] = flashed;
        while let Some(q) = stack.pop() {
            step_flashes += 1;
            for qq in neighbors(r_size, c_size, q.0, q.1, &energy, flashed) {
                energy[qq] += 1;
                if energy[qq] > threshold {
                    stack.push(qq);
                    energy[qq] = flashed;
                }
            }
        }
    }
    step_flashes
}

fn _p1(energy: &mut Matrix<i16>) -> i16 {
    (0..100).map(|idx| step(energy, idx)).sum()
}

fn _p2(energy: &mut Matrix<i16>) -> usize {
    let full = (energy.shape.0 * energy.shape.1) as i16;
    (0..).find_position(|&idx| step(energy, idx) == full).unwrap().0 + 1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 1656);
        assert_eq!(p2(raw_input()), 195);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526"
    }
}
