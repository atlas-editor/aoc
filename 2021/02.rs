fn main() {
    let inp = include_str!("inputs/02.in").lines().collect::<Vec<_>>();
    println!("{:#?}", p1(&inp));
    println!("{:#?}", p2(&inp));
}

fn p1(x: &Vec<&str>) -> i32 {
    let (mut h, mut d) = (0, 0);
    for e in x {
        let y = e.split_whitespace().collect::<Vec<_>>();
        let dir = y[0];
        let units = y[1].parse::<i32>().unwrap();
        if dir == "forward" {
            h += units
        } else if dir == "down" {
            d += units
        } else {
            d -= units
        }
    }
    return h * d;
}

fn p2(x: &Vec<&str>) -> i32 {
    let (mut h, mut d, mut aim) = (0, 0, 0);
    for e in x {
        let y = e.split_whitespace().collect::<Vec<_>>();
        let dir = y[0];
        let units = y[1].parse::<i32>().unwrap();
        if dir == "forward" {
            h += units;
            d += aim * units
        } else if dir == "down" {
            aim += units;
        } else {
            aim -= units;
        }
    }
    return h * d;
}
