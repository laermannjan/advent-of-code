use nom::{branch::alt, combinator::map, sequence::preceded};
use nom::{bytes::complete::tag, IResult};
use std::io::{self, Read};

type Input = Vec<Instruction>;

#[derive(Debug, Clone)]
pub enum Instruction {
    Noop,
    Addx(i32),
}

impl Instruction {
    fn parse(input: &str) -> IResult<&str, Self> {
        alt((
            map(tag("noop"), |_| Instruction::Noop),
            map(preceded(tag("addx "), nom::character::complete::i32), |x| {
                Instruction::Addx(x)
            }),
        ))(input)
    }
}

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    input
        .lines()
        .map(|l| Instruction::parse(l).unwrap().1)
        .collect()
}
pub fn main() -> io::Result<()> {
    let input = parse_input();
    let mut x = 1;
    let mut cycle = 0;

    let mut signal_strengths: Vec<(i32, i32)> = vec![];

    for inst in input {
        let exec_cycles;
        let x_delta;

        match inst {
            Instruction::Noop => {
                exec_cycles = 1;
                x_delta = 0;
            }
            Instruction::Addx(y) => {
                exec_cycles = 2;
                x_delta = y;
            }
        }

        for _ in 0..exec_cycles {
            cycle += 1;
            if (cycle + 20) % 40 == 0 {
                signal_strengths.push((cycle, x));
            }
            // dbg!(cycle, x, x_delta);
        }
        x += x_delta;
    }
    eprintln!(
        "{}",
        signal_strengths.iter().map(|(c, x)| c * x).sum::<i32>()
    );
    Ok(())
}
