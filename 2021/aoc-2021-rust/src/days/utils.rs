use std::collections::HashMap;
use std::fmt;
use std::fmt::{Debug, Display, Formatter};
use std::ops::{Add, Index, IndexMut, Mul, Neg};
use std::str::FromStr;

use bstr::ByteSlice;
use itertools::iproduct;
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

pub fn atopi<T: From<u8> + Mul<Output = T> + Add<Output = T>>(input: &[u8]) -> T {
    input.iter().fold(T::from(0), |res, digit| {
        T::from(10) * res + T::from(digit - 48)
    })
}

pub fn atopi2<T: From<u8> + Mul<Output = T> + Add<Output = T>>(input: &[u8]) -> T {
    input.iter().fold(T::from(0), |res, digit| {
        T::from(2) * res + T::from(digit - 48)
    })
}

pub fn atoi<T: From<u8> + Mul<Output = T> + Add<Output = T> + Neg<Output = T>>(input: &[u8]) -> T {
    if input[0] == b'-' {
        -atopi::<T>(&input[1..])
    } else {
        atopi::<T>(input)
    }
}

#[derive(Debug, Clone)]
pub struct Matrix<T> {
    pub shape: (usize, usize),
    pub data: Vec<T>,
}

impl<T: Display> Display for Matrix<T> {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        for r in 0..self.r_size() {
            for c in 0..self.c_size() {
                write!(f, "{}", self[(r, c)])?;
            }
            if r < self.r_size() - 1 {
                writeln!(f)?;
            }
        }
        Ok(())
    }
}

impl<T> Matrix<T> {
    pub fn from_shape_and_data(shape: (usize, usize), data: Vec<T>) -> Self {
        if data.len() != shape.0 * shape.1 {
            panic!("invalid shape for data")
        }
        Self { shape, data }
    }
    pub fn from_repr<F: Fn(u8) -> T>(repr: &[u8], transform: F) -> Self {
        let mut r_size = 0usize;
        let mut data = vec![];
        for line in repr.lines() {
            r_size += 1;
            for &el in line {
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

#[derive(Debug, Clone)]
pub struct Array3D<T> {
    pub shape: (usize, usize, usize),
    pub data: Vec<T>,
}

impl<T> Array3D<T> {
    pub fn iter(&self) -> impl Iterator<Item = ((usize, usize, usize), &T)> {
        iproduct!(0..self.shape.0, 0..self.shape.1, 0..self.shape.2).map(|pos| (pos, &self[pos]))
    }
}

impl<T> Index<(usize, usize, usize)> for Array3D<T> {
    type Output = T;

    fn index(&self, index: (usize, usize, usize)) -> &Self::Output {
        &self.data[self.shape.0 * self.shape.1 * index.2 + self.shape.1 * index.0 + index.1]
    }
}

impl<T> IndexMut<(usize, usize, usize)> for Array3D<T> {
    fn index_mut(&mut self, index: (usize, usize, usize)) -> &mut Self::Output {
        &mut self.data[self.shape.0 * self.shape.1 * index.2 + self.shape.1 * index.0 + index.1]
    }
}

#[macro_export]
macro_rules! array3d {
    [$val:expr; $X:expr, $Y:expr, $Z:expr] => {
        Array3D{shape: ($X, $Y, $Z), data: vec![$val; $X*$Y*$Z]}
    };
    [$val:expr; $D:expr] => {
        matrix![$val; $D, $D, $D]
    };
}

#[derive(Debug, Clone)]
pub struct ByteMap<T> {
    map: [T; 256],
}

impl<T: Default> ByteMap<T> {
    pub fn new() -> Self {
        Self {
            map: [(); 256].map(|_| T::default()),
        }
    }
}

impl<T> Index<u8> for ByteMap<T> {
    type Output = T;

    fn index(&self, index: u8) -> &Self::Output {
        &self.map[index as usize]
    }
}

impl<T> IndexMut<u8> for ByteMap<T> {
    fn index_mut(&mut self, index: u8) -> &mut Self::Output {
        &mut self.map[index as usize]
    }
}

#[macro_export]
macro_rules! bmap {
    {$val:expr} => {
        ByteMap{map: [(); 256].map(|_| $val)}
    };
    {$($key:expr => $value:expr,)+} => { bmap!{$($key => $value),+} };
    {$($key:expr => $value:expr),*} => {
        {
            let mut _map = ByteMap::new();
            $(
                _map[$key] = $value;
            )*
            _map
        }
    };
}

pub type ByteSet = ByteMap<bool>;

impl ByteSet {
    pub fn insert(&mut self, value: u8) {
        self[value] = true;
    }

    pub fn remove(&mut self, value: u8) {
        self[value] = false;
    }

    pub fn contains(&self, value: u8) -> bool {
        self[value] == true
    }
}

#[macro_export]
macro_rules! bset {
    {$($key:expr,)+} => { bset!{$($key),+} };
    {$($key:expr),*} => {
        {
            let mut _set = ByteSet::new();
            $(
                _set[$key] = true;
            )*
            _set
        }
    };
}

pub type ByteGraph = ByteMap<Vec<u8>>;

impl ByteGraph {
    pub fn neighbors(&self, u: u8) -> impl Iterator<Item = &u8> {
        self[u].iter()
    }

    pub fn insert(&mut self, u: u8, v: u8) {
        if !self[u].contains(&v) {
            self[u].push(v);
        }
        if !self[v].contains(&u) {
            self[v].push(u);
        }
    }
}

#[macro_export]
macro_rules! bgraph {
    ($(($key:expr, $value:expr),)+) => { bgraph!($(($key, $value)),+) };
    ($(($key:expr, $value:expr)),*) => {
        {
            let mut _graph = ByteGraph::new();
            $(
                _graph.insert($key, $value);
            )*
            _graph
        }
    };
}

pub fn parse_ints<
    const N: usize,
    T: From<u8> + Mul<Output = T> + Add<Output = T> + Neg<Output = T> + Debug,
>(
    input: &[u8],
) -> [T; N] {
    let mut nums: Vec<T> = Vec::with_capacity(N);
    let mut i = 0;
    while i < input.len() {
        let start = i;
        while i < input.len() && (input[i] == b'-' || input[i].is_ascii_digit()) {
            i += 1;
        }

        if start < i {
            nums.push(atoi(&input[start..i]));
        }

        i += 1;
    }
    nums.try_into().unwrap()
}


pub struct BidirectionalMap<U, V> {
    forward: HashMap<U, V>,
    reverse: HashMap<V, U>,
}

impl<V: Eq + std::hash::Hash + Clone> BidirectionalMap<usize, V> {
    pub fn from_vec(values: Vec<V>) -> BidirectionalMap<usize, V> {
        let mut forward = HashMap::new();
        let mut reverse = HashMap::new();
        for (idx, value) in values.iter().enumerate() {
            forward.insert(idx, value.to_owned());
            reverse.insert(value.to_owned(), idx);
        }

        Self{forward, reverse}
    }

    pub fn get(&self, idx: usize) -> &V {
        &self.forward[&idx]
    }

    pub fn get_rev(&self, value: V) -> usize {
        self.reverse[&value]
    }
}