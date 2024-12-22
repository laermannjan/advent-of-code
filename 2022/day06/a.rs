use std::io::{self, Read};

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let pos = input
        .as_bytes()
        .windows(4)
        .position(|b| !(0..=2).any(|i| (i + 1..=3).any(|j| b[i] == b[j])))
        .unwrap();

    eprintln!("{}", pos as i32 + 4);
    Ok(())
}

