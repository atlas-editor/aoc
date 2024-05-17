fn main() {
    let input: Vec<Vec<u8>> = include_str!("tmp.in")
        .lines()
        .map(|x| {
            x.chars()
                .map(|y| y.to_digit(10).unwrap() as u8)
                .collect()
        })
        .collect();
    println!("{}", p(&input, false));
    println!("{}", p(&input, true));
}

fn p(input: &[Vec<u8>], p2: bool) -> i32 {
    let (R, C) = (input.len(), input[0].len());
    let mut energy = input.to_owned();
    let mut flashed = vec![vec![false; C]; R];
    let mut res = 0;
    for i in 0.. {
        if i == 100 && !p2 {
            return res;
        }
        let mut nines = vec![];
        for (r, row) in energy.iter_mut().enumerate().take(R) {
            for (c, item) in row.iter_mut().enumerate().take(C) {
                *item += 1;
                if *item > 9 {
                    nines.push((r, c));
                }
            }
        }

        for nine in nines {
            let mut q = vec![nine];
            while let Some((r, c)) = q.pop() {
                if flashed[r][c] {
                    continue;
                }
                flashed[r][c] = true;

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
                    let rr = r as i32 + dr;
                    let cc = c as i32 + dc;

                    if 0 <= rr && rr < R as i32 && 0 <= cc && cc < C as i32 {
                        let rr = rr as usize;
                        let cc = cc as usize;
                        energy[rr][cc] += 1;
                        if energy[rr][cc] > 9 && !flashed[rr][cc] {
                            q.push((rr, cc));
                        }
                    }
                }
            }
        }
        let mut t = 0;
        for r in 0..R {
            for c in 0..C {
                if flashed[r][c] {
                    energy[r][c] = 0;
                    t += 1;
                }
                flashed[r][c] = false;
            }
        }
        if t == R * C && p2 {
            return i + 1;
        }
        res += t as i32;
    }
    -1
}
