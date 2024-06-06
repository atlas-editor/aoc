use std::cmp::Ordering;
use itertools::Itertools;
use std::collections::BinaryHeap;

pub fn p1(raw_input: &str) -> i32 {
    let g = parse_input(raw_input);
    let (r_size, c_size) = (g.len(), g[0].len());
    dijkstra(&g, (0, 0), (r_size as i32 - 1, c_size as i32 - 1))
}

pub fn p2(raw_input: &str) -> i32 {
    let gg = extend_graph(&parse_input(raw_input));
    let (r_size, c_size) = (gg.len(), gg[0].len());
    dijkstra(&gg, (0, 0), (r_size as i32 - 1, c_size as i32 - 1))
}

fn parse_input(raw_input: &str) -> Vec<Vec<u8>> {
    raw_input
        .lines()
        .map(|x| {
            x.chars()
                .map(|y| y.to_digit(10).unwrap() as u8)
                .collect_vec()
        })
        .collect()
}

fn wrap(x: usize) -> u8 {
    if x <= 9 {
        x as u8
    } else {
        ((x % 10) + 1) as u8
    }
}

fn extend_graph(g: &Vec<Vec<u8>>) -> Vec<Vec<u8>> {
    let (r_size, c_size) = (g.len(), g[0].len());
    let mut gg = vec![vec![0; 5*c_size]; 5*r_size];

    for i in 0..r_size {
        for j in 0..5 {
            for k in 0..c_size {
                for l in 0..5 {
                    gg[i+(j*r_size)][k+(l*c_size)] = wrap(g[i][k] as usize + j + l)
                }
            }
        }
    }

    gg
}

#[derive(Eq, PartialEq)]
struct State {
    dist: i32,
    pos: (i32, i32),
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other.dist.cmp(&self.dist)
            .then_with(|| self.pos.cmp(&other.pos))
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn nbrs(r: usize, c: usize, u: (i32, i32)) -> Vec<(i32, i32)> {
    [(-1, 0), (0, 1), (1, 0), (0, -1)]
        .map(|x| (x.0 + u.0, x.1 + u.1))
        .into_iter()
        .filter(|&x| x.0 >= 0 && x.0 < r as i32 && x.1 >= 0 && x.1 < c as i32)
        .collect_vec()
}

fn dijkstra(g: &Vec<Vec<u8>>, start: (i32, i32), dest: (i32, i32)) -> i32 {
    let (r_size, c_size) = (g.len(), g[0].len());
    let mut dist = vec![vec![i32::MAX; c_size]; r_size];
    let mut q = BinaryHeap::new();

    dist[start.0 as usize][start.1 as usize] = 0;
    q.push(State{ dist: 0, pos: start});
    while let Some(State { dist: d, pos: u }) = q.pop() {
        if u == dest {
            return d;
        }
        if d > dist[u.0 as usize][u.1 as usize] {
            continue;
        }

        for v in nbrs(r_size, c_size, u) {
            let alt = dist[u.0 as usize][u.1 as usize] + g[v.0 as usize][v.1 as usize] as i32;
            if alt < dist[v.0 as usize][v.1 as usize] {
                dist[v.0 as usize][v.1 as usize] = alt;
                q.push(State{ dist: alt, pos: v});
            }
        }
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 40);
        assert_eq!(p2(raw_input()), 315);
    }

    fn raw_input<'a>() -> &'a str {
        "1163751742
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
