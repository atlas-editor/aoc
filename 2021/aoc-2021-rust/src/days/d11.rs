use crate::days::utils::*;

pub fn p1(raw_input: &[u8]) -> i16 {
    let mut input = parse_input(raw_input);
    p(&mut input, false)
}

pub fn p2(raw_input: &[u8]) -> i16 {
    let mut input = parse_input(raw_input);
    p(&mut input, true)
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

fn p(energy: &mut Matrix<i16>, p2: bool) -> i16 {
    let (r_size, c_size) = energy.shape;
    let area = (r_size * c_size) as i16;
    let mut flashes = 0;
    let mut step_flashes = 0;
    let mut stack = vec![];
    for i in 0.. {
        step_flashes = 0;
        let flashed = -i - 1;
        let threshold = flashed + 9;
        if i == 100 && !p2 {
            return flashes;
        }

        for r in 0..r_size {
            for c in 0..c_size {
                if energy[r][c] <= threshold {
                    continue;
                }
                stack.push((r, c));
                energy[r][c] = flashed;
                while let Some((r, c)) = stack.pop() {
                    step_flashes += 1;
                    for (rr, cc) in neighbors(r_size, c_size, r, c, &energy, flashed) {
                        energy[rr][cc] += 1;
                        if energy[rr][cc] > threshold {
                            stack.push((rr, cc));
                            energy[rr][cc] = flashed;
                        }
                    }
                }
            }
        }

        if step_flashes == area && p2 {
            return i + 1;
        }
        flashes += step_flashes;
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
