use itertools::Itertools;

pub fn p1(raw_input: &str) -> i32 {
    let input = parse_input(raw_input);
    _p1(&input)
}

pub fn p2(raw_input: &str) -> usize {
    let input = parse_input(raw_input);
    _p2(&input)
}

fn parse_input(raw_input: &str) -> Vec<Vec<i32>> {
    raw_input
        .lines()
        .map(|x| {
            x.chars()
                .map(|y| y.to_digit(10).unwrap() as i32)
                .collect_vec()
        })
        .collect()
}

fn neighbors(r_size: usize, c_size: usize, r: usize, c: usize) -> Vec<(usize, usize)> {
    let mut nbrs = vec![];
    for (dr, dc) in [(0, 1), (0, -1), (1, 0), (-1, 0)] {
        let rr = r as i32 + dr;
        let cc = c as i32 + dc;

        if rr >= 0 && rr < r_size as i32 && cc >= 0 && cc < c_size as i32 {
            nbrs.push((rr as usize, cc as usize))
        }
    }

    nbrs
}

fn _p1(map: &[Vec<i32>]) -> i32 {
    let r_size = map.len();
    let c_size = map[0].len();
    let mut risk_level = 0;
    for r in 0..r_size {
        'outer: for c in 0..c_size {
            let risk = map[r][c];
            for (rr, cc) in neighbors(r_size, c_size, r, c) {
                if map[rr][cc] < risk {
                    continue 'outer;
                }
            }
            risk_level += risk + 1;
        }
    }
    risk_level
}

fn _p2(map: &[Vec<i32>]) -> usize {
    let r_size = map.len();
    let c_size = map[0].len();
    let mut visited = vec![vec![false; c_size]; r_size];
    let mut basins = vec![];
    for i in 0..r_size {
        for j in 0..c_size {
            if map[i][j] == 9 || visited[i][j] {
                continue;
            }

            let mut component_size = 0;
            let mut queue = vec![];
            queue.push((i, j));
            visited[i][j] = true;

            while let Some((r, c)) = queue.pop() {
                component_size += 1;

                for (rr, cc) in neighbors(r_size, c_size, r, c) {
                    if !visited[rr][cc] && map[rr][cc] != 9 {
                        queue.push((rr, cc));
                        visited[rr][cc] = true;
                    }
                }
            }
            basins.push(component_size);
        }
    }

    basins.iter().sorted().rev().take(3).product()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 15);
        assert_eq!(p2(raw_input()), 1134);
    }

    fn raw_input<'a>() -> &'a str {
        "2199943210
3987894921
9856789892
8767896789
9899965678"
    }
}
