use crate::days::utils::Matrix;
use crate::matrix;
use itertools::Itertools;
use std::collections::BinaryHeap;

pub fn p1(raw_input: &[u8]) -> i32 {
    let g = parse_input(raw_input);
    dijkstra(&g)
}

pub fn p2(raw_input: &[u8]) -> i32 {
    let gg = extend_graph(&parse_input(raw_input));
    dijkstra(&gg)
}

fn parse_input(raw_input: &[u8]) -> Matrix<i32> {
    Matrix::from_digits(raw_input)
}

fn wrap(x: i32) -> i32 {
    if x <= 9 {
        x
    } else {
        (x % 10) + 1
    }
}

fn extend_graph(g: &Matrix<i32>) -> Matrix<i32> {
    let mut gg = matrix![0; 5*g.c_size(), 5*g.r_size()];
    for i in 0..g.r_size() {
        for j in 0..5 {
            for k in 0..g.c_size() {
                for l in 0..5 {
                    gg[(i + (j * g.r_size()), k + (l * g.c_size()))] =
                        wrap(g[(i, k)] + j as i32 + l as i32)
                }
            }
        }
    }
    gg
}

fn nbrs(r: usize, c: usize, u: (usize, usize)) -> impl Iterator<Item = (usize, usize)> {
    [(-1, 0), (0, 1), (1, 0), (0, -1)]
        .map(|x| (x.0 + (u.0 as i32), x.1 + (u.1 as i32)))
        .into_iter()
        .filter(move |&x| x.0 >= 0 && x.0 < r as i32 && x.1 >= 0 && x.1 < c as i32)
        .map(|y| (y.0 as usize, y.1 as usize))
}

fn dijkstra(g: &Matrix<i32>) -> i32 {
    let mut dist = matrix![i32::MAX; g.r_size(), g.c_size()];
    let mut heap = BinaryHeap::new();

    dist[(0, 0)] = 0;
    heap.push((0, (0, 0)));
    while let Some((d, u)) = heap.pop() {
        if u == (g.r_size() - 1, g.c_size() - 1) {
            return -d;
        }
        if -d > dist[(u.0, u.1)] {
            continue;
        }

        for v in nbrs(g.r_size(), g.c_size(), (u.0, u.1)) {
            let alt = dist[(u.0, u.1)] + g[(v.0, v.1)];
            if alt < dist[(v.0, v.1)] {
                dist[(v.0, v.1)] = alt;
                heap.push((-alt, (v.0, v.1)));
            }
        }
    }
    panic!("destination not reached")
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 40);
        assert_eq!(p2(raw_input()), 315);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581"
    }
}
