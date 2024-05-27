use regex::Regex;
use std::fmt::Debug;
use std::str::FromStr;

pub fn ints<T: FromStr>(input: &str) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    let re = Regex::new(r"-?\d+").unwrap();
    re.find_iter(input)
        .map(|m| m.as_str().parse().unwrap())
        .collect()
}
