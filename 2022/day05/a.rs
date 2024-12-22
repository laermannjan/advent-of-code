use itertools::Itertools;
use std::io::{self, Read};

type Stack = Vec<char>;
type Instruction = (i32, i32, i32);

fn parse_instructions(input: &str) -> Vec<Instruction> {
    input
        .lines()
        .map(|line| {
            let (n, origin, target) = line
                .split(" ")
                .skip(1)
                .step_by(2)
                .map(|n| n.parse::<i32>().unwrap())
                .collect_tuple()
                .unwrap();
            (n, origin - 1, target - 1)
        })
        .collect::<Vec<(i32, i32, i32)>>()
}

fn parse_stack_configuration(input: &str) -> Vec<Stack> {
    // 3 cols for a stack and 1 between them:
    // n_stacks * 3 + n_stacks - 1 = n_stacks * 4 - 1
    let n_stacks = (input.lines().last().unwrap().len() + 1) / 4;
    let mut stacks = vec![vec![]; n_stacks];
    input.lines().rev().skip(1).for_each(|line| {
        line.chars()
            .skip(1)
            .step_by(4)
            .enumerate()
            .filter(|(_i, c)| *c != ' ')
            .for_each(|(i, c)| stacks[i].push(c))
    });
    stacks
}

fn read_top_create(stacks: Vec<Stack>) -> String {
    stacks.iter().map(|s| s.last().unwrap()).collect::<String>()
}

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let (config, instructions) = input.split_once("\n\n").unwrap();

    let mut stacks = parse_stack_configuration(config);
    let instructions = parse_instructions(instructions);

    instructions.iter().for_each(|(n, origin, target)| {
        (0..*n).for_each(|_| {
            let c = stacks[*origin as usize].pop().unwrap();
            stacks[*target as usize].push(c);
        });
    });
    eprintln!("{}", read_top_create(stacks));
    Ok(())
}
