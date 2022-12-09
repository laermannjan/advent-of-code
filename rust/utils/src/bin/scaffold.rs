use std::{error::Error, fs::OpenOptions, io::Write, process::Command};

use clap::Parser;

// parser that reads the year as u32 and day as u8
#[derive(Parser, Debug)]
#[clap(author, version, about)]
pub struct Scaffold {
    year: u32,
    day: u8,
}

const TEMPLATE: &str = r###"utils = { path = \"../utils\" }"###;
const MODULE_TEMPLATE: &str = r###"pub fn part_one(input: &str) -> Option<u32> {
    None
}
pub fn part_two(input: &str) -> Option<u32> {
    None
}
fn main() {
    let input = &utils:read_file("inputs", DAY);
    utils::solve!(1, part_one, input);
    utils::solve!(2, part_two, input);
}
#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_part_one() {
        let input = utils:read_file("examples", DAY);
        assert_eq!(part_one(&input), None);
    }
    #[test]
    fn test_part_two() {
        let input = utils::read_file("examples", DAY);
        assert_eq!(part_two(&input), None);
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
    create_file(
        &format!("{}aoc-{}/src/bin/{:02}.rs", crate_path_prefix, year, day),
        MODULE_TEMPLATE,
    )?;
    Ok(())
}

fn scaffold_year(year: u32) -> Result<(), std::io::Error> {
    create_crate(year)?;
    Ok(())
}

fn scaffold_day(year: u32, day: u8) -> Result<(), std::io::Error> {
    create_day(year, day)?;
    Ok(())
}

fn main() {
    let scaffold = Scaffold::parse();
    println!("{:?}", scaffold);

    if let Err(e) = scaffold_year(scaffold.year) {
        eprintln!("Error: {}", e);
    } else if let Err(e) = scaffold_day(scaffold.year, scaffold.day) {
        eprintln!("Error: {}", e);
    }
}
