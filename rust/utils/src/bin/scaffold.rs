use std::{fs::OpenOptions, io::Write, process::Command};

use clap::Parser;

// parser that reads the year as u32 and day as u8
#[derive(Parser, Debug)]
#[clap(author, version, about)]
pub struct Scaffold {
    year: u32,
    day: u8,
}

const TEMPLATE: &str = r###"utils = { path = \"../utils\" }"###;
const DAY_TEMPLATE: &str = r###"type Input = Vec<i32>;

pub fn parse_input(input: &str) -> Input {
    input.lines().map(|s| s.parse::<i32>().unwrap_or(0)).collect()
}
pub fn part_one(input: Input) -> Option<i32> {
    None
}

pub fn part_two(input: Input) -> Option<i32> {
    None
}

utils::main!(YEAR, DAY);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(YEAR, DAY);
        let parsed_input = parse_input(&input);
        assert_eq!(part_one(parsed_input), Some(1337));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(YEAR, DAY);
        let parsed_input = parse_input(&input);
        assert_eq!(part_two(parsed_input), Some(1337));
    }
}
"###;

fn create_crate(year: u32) -> Result<(), std::io::Error> {
    let crate_path_prefix = format!("{}/../", env!("CARGO_MANIFEST_DIR"));
    Command::new("cargo")
        .args(&["new", &format!("aoc-{}", year)])
        .current_dir(format!("{}", crate_path_prefix))
        .output()?;

    // append the crate to the workspace
    if !Command::new("bash")
        .args(&[
            "-c",
            &format!(
                "echo {} >> {}/aoc-{}/Cargo.toml",
                TEMPLATE, crate_path_prefix, year
            ),
        ])
        .status()?
        .success()
    {
        return Err(std::io::Error::new(
            std::io::ErrorKind::Other,
            "Failed to add utils as a dependency",
        ));
    }

    Ok(())
}

fn create_file(path: &str, content: &str) -> Result<(), std::io::Error> {
    //create path of parent dir
    std::fs::create_dir_all(std::path::Path::new(path).parent().unwrap())?;

    let mut file = OpenOptions::new().write(true).create_new(true).open(path)?;
    file.write_all(content.as_bytes())?;
    Ok(())
}
fn create_day(year: u32, day: u8) -> Result<(), std::io::Error> {
    let crate_path_prefix = format!("{}/../", env!("CARGO_MANIFEST_DIR"));
    let path = format!("{}aoc-{}/src/bin/{:02}.rs", crate_path_prefix, year, day);
    let template = DAY_TEMPLATE
        .replace("YEAR", &format!("{}", year))
        .replace("DAY", &format!("{}", day));

    if std::path::Path::new(&path).exists() {
        return Err(std::io::Error::new(
            std::io::ErrorKind::AlreadyExists,
            "File already exists",
        ));
    }

    create_file(&path, &template)?;
    Ok(())
}

fn scaffold_year(year: u32) -> Result<(), std::io::Error> {
    let crate_path = format!("{}/../aoc-{}", env!("CARGO_MANIFEST_DIR"), year);

    if std::path::Path::new(&crate_path).exists() {
        return Err(std::io::Error::new(
            std::io::ErrorKind::AlreadyExists,
            "Crate already exists",
        ));
    }

    create_crate(year)?;
    Ok(())
}

fn scaffold_day(year: u32, day: u8) -> Result<(), std::io::Error> {
    create_day(year, day)?;
    Ok(())
}

fn main() {
    let scaffold = Scaffold::parse();

    match scaffold_year(scaffold.year) {
        Ok(_) => println!(
            "Created create for {} => ./aoc-{}",
            scaffold.year, scaffold.year
        ),
        Err(e) => match e.kind() {
            std::io::ErrorKind::AlreadyExists => {
                println!(
                    "Crate for {} already exists, skipping creation",
                    scaffold.year
                );
            }
            _ => {
                println!("Error: {}", e);
                std::process::exit(1);
            }
        },
    }
    match scaffold_day(scaffold.year, scaffold.day) {
        Ok(_) => println!(
            "Created binary for day {} => ./aoc-{}/src/bin/{:02}.rs",
            scaffold.day, scaffold.year, scaffold.day
        ),
        Err(e) => match e.kind() {
            std::io::ErrorKind::AlreadyExists => {
                println!(
                    "File for Day {} binary already exists, skipping creation",
                    scaffold.day
                );
            }
            _ => {
                println!("Error: {}", e);
                std::process::exit(1);
            }
        },
    }
}
