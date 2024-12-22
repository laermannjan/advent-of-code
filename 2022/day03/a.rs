use std::io::{self, Read};


fn convert_to_priority(c: char) -> i32 {
    ('a'..='z').chain('A'..='Z').position(|x| x == c).unwrap() as i32 + 1
}
pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;
    let priorities = input
        .lines()
        .map(|rucksack| {
            let comp_size = rucksack.len() / 2;
            let comp_a = &rucksack[..comp_size].to_string();
            let comp_b = &rucksack[comp_size..].to_string();
            (comp_a.to_owned(), comp_b.to_owned())
        })
        .map(|(comp_a, comp_b)| {
            let common_char = comp_a.chars().find(|c| comp_b.contains(*c)).unwrap();
            convert_to_priority(common_char)
        })
        .sum::<i32>();
    eprintln!("{}", priorities as i32);
    Ok(())
}

