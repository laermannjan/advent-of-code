use itertools::Itertools;
use std::io::{self, Read};

fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;
    let (horizontal, depth, _) = input
        .split_whitespace()
        .tuples()
        .map(|(inst, val)| (inst.as_bytes()[0], val.parse().unwrap()))
        .collect::<Vec<_>>()
        .iter()
        .fold(
        (0, 0, 0),
        |(horizontal, depth, aim), (inst, val)| match inst {
            b'f' => (horizontal + val, depth + aim * val, aim),
            b'u' => (horizontal, depth, aim - val),
            b'd' => (horizontal, depth, aim + val),
            _ => unreachable!(),
        },
    );

    eprintln!("{}", (horizontal * depth) as i32);
    Ok(())
}

