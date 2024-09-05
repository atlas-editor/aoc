use crate::days::utils::Matrix;
use bstr::ByteSlice;
use itertools::Itertools;
use std::fmt::{Display, Formatter};
use std::ops::Index;

pub fn p1(raw_input: &[u8]) -> usize {
    let (ie_algorithm, mut input_image) = parse_input(raw_input);
    input_image.apply_ie_algorithm(&ie_algorithm, 2);
    input_image.lit_pixels()
}

pub fn p2(raw_input: &[u8]) -> usize {
    let (ie_algorithm, mut input_image) = parse_input(raw_input);
    input_image.apply_ie_algorithm(&ie_algorithm, 50);
    input_image.lit_pixels()
}

fn parse_input(raw_input: &[u8]) -> (IEAlgorithm, Image) {
    let (ie_algorithm, input_image) = raw_input.split_once_str(b"\n\n").unwrap();

    (
        IEAlgorithm::from_input(ie_algorithm),
        Image::from_input(input_image),
    )
}

#[derive(Debug, Copy, Clone)]
enum Pixel {
    Light,
    Dark,
}

#[derive(Debug, Clone, Copy)]
struct IEAlgorithm {
    algorithm: [Pixel; 512],
}

impl IEAlgorithm {
    fn from_input(input: &[u8]) -> Self {
        let mut algorithm = [Pixel::Dark; 512];

        for (i, &byte) in input.iter().enumerate() {
            if byte == b'#' {
                algorithm[i] = Pixel::Light
            }
        }

        Self { algorithm }
    }

    fn full(&self, pixel: Pixel) -> Pixel {
        match pixel {
            Pixel::Light => self.algorithm[511],
            Pixel::Dark => self.algorithm[0],
        }
    }
}

impl Index<[Pixel; 9]> for IEAlgorithm {
    type Output = Pixel;

    fn index(&self, index: [Pixel; 9]) -> &Self::Output {
        &self.algorithm[index.iter().fold(0usize, |acc, p| match p {
            Pixel::Light => acc << 1 | 1,
            Pixel::Dark => acc << 1,
        })]
    }
}

#[derive(Debug)]
struct Image {
    image: Matrix<Pixel>,
    infinite_pixel: Pixel,
}

impl Image {
    fn from_input(input: &[u8]) -> Self {
        Self {
            image: Matrix::from_repr(
                input,
                |x| if x == b'#' { Pixel::Light } else { Pixel::Dark },
            ),
            infinite_pixel: Pixel::Dark,
        }
    }

    fn apply_ie_algorithm_once(&mut self, ie_algorithm: &IEAlgorithm) {
        let mut new_data =
            Vec::with_capacity((self.image.r_size() + 2) * (self.image.c_size() + 2));
        for i in -1..(self.image.r_size() as i32) + 1 {
            for j in -1..(self.image.c_size() as i32) + 1 {
                let mut neighborhood = [self.infinite_pixel; 9];
                for (idx, n) in neighbors(self.image.r_size(), self.image.c_size(), (i, j)) {
                    neighborhood[idx] = self.image[n]
                }
                new_data.push(ie_algorithm[neighborhood])
            }
        }
        let new_image = Matrix::from_shape_and_data(
            (self.image.r_size() + 2, self.image.c_size() + 2),
            new_data,
        );

        self.image = new_image;
        self.infinite_pixel = ie_algorithm.full(self.infinite_pixel);
    }

    fn apply_ie_algorithm(&mut self, ie_algorithm: &IEAlgorithm, count: usize) {
        for _ in 0..count {
            self.apply_ie_algorithm_once(ie_algorithm);
        }
    }

    fn lit_pixels(&self) -> usize {
        self.image
            .data
            .iter()
            .filter(|x| matches!(x, Pixel::Light))
            .count()
    }
}

fn neighbors(
    r: usize,
    c: usize,
    index: (i32, i32),
) -> impl Iterator<Item = (usize, (usize, usize))> {
    [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 0),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ]
    .iter()
    .enumerate()
    .map(move |(idx, x)| (idx, (x.0 + index.0, x.1 + index.1)))
    .filter(move |(idx, x)| x.0 >= 0 && x.0 < r as i32 && x.1 >= 0 && x.1 < c as i32)
    .map(|(idx, y)| (idx, (y.0 as usize, y.1 as usize)))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 35);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), 3351);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###"
    }
}
