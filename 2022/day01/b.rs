use std::io::{self, Read};

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let mut inp: Vec<i32> = input.split("\n\n")
        .map(|elf_load| {
            elf_load
                .lines()
                .map(|calories| calories.parse::<i32>().unwrap())
                .sum()
        }).collect::<Vec<_>>().clone();

    inp.sort();


    let result: i32 = inp.into_iter().rev().take(3).sum();

    eprintln!("{}", result as i32);
    Ok(())
}
