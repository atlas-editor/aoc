use ahash::{HashMap, HashMapExt};
use bstr::ByteSlice;
use itertools::iproduct;
use std::ops::{Add, AddAssign};

pub fn p1(raw_input: &[u8]) -> u32 {
    let (a, b) = raw_input.split_once_str("\n").unwrap();
    let player_a = Player::from_input(a);
    let player_b = Player::from_input(b);
    simulate_simple_game(player_a, player_b)
}

pub fn p2(raw_input: &[u8]) -> u64 {
    let (a, b) = raw_input.split_once_str("\n").unwrap();
    let mut cache = HashMap::new();
    simulate_quantum_game(a[a.len() - 1] - 48, 0, b[b.len() - 1] - 48, 0, &mut cache).max()
}

struct Player {
    position: u8,
    score: u32,
}

impl Player {
    fn from_input(input: &[u8]) -> Self {
        Player {
            position: input[input.len() - 1] - 48,
            score: 0,
        }
    }

    fn r#move(&mut self, roll: u16) {
        self.position = (((self.position as u16 + roll - 1) % 10) + 1) as u8;
        self.score += self.position as u32;
    }
}

struct Dice {
    state: (u8, u8, u8),
    rolls: u32,
}

impl Dice {
    fn new() -> Self {
        Self {
            state: (1, 2, 3),
            rolls: 0,
        }
    }

    fn roll(&mut self) -> u16 {
        let curr = self.state.0 as u16 + self.state.1 as u16 + self.state.2 as u16;
        self.state = (
            ((self.state.0 + 2) % 100) + 1,
            ((self.state.1 + 2) % 100) + 1,
            ((self.state.2 + 2) % 100) + 1,
        );
        self.rolls += 3;
        curr
    }
}

fn simulate_simple_game(mut player_a: Player, mut player_b: Player) -> u32 {
    let mut dice = Dice::new();
    loop {
        player_a.r#move(dice.roll());
        if player_a.score >= 1000 {
            return player_b.score * dice.rolls;
        }

        player_b.r#move(dice.roll());
        if player_b.score >= 1000 {
            return player_a.score * dice.rolls;
        }
    }
}

type QuantumState = (u8, u64, u8, u64);

#[derive(Clone, Copy)]
struct Score {
    a: u64,
    b: u64,
}

impl Score {
    fn max(&self) -> u64 {
        if self.a > self.b {
            self.a
        } else {
            self.b
        }
    }

    fn flip(self) -> Self {
        Score {
            a: self.b,
            b: self.a,
        }
    }
}

impl AddAssign<Score> for Score {
    fn add_assign(&mut self, rhs: Score) {
        self.a += rhs.a;
        self.b += rhs.b;
    }
}

fn simulate_quantum_game(
    position_a: u8,
    score_a: u64,
    position_b: u8,
    score_b: u64,
    mut cache: &mut HashMap<QuantumState, Score>,
) -> Score {
    let state: QuantumState = (position_a, score_a, position_b, score_b);
    if let Some(k) = cache.get(&state) {
        return *k;
    }
    if score_a >= 21 {
        return Score { a: 1, b: 0 };
    }
    if score_b >= 21 {
        return Score { a: 0, b: 1 };
    }

    let mut score = Score { a: 0, b: 0 };
    for i in 1..4 {
        for j in 1..4 {
            for k in 1..4 {
                let new_position_a = (position_a + (i + j + k) - 1) % 10 + 1;
                score += simulate_quantum_game(
                    position_b,
                    score_b,
                    new_position_a,
                    score_a + new_position_a as u64,
                    cache,
                )
                .flip();
            }
        }
    }

    cache.insert(state, score);
    score
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 739785);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), 444356092776315);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"Player 1 starting position: 4
Player 2 starting position: 8"
    }
}
