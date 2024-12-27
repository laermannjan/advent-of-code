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
    let n_cycles = 240;
    let width = 40;

    let mut x = 1;
    let mut cycle = 0;

    let mut pixels = vec!['-'; n_cycles];

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

        let sprite_positions = x - 1..=x + 1;

        for _ in 0..exec_cycles {
            cycle += 1;

            let horizontal_pos = (cycle - 1) % width;

            pixels[cycle as usize - 1] = if sprite_positions.contains(&horizontal_pos) {
                '#'
            } else {
                '.'
            };
            // dbg!(cycle, x, x_delta);
        }
        x += x_delta;
    }

    eprintln!();
    pixels.chunks(40).for_each(|row| {
        eprintln!("{}", row.iter().collect::<String>());
    });

    todo!("Read the message from the pixels")
}
