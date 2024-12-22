use std::io::{self, Read};

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let pos = input
        .as_bytes()
        .windows(14)
        .position(|b| !(0..=12).any(|i| (i + 1..=13).any(|j| b[i] == b[j])))
        .unwrap();

    eprintln!("{}", pos as i32 + 14);
    Ok(())
}

