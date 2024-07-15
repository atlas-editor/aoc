use bstr::ByteSlice;
use itertools::Itertools;

use crate::days::utils::{ByteGraph, ByteMap};

pub fn p1(raw_input: &[u8]) -> u32 {
    let graph = parse_input(raw_input);
    graph.paths_count(0, 1)
}

pub fn p2(raw_input: &[u8]) -> u32 {
    let graph = parse_input(raw_input);
    graph.paths_count2(0, 1)
}

fn assign_u8<'a>(name: &[u8], i: u8, assignment: &mut [u8; 65536]) -> u8 {
    match name {
        b"start" => 0,
        b"end" => 1,
        other => {
            let mut key = (other[0] as usize) << 8;
            if other.len() == 2 {
                key += other[1] as usize;
            }
            if assignment[key] != 0 {
                assignment[key]
            } else {
                if other[0] >= b'a' {
                    assignment[key] = 2 * i + 2;
                    assignment[key]
                } else {
                    assignment[key] = 2 * i + 3;
                    assignment[key]
                }
            }
        }
    }
}

fn parse_input(raw_input: &[u8]) -> ByteGraph {
    let mut graph = ByteGraph::new();
    let mut assigned = [0u8; 65536];
    raw_input
        .lines()
        .map(|x| x.split_once_str(b"-").unwrap())
        .enumerate()
        .for_each(|(i, (u, v))| {
            graph.insert(
                assign_u8(u, (2 * i) as u8, &mut assigned),
                assign_u8(v, (2 * i + 1) as u8, &mut assigned),
            );
        });
    graph
}

impl ByteGraph {
    fn dfs_custom<F: Fn(&ByteMap<i16>, u8) -> bool>(
        &self,
        u: u8,
        v: u8,
        visited: &mut ByteMap<i16>,
        visit_fn: &F,
    ) -> u32 {
        visited[u] += 1;
        let mut count = 0;
        if u == v {
            count += 1;
        } else {
            for &w in self.neighbors(u) {
                if w & 1 == 1 || visit_fn(visited, w) {
                    count += self.dfs_custom(w, v, visited, visit_fn);
                }
            }
        }
        visited[u] -= 1;
        count
    }
    fn paths_count(&self, u: u8, v: u8) -> u32 {
        self.dfs_custom(u, v, &mut ByteMap::new(), &(|visited, w| visited[w] == 0))
    }

    fn paths_count2(&self, u: u8, v: u8) -> u32 {
        self.dfs_custom(
            u,
            v,
            &mut ByteMap::new(),
            &(|visited, w| {
                w > 0 && (visited[w] == 0 || (2..255).step_by(2).all(|i| visited[i] < 2))
            }),
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 226);
        assert_eq!(p2(raw_input()), 3509);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW"
    }
}
