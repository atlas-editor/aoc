use crate::days::utils::{atoi, atopi, ByteGraph, ByteSet};
use crate::{bgraph, bset};
use bstr::ByteSlice;
use itertools::Itertools;
use maplit::hashset;
use std::collections::HashSet;

pub fn p1(raw_input: &[u8]) -> i32 {
    let scanners = parse_input(raw_input);
    _p1(scanners)
}

pub fn p2(raw_input: &[u8]) -> i16 {
    let scanners = parse_input(raw_input);
    _p2(scanners)
}

fn read_int(input: &[u8]) -> usize {
    for i in 0..input.len() {
        if input[i] < b'0' || input[i] > b'9' {
            return atopi(&input[..i]);
        }
    }
    atopi(input)
}
fn parse_scanner(scanner_data: &[u8]) -> Scanner {
    let (header, data) = scanner_data.split_once_str(b"\n").unwrap();
    let beacons = data
        .lines()
        .map(|line| {
            let pos = line.split_str(b",").collect_vec();
            Beacon::from_vec([atoi(pos[0]), atoi(pos[1]), atoi(pos[2])])
        })
        .collect_vec();

    Scanner::new(read_int(&header[12..14]), beacons)
}

fn parse_input(raw_input: &[u8]) -> Vec<Scanner> {
    raw_input
        .split_str(b"\n\n")
        .map(|scanner_data| parse_scanner(scanner_data))
        .collect_vec()
}

#[derive(Debug, Clone)]
struct Beacon {
    position: Vector3D,
}

impl Beacon {
    fn distance_hash_from(&self, other: &Self) -> u64 {
        let mut per_coor = [
            (self.position[0] - other.position[0]).abs() as u64,
            (self.position[1] - other.position[1]).abs() as u64,
            (self.position[2] - other.position[2]).abs() as u64,
        ]
        .into_iter()
        .sorted()
        .collect_vec();
        per_coor[0] + (per_coor[1] << 16) + (per_coor[2] << 32)
    }

    fn from_vec(vec: Vector3D) -> Self {
        Self { position: vec }
    }

    fn to_vec(&self) -> Vector3D {
        self.position
    }
}

#[derive(Debug, Clone)]
struct Scanner {
    id: usize,
    beacons: Vec<Beacon>,
    abs_position: Option<Vector3D>,
    transformation: Option<Transformation>,
}

impl Scanner {
    fn new(id: usize, beacons: Vec<Beacon>) -> Self {
        Self {
            id,
            beacons,
            abs_position: None,
            transformation: None,
        }
    }

    fn get_distances(&self) -> Vec<HashSet<u64>> {
        self.beacons
            .iter()
            .map(|b0| {
                self.beacons
                    .iter()
                    .map(|b1| b0.distance_hash_from(b1))
                    .collect()
            })
            .collect()
    }

    fn find_common_beacons(self, other: &Self, in_common: usize) -> Vec<(Beacon, Beacon)> {
        self.find_common_beacon_indices(other, in_common)
            .iter()
            .map(|&(i, j)| {
                (
                    self.clone().beacons_abs().unwrap()[i].clone(),
                    other.beacons.clone()[j].clone(),
                )
            })
            .collect_vec()
    }

    fn find_common_beacon_indices(&self, other: &Self, in_common: usize) -> Vec<(usize, usize)> {
        self.get_distances()
            .iter()
            .enumerate()
            .map(|(i, hs0)| {
                other
                    .get_distances()
                    .iter()
                    .enumerate()
                    .filter_map(|(j, hs1)| {
                        if hs0.intersection(hs1).count() >= in_common {
                            Some((i, j))
                        } else {
                            None
                        }
                    })
                    .collect_vec()
            })
            .flatten()
            .collect_vec()
    }

    fn are_overlapping(&self, other: &Self) -> bool {
        self.find_common_beacon_indices(other, 12).len() >= 12
    }

    fn is_determined(&self) -> bool {
        self.abs_position.is_some() && self.transformation.is_some()
    }

    fn beacons_abs(self) -> Option<Vec<Beacon>> {
        if !self.is_determined() {
            None
        } else {
            Some(
                self.beacons
                    .iter()
                    .map(|b| {
                        Beacon::from_vec(vec_add(
                            self.abs_position.unwrap(),
                            self.transformation.clone().unwrap().apply(b.to_vec()),
                        ))
                    })
                    .collect(),
            )
        }
    }
}

type Vector3D = [i16; 3];

fn vec_diff(v0: Vector3D, v1: Vector3D) -> Vector3D {
    [v0[0] - v1[0], v0[1] - v1[1], v0[2] - v1[2]]
}

fn vec_add(v0: Vector3D, v1: Vector3D) -> Vector3D {
    [v0[0] + v1[0], v0[1] + v1[1], v0[2] + v1[2]]
}

fn vec_manhattan_dist(v0: Vector3D, v1: Vector3D) -> i16 {
    (v0[0] - v1[0]).abs() + (v0[1] - v1[1]).abs() + (v0[2] - v1[2]).abs()
}

#[derive(Debug, Clone)]
struct Transformation {
    matrix: [i16; 9],
}

impl Transformation {
    fn new(matrix: [i16; 9]) -> Self {
        Transformation { matrix }
    }

    fn from_vecs(vecs: Vec<&[i16; 3]>) -> Self {
        Self::new([
            vecs[0][0], vecs[0][1], vecs[0][2], vecs[1][0], vecs[1][1], vecs[1][2], vecs[2][0],
            vecs[2][1], vecs[2][2],
        ])
    }

    fn identity() -> Self {
        Self::new([1, 0, 0, 0, 1, 0, 0, 0, 1])
    }

    fn apply(&self, vec: Vector3D) -> Vector3D {
        [
            self.matrix[0] * vec[0] + self.matrix[1] * vec[1] + self.matrix[2] * vec[2],
            self.matrix[3] * vec[0] + self.matrix[4] * vec[1] + self.matrix[5] * vec[2],
            self.matrix[6] * vec[0] + self.matrix[7] * vec[1] + self.matrix[8] * vec[2],
        ]
    }

    fn determinant(&self) -> i16 {
        self.matrix[0] * (self.matrix[4] * self.matrix[8] - self.matrix[5] * self.matrix[7])
            - self.matrix[1] * (self.matrix[3] * self.matrix[8] - self.matrix[5] * self.matrix[6])
            + self.matrix[2] * (self.matrix[3] * self.matrix[7] - self.matrix[4] * self.matrix[6])
    }

    fn generate_all() -> impl Iterator<Item = Self> {
        let x = [[1, 0, 0], [-1, 0, 0]];
        let y = [[0, 1, 0], [0, -1, 0]];
        let z = [[0, 0, 1], [0, 0, -1]];

        (0..2).flat_map(move |i| {
            (0..2).flat_map(move |j| {
                (0..2).flat_map(move |k| {
                    [x[i], y[j], z[k]]
                        .iter()
                        .permutations(3)
                        .map(|p| Self::from_vecs(p))
                        .filter_map(|s| if s.determinant() == 1 { Some(s) } else { None })
                        .collect_vec()
                })
            })
        })
    }
}

fn determine_scanner_position_by_transformation(
    beacon_abs: &Beacon,
    beacon_alt: &Beacon,
    transformation: &Transformation,
) -> Vector3D {
    vec_diff(
        beacon_abs.to_vec(),
        transformation.apply(beacon_alt.to_vec()),
    )
}

fn determine_abs_scanner_position_and_transformation(
    determined_scanner: Scanner,
    unknown_scanner: &Scanner,
) -> (Vector3D, Transformation) {
    let common_beacons = determined_scanner
        .clone()
        .find_common_beacons(unknown_scanner, 12);
    if common_beacons.len() < 12 {
        panic!(
            "scanner {} and scanner {} do not share 12 beacons",
            unknown_scanner.id, determined_scanner.id
        )
    }
    'outer: for t in Transformation::generate_all() {
        let pos = determine_scanner_position_by_transformation(
            &common_beacons[0].0,
            &common_beacons[0].1,
            &t,
        );
        for (determined, unknown) in &common_beacons[1..] {
            if determine_scanner_position_by_transformation(determined, unknown, &t) != pos {
                continue 'outer;
            }
        }
        return (pos, t);
    }
    panic!(
        "scanner {}'s position could not be determined using scanner {}",
        unknown_scanner.id, determined_scanner.id
    )
}

fn _p1(scanners: Vec<Scanner>) -> i32 {
    let mut duplicated_pairs = HashSet::new();
    let mut max_beacons = 0;
    let mut no_duplicates = 0;
    for i in 0..scanners.len() {
        max_beacons += scanners[i].beacons.len() as i32;
        for j in i + 1..scanners.len() {
            for (k, l) in scanners[i].find_common_beacon_indices(&scanners[j], 3) {
                if duplicated_pairs
                    .intersection(&hashset! {(i, k), (j, l)})
                    .count()
                    <= 1
                {
                    duplicated_pairs.extend([(i, k), (j, l)]);
                    no_duplicates += 1;
                }
            }
        }
    }

    max_beacons - no_duplicates
}

fn _p2(mut scanners: Vec<Scanner>) -> i16 {
    let mut g = bgraph! {};

    // scanner 0 will act as the origin whose orientation is the identity
    scanners[0].abs_position = Some([0, 0, 0]);
    scanners[0].transformation = Some(Transformation::identity());

    let no_of_scanners = scanners.len();

    for i in 0..no_of_scanners {
        let s0 = &scanners[i];
        for j in 0..no_of_scanners {
            let s1 = &scanners[j];
            if s0.are_overlapping(s1) {
                g.insert(i as u8, j as u8);
            }
        }
    }

    for (i, j) in g.visit_seq() {
        let s0 = scanners[i].clone();
        let mut s1 = scanners[j].clone();
        let (pos, t) = determine_abs_scanner_position_and_transformation(s0, &s1);
        s1.abs_position = Some(pos);
        s1.transformation = Some(t);
        scanners[j] = s1;
    }

    (0..no_of_scanners)
        .map(|i| {
            (0..no_of_scanners)
                .map(|j| {
                    vec_manhattan_dist(
                        scanners[i].abs_position.unwrap(),
                        scanners[j].abs_position.unwrap(),
                    )
                })
                .max()
                .unwrap()
        })
        .max()
        .unwrap()
}

impl ByteGraph {
    fn visit_seq(&self) -> impl Iterator<Item = (usize, usize)> {
        let mut seq = vec![];
        let mut visited = bset! {0};
        let mut queue = vec![0];

        while let Some(u) = queue.pop() {
            for &v in self.neighbors(u) {
                if !visited.contains(v) {
                    visited.insert(v);
                    queue.push(v);
                    seq.push((u as usize, v as usize))
                }
            }
        }
        seq.into_iter()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 79);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), 3621);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b"--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14"
    }
}
