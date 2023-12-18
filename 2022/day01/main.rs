use std::path::Path;

use utils::Day;

type Input = Vec<i32>;

pub fn parse_input(input: &str) -> Input {
    input
        .split("\n\n")
        .map(|elf_load| {
            elf_load
                .lines()
                .map(|calories| calories.parse::<i32>().unwrap())
                .sum()
        })
        .collect::<Vec<_>>()
}
pub fn part_one(input: String) -> Option<i32> {
    parse_input(&input).into_iter().max()
}

pub fn part_two(input: String) -> Option<i32> {
    let mut input = parse_input(&input);
    input.sort();
    Some(input.into_iter().rev().take(3).sum())
}

pub fn main() {
    let day = Day {
        loc: Path::new(file!()).into(),
        part_one,
        part_two,
    };
    let _ = day.run();
}
