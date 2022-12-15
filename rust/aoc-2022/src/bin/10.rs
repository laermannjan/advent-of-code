use nom::{branch::alt, combinator::map, sequence::preceded};
use nom::{bytes::complete::tag, IResult};

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

pub fn parse_input(input: &str) -> Input {
    input
        .lines()
        .map(|l| Instruction::parse(l).unwrap().1)
        .collect()
}
pub fn part_one(input: Input) -> Option<i32> {
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
    Some(signal_strengths.iter().map(|(c, x)| c * x).sum())
}

pub fn part_two(input: Input) -> Option<i32> {
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

    println!();
    pixels.chunks(40).for_each(|row| {
        println!("{}", row.iter().collect::<String>());
    });

    todo!("Read the message from the pixels")
}

utils::main!(2022, 10);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let test_no = 2;
        let input = utils::get_test_input(2022, 10, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 10, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 10, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 10, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
