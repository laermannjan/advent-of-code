use std::io::{self, Read};


pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let result: i32 = input
        .lines()
        .map(|line| {
            let (elf_one, elf_two) = line.split_once(",").unwrap();
            let (one_lower, one_upper) = elf_one.split_once("-").unwrap();
            let (two_lower, two_upper) = elf_two.split_once("-").unwrap();
            (
                one_lower.parse::<i32>().unwrap(),
                one_upper.parse::<i32>().unwrap(),
                two_lower.parse::<i32>().unwrap(),
                two_upper.parse::<i32>().unwrap(),
            )
        })
        .filter(|(a, b, c, d)| {
            (a <= c && b >= d) || (c <= a && d >= b) || (a >= c && a <= d) || (b >= c && a <= d)
        })
        .count() as i32;

    eprintln!("{}", result as i32);
    Ok(())
}
