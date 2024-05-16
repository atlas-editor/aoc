fn main() {
    let inp = include_str!("tmp.test")
        .lines()
        .map(|x| {
            x.chars()
                .map(|y| y.to_string().parse::<i32>().unwrap())
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>();
    println!("{}", p1(&inp));
    println!("{}", p2(&inp));
}

fn p1(map: &[Vec<i32>]) -> i32 {
    let R = map.len();
    let C = map[0].len();
    let mut res = 0;
    for i in 0..R {
        for j in 0..C {
            let mut nbrs = vec![];
            if i > 0 {
                nbrs.push(map[i - 1][j]);
            }
            if i < R - 1 {
                nbrs.push(map[i + 1][j]);
            }
            if j > 0 {
                nbrs.push(map[i][j - 1]);
            }
            if j < C - 1 {
                nbrs.push(map[i][j + 1]);
            }
            if map[i][j] < *nbrs.iter().min().unwrap() {
                res += map[i][j] + 1;
            }
        }
    }

    res
}

fn p2(map: &[Vec<i32>]) -> usize {
    0
}
