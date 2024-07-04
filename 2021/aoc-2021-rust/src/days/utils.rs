use std::fmt;
use std::fmt::{Debug, Display};
use std::io::BufRead;
use std::ops::{Index, IndexMut};
use std::str::FromStr;

use regex::Regex;

pub fn ints<T: FromStr>(input: &str) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    let re = Regex::new(r"-?\d+").unwrap();
    re.find_iter(input)
        .map(|m| m.as_str().parse().unwrap())
        .collect()
}

#[derive(Debug, Clone)]
pub struct Matrix<T> {
    pub shape: (usize, usize),
    pub data: Vec<T>,
}

impl<T> Matrix<T> {
    pub fn from_repr<F: Fn(u8) -> T>(repr: &[u8], transform: F) -> Self {
        let mut r_size = 0usize;
        let mut data = vec![];
        for line in repr.lines() {
            r_size += 1;
            for &el in line.unwrap().as_bytes() {
                data.push(transform(el));
            }
        }
        Self {
            shape: (r_size, data.len() / r_size),
            data,
        }
    }

    pub fn r_size(&self) -> usize {
        self.shape.0
    }

    pub fn c_size(&self) -> usize {
        self.shape.1
    }

    pub fn len(&self) -> usize {
        self.r_size() * self.c_size()
    }
}

impl Matrix<u8> {
    pub fn bytes(repr: &[u8]) -> Self {
        Self::from_repr(repr, |x| x)
    }
}

impl<T: From<u8>> Matrix<T> {
    pub fn from_digits(repr: &[u8]) -> Self {
        Self::from_repr(repr, |x| T::from(x - 48))
    }
}

impl<T: Debug> Display for Matrix<T> {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for i in 0..self.shape.0 {
            writeln!(f, "{:?}", &self[i])?;
        }
        Ok(())
    }
}

impl<T> Index<usize> for Matrix<T> {
    type Output = T;

    fn index(&self, index: usize) -> &Self::Output {
        &self.data[index]
    }
}

impl<T> IndexMut<usize> for Matrix<T> {
    fn index_mut(&mut self, index: usize) -> &mut Self::Output {
        &mut self.data[index]
    }
}

impl<T> Index<(usize, usize)> for Matrix<T> {
    type Output = T;

    fn index(&self, index: (usize, usize)) -> &Self::Output {
        &self.data[index.0 * self.shape.1 + index.1]
    }
}

impl<T> IndexMut<(usize, usize)> for Matrix<T> {
    fn index_mut(&mut self, index: (usize, usize)) -> &mut Self::Output {
        &mut self.data[index.0 * self.shape.1 + index.1]
    }
}

#[macro_export]
macro_rules! matrix {
    [$val:expr; $R:expr, $C:expr] => {
        Matrix{shape: ($R, $C), data: vec![$val; $R*$C]}
    };
    [$val:expr; $D:expr] => {
        matrix![$val; $D, $D]
    };
}

#[cfg(test)]
mod tests {
    use super::*;

    fn inner() -> impl Iterator {
        [(-1, 1), (2, 2), (-11, 1), (111, 1)].iter()
    }

    #[test]
    fn it_works() {
        for i in inner() {
            let k = inner().next().unwrap();
            println!("ok");
        }
    }
}
