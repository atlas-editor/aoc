pub fn p1(raw_input: &str) -> usize {
    let input = parse_input(raw_input);
    _p1(&input)
}

pub fn p2(raw_input: &str) -> u32 {
    let input = parse_input(raw_input);
    _p2(&input)
}

fn parse_input(raw_input: &str) -> Vec<&str> {
    raw_input.lines().collect()
}

fn _p1(entries: &[&str]) -> usize {
    entries
        .iter()
        .map(|e| {
            e.split('|')
                .last()
                .unwrap()
                .split_ascii_whitespace()
                .filter(|x| [2, 4, 3, 7].contains(&x.len()))
                .count()
        })
        .sum()
}

fn set(digit: &str) -> u8 {
    let letters = "abcdefgh";
    digit
        .chars()
        .map(|c| letters.find(c).unwrap())
        .fold(0, |result, index| result | (1 << index))
}

fn find_pop<T, F: Fn(&T) -> bool>(vec: &mut Vec<T>, predicate: F) -> T {
    vec.iter()
        .position(predicate)
        .map(|index| vec.remove(index))
        .unwrap()
}

/// Checks if `x` is subset of `y`.
fn is_subset(x: u8, y: u8) -> bool {
    (x & y) == x
}

fn decode(codes_raw: &str, output: &str) -> u32 {
    let mut codes = codes_raw.split_ascii_whitespace().map(set).collect();

    let one = find_pop(&mut codes, |x| x.count_ones() == 2);
    let four = find_pop(&mut codes, |x| x.count_ones() == 4);
    let seven = find_pop(&mut codes, |x| x.count_ones() == 3);
    let eight = find_pop(&mut codes, |x| x.count_ones() == 7);
    let nine = find_pop(&mut codes, |x| x.count_ones() == 6 && is_subset(four, *x));
    let zero = find_pop(&mut codes, |x| x.count_ones() == 6 && is_subset(one, *x));
    let six = find_pop(&mut codes, |x| x.count_ones() == 6);
    let three = find_pop(&mut codes, |x| x.count_ones() == 5 && is_subset(seven, *x));
    let five = find_pop(&mut codes, |x| x.count_ones() == 5 && is_subset(*x, nine));
    let two = codes.pop().unwrap();

    let decoded = [zero, one, two, three, four, five, six, seven, eight, nine];

    output
        .split_ascii_whitespace()
        .map(|x| {
            decoded
                .iter()
                .position(|y| *y == set(x))
                .unwrap()
                .to_string()
        })
        .collect::<String>()
        .parse()
        .unwrap()
}

fn _p2(entries: &[&str]) -> u32 {
    entries
        .iter()
        .map(|e| {
            let (codes_raw, output) = e.split_once('|').unwrap();
            decode(codes_raw, output)
        })
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(p1(raw_input()), 26);
        assert_eq!(p2(raw_input()), 61229);
    }

    fn raw_input<'a>() -> &'a str {
        "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"
    }
}
