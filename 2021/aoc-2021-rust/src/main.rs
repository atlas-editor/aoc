use std::fs;

mod days;
mod utils;

macro_rules! run {
    ($day:ident, $part:ident) => {{
        let input_path = format!("inputs/{}.in", &stringify!($day).to_string()[1..]);
        let raw_input = fs::read_to_string(input_path).unwrap();

        let start = std::time::SystemTime::now();
        let result = days::$day::$part(&raw_input);
        let duration = start.elapsed().unwrap();

        println!(
            "{} {} took {:.3} ms",
            stringify!($day),
            stringify!($part),
            duration.as_secs_f64() * 1000.0
        );
        println!("{}", result);
    }};
}

fn run_all() {
    // day1
    run!(d01, p1);
    run!(d01, p2);

    // day2
    run!(d02, p1);
    run!(d02, p2);

    // day3
    run!(d03, p1);
    run!(d03, p2);

    // day4
    run!(d04, p1);
    run!(d04, p2);

    // day5
    run!(d05, p1);
    run!(d05, p2);

    // day6
    run!(d06, p1);
    run!(d06, p2);

    // day7
    run!(d07, p1);
    run!(d07, p2);

    // day8
    run!(d08, p1);
    run!(d08, p2);

    // day9
    run!(d09, p1);
    run!(d09, p2);

    // day10
    run!(d10, p1);
    run!(d10, p2);

    // day11
    run!(d11, p1);
    run!(d11, p2);

    // day12
    run!(d12, p1);
    run!(d12, p2);

    // day13
    run!(d13, p1);
    run!(d13, p2);

    // day14
    run!(d14, p1);
    run!(d14, p2);
}

fn main() {
    run_all()
}
