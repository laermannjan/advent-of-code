use std::io::{self, Read};

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let result: i32 = input.split("\n\n")
        .map(|elf_load| {
            elf_load
                .lines()
                .map(|calories| calories.parse::<i32>().unwrap())
                .sum()
        }).max().unwrap();

    eprintln!("{}", result as i32);
    Ok(())

}

