use crate::days::utils::Matrix;

pub fn p1(raw_input: &[u8]) -> i32 {
    run(raw_input)
}

pub fn p2(_raw_input: &[u8]) {}

fn run(raw_input: &[u8]) -> i32 {
    let mut sea = Matrix::bytes(raw_input);
    let mut sea_copy: Matrix<u8>;

    let rows = sea.r_size();
    let cols = sea.c_size();

    for i in 0.. {
        sea_copy = sea.clone();
        let mut moved = false;

        for r in 0..rows {
            for c in 0..cols {
                if sea[(r, c)] == b'>' && sea[(r, (c + 1) % cols)] == b'.' {
                    moved = true;
                    sea_copy[(r, c)] = b'.';
                    sea_copy[(r, (c + 1) % cols)] = b'>';
                }
            }
        }

        sea = sea_copy.clone();

        for r in 0..rows {
            for c in 0..cols {
                if sea_copy[(r, c)] == b'v' && sea_copy[((r + 1) % rows, c)] == b'.' {
                    moved = true;
                    sea[(r, c)] = b'.';
                    sea[((r + 1) % rows, c)] = b'v';
                }
            }
        }

        if !moved {
            return i + 1;
        }
    }

    unreachable!()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 58);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), ());
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>"
    }
}
