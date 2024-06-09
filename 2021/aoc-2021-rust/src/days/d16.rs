use itertools::Itertools;
use std::any::Any;
use std::usize;

pub fn p1(raw_input: &str) -> i32 {
    let code = parse_input(raw_input);
    let packet = parse_packets(&code);
    version_sum(packet)
}

pub fn p2(raw_input: &str) -> u64 {
    let code = parse_input(raw_input);
    let packet = parse_packets(&code);
    value(&packet)
}

fn parse_input(raw_input: &str) -> String {
    raw_input
        .chars()
        .map(|x| {
            let y = x.to_digit(16).unwrap();
            format!("{:04b}", y)
        })
        .collect::<String>()
}

#[derive(Debug, Clone)]
struct Header {
    version: u8,
    type_id: u8,
}

#[derive(Debug, Clone)]
struct Packet {
    header: Header,
    length_type_id: Option<u8>,
    literal_values: Vec<String>,
    sub_packets: Vec<Box<Packet>>,
}

impl Packet {
    fn literal_value(&self) -> Option<u64> {
        if self.literal_values.is_empty() {
            None
        } else {
            Some(u64::from_str_radix(self.literal_values.join("").as_str(), 2).unwrap())
        }
    }
}

fn parse_header(code: &str) -> Header {
    Header {
        version: u8::from_str_radix(&code[0..3], 2).unwrap(),
        type_id: u8::from_str_radix(&code[3..6], 2).unwrap(),
    }
}

fn parse_length_type_id(code: &str) -> u8 {
    u8::from_str_radix(&code[6..7], 2).unwrap()
}

fn parse_literal_values(code: &str) -> (Vec<String>, usize) {
    let mut res = vec![];
    let mut i = 6;
    while i < code.len() {
        if code.chars().nth(i).unwrap() == '1' {
            res.push(&code[i + 1..i + 5]);
            i += 5;
        } else {
            res.push(&code[i + 1..i + 5]);
            i += 5;
            break;
        }
    }
    (
        res.iter().map(|x| x.to_string()).collect(),
        6 + res.len() * 5,
    )
}

fn parse_packets(code: &str) -> Box<Packet> {
    _parse_packets(code, None).0[0].clone()
}

fn _parse_packets(code: &str, sub_packets: Option<usize>) -> (Vec<Box<Packet>>, usize) {
    let mut packets = vec![];
    let mut i = 0;

    while i + 11 <= code.len() {
        match sub_packets {
            Some(n) if n == packets.len() => return (packets, i),
            _ => {}
        }
        let header = parse_header(&code[i..]);
        if header.type_id == 4 {
            let (literal_values, end) = parse_literal_values(&code[i..]);
            packets.push(Box::new(Packet {
                header,
                length_type_id: None,
                literal_values,
                sub_packets: vec![],
            }));
            i += end;
        } else {
            let length_type_id = parse_length_type_id(&code[i..]);
            if length_type_id == 0 {
                let total_length = usize::from_str_radix(&code[i + 7..i + 22], 2).unwrap();
                let (sub_packets, _) = _parse_packets(&code[i + 22..i + 22 + total_length], None);
                packets.push(Box::new(Packet {
                    header,
                    length_type_id: Some(length_type_id),
                    literal_values: vec![],
                    sub_packets,
                }));
                i += 22 + total_length;
            } else {
                let no_sub_packets = usize::from_str_radix(&code[i + 7..i + 18], 2).unwrap();
                let (sub_packets, end) = _parse_packets(&code[i + 18..], Some(no_sub_packets));
                packets.push(Box::new(Packet {
                    header,
                    length_type_id: Some(length_type_id),
                    literal_values: vec![],
                    sub_packets,
                }));
                i += 18 + end;
            }
        }
    }
    (packets, i)
}

fn version_sum(packet: Box<Packet>) -> i32 {
    let mut s = 0;
    let mut q = vec![&packet];
    while !q.is_empty() {
        let curr = q.pop().unwrap();
        s += curr.header.version as i32;

        for p in &curr.sub_packets {
            q.push(p);
        }
    }
    s
}

fn value(packet: &Box<Packet>) -> u64 {
    match packet.header.type_id {
        0 => packet.sub_packets.iter().map(|x| value(x)).sum(),
        1 => packet.sub_packets.iter().map(|x| value(x)).product(),
        2 => packet.sub_packets.iter().map(|x| value(x)).min().unwrap(),
        3 => packet.sub_packets.iter().map(|x| value(x)).max().unwrap(),
        4 => packet.literal_value().unwrap(),
        5 => (value(&packet.sub_packets[0]) > value(&packet.sub_packets[1])) as u64,
        6 => (value(&packet.sub_packets[0]) < value(&packet.sub_packets[1])) as u64,
        7 => (value(&packet.sub_packets[0]) == value(&packet.sub_packets[1])) as u64,
        _ => panic!("invalid type_id"),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut raw_input = "D2FE28";
        assert_eq!(p1(raw_input), 6);
        raw_input = "38006F45291200";
        assert_eq!(p1(raw_input), 9);
        raw_input = "EE00D40C823060";
        assert_eq!(p1(raw_input), 14);
        raw_input = "8A004A801A8002F478";
        assert_eq!(p1(raw_input), 16);
        raw_input = "620080001611562C8802118E34";
        assert_eq!(p1(raw_input), 12);
        raw_input = "C0015000016115A2E0802F182340";
        assert_eq!(p1(raw_input), 23);
        raw_input = "A0016C880162017C3686B18A3D4780";
        assert_eq!(p1(raw_input), 31);

        raw_input = "C200B40A82";
        assert_eq!(p2(raw_input), 3);
        raw_input = "04005AC33890";
        assert_eq!(p2(raw_input), 54);
        raw_input = "880086C3E88112";
        assert_eq!(p2(raw_input), 7);
        raw_input = "CE00C43D881120";
        assert_eq!(p2(raw_input), 9);
        raw_input = "D8005AC2A8F0";
        assert_eq!(p2(raw_input), 1);
        raw_input = "F600BC2D8F";
        assert_eq!(p2(raw_input), 0);
        raw_input = "9C005AC2F8F0";
        assert_eq!(p2(raw_input), 0);
        raw_input = "9C0141080250320F1802104A08";
        assert_eq!(p2(raw_input), 1);
    }
}
