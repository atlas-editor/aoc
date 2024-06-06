use crate::utils::{benchmark_run, print_day, print_header};
use std::fs;
mod days;
mod utils;

macro_rules! benchmark_all {
    ($($day:ident),*) => {{
        print_header();
        $(
        let input_path = format!("inputs/{}.in", &stringify!($day).to_string()[1..]);
        let raw_input = fs::read_to_string(input_path).unwrap();

        let p1_duration = benchmark_run(days::$day::p1, &raw_input);
        let p2_duration = benchmark_run(days::$day::p2, &raw_input);

        print_day(stringify!($day).to_string()[1..].parse().unwrap(), p1_duration, p2_duration);
        )*
    }};
}

fn main() {
    benchmark_all!(d01, d02, d03, d04, d05, d06, d07, d08, d09, d10, d11, d12, d13, d14, d15);
}
