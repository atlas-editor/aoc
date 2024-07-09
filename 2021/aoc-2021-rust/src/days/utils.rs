use bstr::ByteSlice;
use std::fmt;
use std::fmt::{Debug, Display};
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
}

pub type ByteGraph = ByteMap<Vec<u8>>;

impl ByteGraph {
    fn neighbors(&self, u: u8) -> impl Iterator<Item = &u8> {
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

    fn from_pairs(pairs: Vec<(u8, u8)>) -> Self {
        let mut _graph = Self::new();
        for (u, v) in pairs {
            if !_graph[u].contains(&v) {
                _graph[u].push(v);
            }
            if !_graph[v].contains(&u) {
                _graph[v].push(u);
            }
        }
        _graph
    }

    fn dfs_custom(&self, u: u8, v: u8, visited: &mut ByteMap<i32>, allowed: i32, has_duplicate: &mut bool) -> u32 {
        if u & 1 == 0 {
            visited[u] += 1;
            if visited[u] == allowed {
                *has_duplicate = true;
            }
        }
        let mut count = 0;
        if u == v {
            count += 1;
        } else {
            for &w in self.neighbors(u) {
                if w & 1 == 1 || visited[w] < allowed && !*has_duplicate {
                    count += self.dfs_custom(w, v, visited, allowed, has_duplicate);
                }
            }
        }
        if u & 1 == 0 {
            visited[u] -= 1;
            *has_duplicate = false;
        }
        count
    }

    pub fn paths_count(&self, u: u8, v: u8, allowed: i32) -> u32 {
        self.dfs_custom(u, v, &mut bmap! {0 => allowed-1}, allowed, &mut false)
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
