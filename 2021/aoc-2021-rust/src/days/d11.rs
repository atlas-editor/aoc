use crate::days::utils::*;
use itertools::Itertools;

pub fn p1(raw_input: &[u8]) -> i16 {
    let input = parse_input(raw_input);
    _p1(input)
}

pub fn p2(raw_input: &[u8]) -> usize {
    let input = parse_input(raw_input);
    _p2(input)
}

fn parse_input(raw_input: &[u8]) -> Matrix<i16> {
    Matrix::from_digits(raw_input)
}

fn neighbors(idx: usize, r_size: usize, c_size: usize) -> impl Iterator<Item = usize> {
    let idx = idx as i16;
    let r_size = r_size as i16;
    let c_size = c_size as i16;
    let (r, c) = (idx / c_size, idx % c_size);
    [
        (idx - c_size - 1, (r - 1, c - 1)),
        (idx - c_size, (r - 1, c)),
        (idx - c_size + 1, (r - 1, c + 1)),
        (idx - 1, (r, c - 1)),
        (idx + 1, (r, c + 1)),
        (idx + c_size - 1, (r + 1, c - 1)),
        (idx + c_size, (r + 1, c)),
        (idx + c_size + 1, (r + 1, c + 1)),
    ]
    .into_iter()
    .filter(move |(_, (rr, cc))| *rr >= 0 && *rr < r_size && *cc >= 0 && *cc < c_size)
    .map(|(i, _)| i as usize)
}

struct StepIterator<S, T, U> {
    matrix: Matrix<S>,
    step: T,
    stack: Vec<U>,
}

impl StepIterator<i16, i16, usize> {
    fn new(m: Matrix<i16>) -> Self {
        Self {
            matrix: m,
            step: 0,
            stack: vec![],
        }
    }

    fn flashed(&self) -> i16 {
        -self.step - 1
    }

    fn threshold(&self) -> i16 {
        self.flashed() + 9
    }
}

impl Iterator for StepIterator<i16, i16, usize> {
    type Item = i16;

    fn next(&mut self) -> Option<Self::Item> {
        let mut flashes = 0;
        for p in 0..self.matrix.len() {
            if self.matrix[p] <= self.threshold() {
                continue;
            }

            self.stack.push(p);
            self.matrix[p] = self.flashed();

            while let Some(q) = self.stack.pop() {
                flashes += 1;

                for qq in neighbors(q, self.matrix.r_size(), self.matrix.c_size()) {
                    if self.matrix[qq] == self.flashed() {
                        continue;
                    }

                    self.matrix[qq] += 1;
                    if self.matrix[qq] > self.threshold() {
                        self.stack.push(qq);
                        self.matrix[qq] = self.flashed();
                    }
                }
            }
        }
        self.step += 1;
        if flashes != self.matrix.len() as i16 {
            Some(flashes)
        } else {
            None
        }
    }
}

fn _p1(energy: Matrix<i16>) -> i16 {
    StepIterator::new(energy).take(100).sum()
}

fn _p2(energy: Matrix<i16>) -> usize {
    StepIterator::new(energy).count() + 1
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
