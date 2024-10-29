use bstr::ByteSlice;
use itertools::Itertools;
use std::cmp::{max, min};

pub fn p1(raw_input: &[u8]) -> i64 {
    let restrictions = interpret_instructions(raw_input);
    let minmax = get_min_max(restrictions);
    minmax.1
}

pub fn p2(raw_input: &[u8]) -> i64 {
    let restrictions = interpret_instructions(raw_input);
    let minmax = get_min_max(restrictions);
    minmax.0
}

fn parse_second_argument(second: Option<Vec<u8>>) -> Option<Expression> {
    if [b'w', b'x', b'y', b'z'].contains(&second.clone()?[0]) {
        Some(Expression::new_var(second?[0]))
    } else {
        Some(Expression::new_const(
            String::from_utf8(second?).unwrap().parse().unwrap(),
        ))
    }
}

fn interpret_instructions(raw_instructions: &[u8]) -> Vec<Expression> {
    let mut vars = vec![Expression::new_const(0); (b'z' + 1) as usize];
    for v in b'w'..b'z' + 1 {
        vars[v as usize] = Expression::new_const(0);
    }

    let mut restrictions = vec![];
    let mut inp_no = 0;

    for line in raw_instructions.lines() {
        let fields = line.fields().collect_vec();
        let op = fields[0];
        let mut second = None;
        if fields.len() == 3 {
            second = Some(fields[2].to_vec());
        }
        let arg0 = fields[1][0];
        let lhs = vars[arg0 as usize].clone();

        let arg1 = parse_second_argument(second);
        let mut expr = match op {
            b"inp" => {
                inp_no += 1;
                Expression::new_var(inp_no)
            }
            bin => {
                let rhs = match arg1 {
                    Some(Expression { op: Op::Var(v), .. }) => vars[v as usize].clone(),
                    Some(Expression {
                        op: Op::Const(_), ..
                    }) => arg1.unwrap(),
                    _ => panic!("!instruction: rhs must be var or const"),
                };
                match bin {
                    b"mul" => Expression::new_mul(lhs, rhs),
                    b"add" => Expression::new_add(lhs, rhs),
                    b"mod" => Expression::new_mod(lhs, rhs),
                    b"div" => Expression::new_div(lhs, rhs),
                    b"eql" => Expression::new_eql(lhs, rhs),
                    _ => panic!("!instruction: unknown op"),
                }
            }
        };

        expr = simplify_expression(expr);

        // instructions of this type must be forced to achieve z == 0 at the end
        if line == b"eql x w" {
            match expr.op {
                Op::Const(_) => {
                    // in some cases this already evaluates to a constant
                }
                _ => {
                    restrictions.push(expr);
                    expr = Expression::new_const(1);
                }
            }
        }

        vars[arg0 as usize] = expr;
    }

    restrictions
}

fn get_min_max(restrictions: Vec<Expression>) -> (i64, i64) {
    let mut mm = vec![(0, 0); 14];
    for r in restrictions {
        if let Op::Eql((lhs, rhs)) = r.clone().op {
            if let Op::Add((inner_lhs, inner_rhs)) = lhs.op {
                if let Op::Var(v_lhs) = inner_lhs.op {
                    if let Op::Const(c_lhs) = inner_rhs.op {
                        if let Op::Var(v_rhs) = rhs.op {
                            // u, v -- variables
                            // C -- constant
                            // u + C == v
                            mm[v_lhs as usize - 1] = (max(1 - c_lhs, 1), min(9 - c_lhs, 9));
                            mm[v_rhs as usize - 1] = (
                                mm[v_lhs as usize - 1].0 + c_lhs,
                                mm[v_lhs as usize - 1].1 + c_lhs,
                            );
                            continue;
                        }
                    }
                }
            }
        }

        panic!("!restriction: incorrect format {:?}", r.clone());
    }
    let min_val: i64 = mm
        .iter()
        .map(|&(first, _)| first.to_string())
        .collect::<String>()
        .parse()
        .unwrap();
    let max_val: i64 = mm
        .iter()
        .map(|&(_, second)| second.to_string())
        .collect::<String>()
        .parse()
        .unwrap();

    (min_val, max_val)
}

fn simplify_expression(e: Expression) -> Expression {
    match e.clone().op {
        // e, f -- expressions
        // C, D -- constants
        Op::Mul((lhs, rhs)) => {
            if lhs.op == Op::Const(0) || rhs.op == Op::Const(0) {
                // 0 * e == 0
                // e * 0 == 0
                return Expression::new_const(0);
            }

            if rhs.op == Op::Const(1) {
                // e * 1 == e
                return *lhs;
            }

            if let Op::Const(c_lhs) = (*lhs).op {
                if let Op::Const(c_rhs) = (*rhs).op {
                    // C * D
                    return Expression::new_const(c_lhs * c_rhs);
                }
            }

            e
        }
        Op::Add((lhs, rhs)) => {
            if lhs.op == Op::Const(0) {
                // e + 0 == e
                return *rhs;
            }

            if let Op::Const(c_rhs) = rhs.op {
                if let Op::Const(c_lhs) = lhs.op {
                    // C + D
                    return Expression::new_const(c_lhs + c_rhs);
                }

                if let Op::Add((inner_lhs, inner_rhs)) = lhs.op {
                    if let Op::Const(c_inner_rhs) = inner_rhs.op {
                        // (e + C) + D == e + (C+D)
                        return simplify_expression(Expression::new_add(
                            *inner_lhs,
                            Expression::new_const(c_rhs + c_inner_rhs),
                        ));
                    }
                }
            }

            e
        }
        Op::Mod((lhs, rhs)) => {
            if let Op::Const(m) = rhs.op {
                match lhs.clone().op {
                    Op::Mul((_, inner_rhs)) => {
                        if inner_rhs.op == Op::Const(m) {
                            // (e * C) % C == 0
                            return Expression::new_const(0);
                        }
                    }
                    Op::Add((inner_lhs, inner_rhs)) => {
                        // (e + f) % C == (e % C) + (f % C)
                        return simplify_expression(Expression::new_add(
                            simplify_expression(Expression::new_mod(*inner_lhs, *rhs.clone())),
                            simplify_expression(Expression::new_mod(*inner_rhs, *rhs)),
                        ));
                    }
                    Op::Const(c) => {
                        // C % D
                        return Expression::new_const(c % m);
                    }
                    _ => {}
                }
                if lhs.max < m {
                    // e % C == e if e < C
                    return *lhs;
                }
            }

            e
        }
        Op::Div((lhs, rhs)) => {
            if rhs.op == Op::Const(1) {
                // e / 1 == e
                return *lhs;
            }

            if let Op::Const(c_rhs) = rhs.op {
                if let Op::Const(c_lhs) = lhs.op {
                    // C / D
                    return Expression::new_const(c_lhs / c_rhs);
                }

                match lhs.op {
                    Op::Mul((inner_lhs, inner_rhs)) => {
                        if inner_rhs.op == rhs.op {
                            // (e * C) / C == e
                            return *inner_lhs;
                        }
                    }
                    Op::Add((inner_lhs, inner_rhs)) => {
                        if inner_rhs.max < c_rhs {
                            // (e + f) / C == e / C if f < C
                            return simplify_expression(Expression::new_div(*inner_lhs, *rhs));
                        }
                    }
                    _ => {}
                }
            }

            e
        }
        Op::Eql((lhs, rhs)) => {
            if lhs.max < rhs.min || lhs.min > rhs.max {
                // e != f if e < f or f < e
                return Expression::new_const(0);
            }

            if let Op::Const(c_lhs) = lhs.op {
                if let Op::Const(c_rhs) = rhs.op {
                    // C == D
                    return Expression::new_const((c_lhs == c_rhs) as i64);
                }
            }

            e
        }
        _ => e,
    }
}

#[derive(Debug, Clone, PartialEq)]
enum Op {
    Mul((Box<Expression>, Box<Expression>)),
    Add((Box<Expression>, Box<Expression>)),
    Mod((Box<Expression>, Box<Expression>)),
    Div((Box<Expression>, Box<Expression>)),
    Eql((Box<Expression>, Box<Expression>)),
    Const(i64),
    Var(u8),
}

#[derive(Clone, Debug, PartialEq)]
struct Expression {
    op: Op,
    min: i64,
    max: i64,
}

impl Expression {
    fn new_var(var: u8) -> Self {
        Expression {
            op: Op::Var(var),
            min: 1,
            max: 9,
        }
    }

    fn new_const(r#const: i64) -> Self {
        Expression {
            op: Op::Const(r#const),
            min: r#const,
            max: r#const,
        }
    }

    fn new_mul(lhs: Expression, rhs: Expression) -> Self {
        Expression {
            op: Op::Mul((Box::from(lhs.clone()), Box::from(rhs.clone()))),
            min: lhs.min * rhs.min,
            max: lhs.max * rhs.max,
        }
    }

    fn new_add(lhs: Expression, rhs: Expression) -> Self {
        Expression {
            op: Op::Add((Box::from(lhs.clone()), Box::from(rhs.clone()))),
            min: lhs.min + rhs.min,
            max: lhs.max + rhs.max,
        }
    }

    fn new_mod(lhs: Expression, rhs: Expression) -> Self {
        if let Expression {
            op: Op::Const(m), ..
        } = rhs
        {
            Expression {
                op: Op::Mod((Box::from(lhs.clone()), Box::from(rhs.clone()))),
                min: 0,
                max: m - 1,
            }
        } else {
            panic!("mod with variable rhs")
        }
    }

    fn new_div(lhs: Expression, rhs: Expression) -> Self {
        Expression {
            op: Op::Div((Box::from(lhs.clone()), Box::from(rhs.clone()))),
            min: lhs.min / rhs.min,
            max: lhs.max / rhs.max,
        }
    }

    fn new_eql(lhs: Expression, rhs: Expression) -> Self {
        Expression {
            op: Op::Eql((Box::from(lhs.clone()), Box::from(rhs.clone()))),
            min: 0,
            max: 1,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn p1_works() {
        assert_eq!(p1(raw_input()), 0);
    }

    #[test]
    fn p2_works() {
        assert_eq!(p2(raw_input()), 0);
    }

    fn raw_input<'a>() -> &'a [u8] {
        b""
    }
}
