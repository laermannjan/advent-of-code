use itertools::Itertools;
use std::io::{self, Read};

fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;
    let numbers: Vec<i32> = input.lines().map(|s| s.parse::<i32>().unwrap()).collect();
    let result = numbers.iter().tuple_windows().filter(|(a, b)| a < b).count();
    eprintln!("{}", result);
    Ok(())
}
