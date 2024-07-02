use crate::days::utils::*;
use crate::matrix;

pub fn p1(raw_input: &[u8]) -> i32 {
    let input = parse_input(raw_input);
    p(input, false)
}

pub fn p2(raw_input: &[u8]) -> i32 {
    let input = parse_input(raw_input);
    p(input, true)
}

fn parse_input(raw_input: &[u8]) -> Matrix<u8> {
    Matrix::digits(raw_input)
}

fn neighbors(r_size: usize, c_size: usize, r: usize, c: usize) -> Vec<(usize, usize)> {
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

        if rr >= 0 && rr < r_size as i8 && cc >= 0 && cc < c_size as i8 {
            nbrs.push((rr as usize, cc as usize))
        }
    }

    nbrs
}

fn p(mut energy: Matrix<u8>, p2: bool) -> i32 {
    let (r_size, c_size) = energy.shape;
    let mut flashed = matrix![false; r_size, c_size];
    let mut res = 0;
    for i in 0.. {
        if i == 100 && !p2 {
            return res;
        }
        let mut nines = vec![];
        for r in 0..r_size {
            for c in 0..c_size {
                energy[r][c] += 1;
                if energy[r][c] > 9 {
                    nines.push((r, c))
                }
            }
        }

        for nine in nines {
            if flashed[nine.0][nine.1] {
                continue;
            }
            let mut stack = vec![nine];
            flashed[nine.0][nine.1] = true;
            while let Some((r, c)) = stack.pop() {
                for (rr, cc) in neighbors(r_size, c_size, r, c) {
                    energy[rr][cc] += 1;
                    if energy[rr][cc] > 9 && !flashed[rr][cc] {
                        stack.push((rr, cc));
                        flashed[rr][cc] = true;
                    }
                }
            }
        }
        let mut t = 0;
        for r in 0..r_size {
            for c in 0..c_size {
                if flashed[r][c] {
                    energy[r][c] = 0;
                    t += 1;
                }
                flashed[r][c] = false;
            }
        }
        if t == r_size * c_size && p2 {
            return i + 1;
        }
        res += t as i32;
    }
    unreachable!()
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
