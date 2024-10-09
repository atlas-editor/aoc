use ahash::{HashMap, HashMapExt};
use bstr::ByteSlice;
use std::collections::BinaryHeap;

/*
0 == A
1 == B
2 == C
3 == D
4 == .

[0, 11) hallway
[11, 15) room level 1
[15, 19) room level 2
[19, 23) room level 3
[23, 27) room level 4
 */

const ENERGY: [i32; 4] = [1, 10, 100, 1000];
const HOME: [u8; 4] = [11, 12, 13, 14];
const ROOM_TO_HALLWAY: [u8; 15] = {
    let mut arr = [0; 15];
    arr[11] = 2;
    arr[12] = 4;
    arr[13] = 6;
    arr[14] = 8;
    arr
};

const fn goal<const N: usize>() -> [u8; N] {
    let mut end = [4; N];
    let mut i = 11;
    while i < N {
        end[i] = ((i - 11) % 4) as u8;
        i += 1;
    }
    end
}

pub fn p1(raw_input: &[u8]) -> i32 {
    let mut start = [4; 19];
    for i in 11..19 {
        if i <= 14 {
            start[i] = raw_input[31 + 2 * (i - 11)] - 65
        } else {
            start[i] = raw_input[45 + 2 * (i - 15)] - 65
        }
    }

    dijkstra(start, goal())
}

pub fn p2(raw_input: &[u8]) -> i32 {
    let mut start = [4; 27];
    for i in 11..27 {
        if i <= 14 {
            start[i] = raw_input[31 + 2 * (i - 11)] - 65;
        } else if i <= 18 {
            start[i] = [3, 2, 1, 0][(i - 15) % 4];
        } else if i <= 22 {
            start[i] = [3, 1, 0, 2][(i - 19) % 4];
        } else {
            start[i] = raw_input[45 + 2 * (i - 23)] - 65;
        }
    }

    dijkstra(start, goal())
}

fn dijkstra<const N: usize>(start: [u8; N], end: [u8; N]) -> i32 {
    let mut dist = HashMap::new();
    let mut heap = BinaryHeap::new();

    dist.insert(start, 0);
    heap.push((0, start));
    while let Some((cu, u)) = heap.pop() {
        if u == end {
            return -cu;
        }

        for (cv, v) in diagram_neighbors(u) {
            let alt = -cu + cv;
            if alt < *dist.get(&v).unwrap_or(&i32::MAX) {
                dist.insert(v, alt);
                heap.push((-alt, v));
            }
        }
    }
    panic!("destination not reached")
}

fn diagram_neighbors<const N: usize>(diagram: [u8; N]) -> Vec<(i32, [u8; N])> {
    let mut neighbors = vec![];
    for i in 0..N {
        let a = diagram[i];
        if a == 4 {
            continue;
        }
        let curr_home = HOME[a as usize];
        if i < 11 {
            // in the hallway
            let (reachable_entries, _) = explore_hallway(i as u8, &diagram);

            // we can only go home
            if reachable_entries.contains(&ROOM_TO_HALLWAY[curr_home as usize]) {
                if let Some(single_neighbor) = go_home_from_hallway(i as u8, i as u8, 0, &diagram) {
                    return vec![single_neighbor];
                }
            }
        } else {
            // in a room
            let curr_hallway = ROOM_TO_HALLWAY[(i - 11) % 4 + 11];

            // at home and below only same amphipods OR blocked from leaving room
            if (i.abs_diff(curr_home as usize) % 4 == 0
                && (i + 4..N).step_by(4).all(|x| diagram[x] == a))
                || (11..=(i - 4)).rev().step_by(4).any(|x| diagram[x] != 4)
                || (diagram[curr_hallway as usize - 1] != 4
                    && diagram[curr_hallway as usize + 1] != 4)
            {
                continue;
            }

            let steps_to_hallway = (i - 11) / 4 + 1;

            // we must and can step out
            let (reachable_entries, reachable_hallway) = explore_hallway(curr_hallway, &diagram);

            // if we can go home we should
            if reachable_entries.contains(&ROOM_TO_HALLWAY[curr_home as usize]) {
                if let Some(single_neighbor) =
                    go_home_from_hallway(i as u8, curr_hallway, steps_to_hallway, &diagram)
                {
                    return vec![single_neighbor];
                }
            }

            // we cannot go home so we go to each reachable spot in the hallway
            for j in reachable_hallway {
                let steps_in_hallway = curr_hallway.abs_diff(j) as usize;
                let mut neighbor = diagram;
                neighbor[j as usize] = a;
                neighbor[i] = 4;

                neighbors.push((
                    ENERGY[a as usize] * (steps_to_hallway + steps_in_hallway) as i32,
                    neighbor,
                ));
            }
        }
    }

    neighbors
}

fn explore_hallway<const N: usize>(start: u8, diagram: &[u8; N]) -> (Vec<u8>, Vec<u8>) {
    let mut reachable_entries = vec![];
    let mut hallway_left = vec![];
    for j in (0..=start.saturating_sub(1)).rev() {
        if diagram[j as usize] != 4 {
            break;
        }
        if ![2, 4, 6, 8].contains(&j) {
            hallway_left.push(j);
        } else {
            reachable_entries.push(j);
        }
    }

    let mut hallway_right = vec![];
    for j in start + 1..11 {
        if diagram[j as usize] != 4 {
            break;
        }
        if ![2, 4, 6, 8].contains(&j) {
            hallway_right.push(j);
        } else {
            reachable_entries.push(j);
        }
    }

    (reachable_entries, [hallway_left, hallway_right].concat())
}

fn go_home_from_hallway<const N: usize>(
    pos_actual: u8,
    pos_hallway: u8,
    steps_to_hallway: usize,
    diagram: &[u8; N],
) -> Option<(i32, [u8; N])> {
    let amphipod = diagram[pos_actual as usize];
    let amphipod_home = HOME[amphipod as usize];
    for j in (amphipod_home as usize..N).step_by(4).rev() {
        if diagram[j] != 4 && diagram[j] != amphipod {
            break;
        }
        if diagram[j] == 4 {
            let mut single_neighbor = diagram.clone();
            single_neighbor[j] = amphipod;
            single_neighbor[pos_actual as usize] = 4;

            let steps_in_hallway =
                pos_hallway.abs_diff(ROOM_TO_HALLWAY[amphipod_home as usize]) as usize;
            let steps_from_hallway = (j - 11) / 4 + 1;

            return Some((
                ENERGY[amphipod as usize]
                    * (steps_to_hallway + steps_in_hallway + steps_from_hallway) as i32,
                single_neighbor,
            ));
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 12521);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), 44169);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########"
    }
}
