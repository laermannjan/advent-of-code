use std::{fs::OpenOptions, io::Write, process::Command};

use clap::Parser;
use utils::{
    create_puzzle_input_dummy, create_test_input_dummy, create_test_result_dummy,
    get_puzzle_input_path, get_test_input_path, get_test_result_path,
};

// parser that reads the year as u32 and day as u8
#[derive(Parser, Debug)]
#[clap(author, version, about)]
pub struct Scaffold {
    year: u32,
    day: u8,
}

const DEFAULT_RUN_TEMPLATE: &str = r###"
default-run = "aoc-YEAR"
"###;

const DEPS_TEMPLATE: &str = r###"utils = { path = \"../utils\" }"###;
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
        let test_no = 1;
        let input = utils::get_test_input(YEAR, DAY, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(YEAR, DAY, 1, test_no).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(YEAR, DAY, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(YEAR, DAY, 2, test_no).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
"###;

const MAIN_TEMPLATE: &str = r###"use std::path::Path;
use utils::{get_puzzle_bin, run_day};

fn main() {
    for day in 1..=25 {
        if !Path::new(&get_puzzle_bin(YEAR, day)).exists() {
            println!(" 🚫 Day {} binary not found.", day);
        } else if let Err(_) = run_day(YEAR, day) {
            println!(" 🚫 Day {} failed to run", day)
        }
        println!("\n🎄🎊🎄✨🎄🎁🎄🎊🎄✨🎄🎁🎄\n");
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
                DEPS_TEMPLATE, crate_path_prefix, year
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

    add_default_run_to_cargo_toml(year)?;

    create_file(
        &format!("{}/aoc-{}/src/main.rs", crate_path_prefix, year),
        &MAIN_TEMPLATE.replace("YEAR", &format!("{}", year)),
    )?;

    update_workspace_toml(year)?;

    Ok(())
}

fn add_default_run_to_cargo_toml(year: u32) -> Result<(), std::io::Error> {
    let toml_path = format!("{}/../aoc-{}/Cargo.toml", env!("CARGO_MANIFEST_DIR"), year);
    let mut toml = std::fs::read_to_string(&toml_path).unwrap();
    toml.insert_str(
        toml.find("[dependencies]").unwrap(),
        &DEFAULT_RUN_TEMPLATE.replace("YEAR", &format!("{}", year)),
    );
    create_file(&toml_path, &toml)?;
    println!("Added default-run to {}", toml_path);
    Ok(())
}

fn update_workspace_toml(year: u32) -> Result<(), std::io::Error> {
    let toml_path = format!("{}/../Cargo.toml", env!("CARGO_MANIFEST_DIR"));
    let mut toml = std::fs::read_to_string(&toml_path).unwrap();
    toml.insert_str(
        toml.find("members = [").unwrap() + 11,
        &format!(r###""aoc-{}", "###, year),
    );
    create_file(&toml_path, &toml)?;
    println!("🎉 Successfully added aoc-{} to the workspace", year);
    Ok(())
}

fn create_file(path: &str, content: &str) -> Result<(), std::io::Error> {
    //create path of parent dir
    std::fs::create_dir_all(std::path::Path::new(path).parent().unwrap())?;

    let mut file = OpenOptions::new().write(true).create(true).open(path)?;
    file.write_all(content.as_bytes())?;
    Ok(())
}

fn create_new_file(path: &str, content: &str) -> Result<(), std::io::Error> {
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
        println!("Solution binary already exists => {}", &path);
    } else {
        create_new_file(&path, &template)?;
        println!("Created solution binary => {}", &path);
    }
    Ok(())
}

fn create_data(year: u32, day: u8) -> Result<(), std::io::Error> {
    let puzzle_input_path = get_puzzle_input_path(year, day);

    if std::path::Path::new(&puzzle_input_path).exists() {
        println!("Puzzle input file already exists => {}", puzzle_input_path);
    } else {
        create_puzzle_input_dummy(year, day);
        println!("Created puzzle input file => {}", puzzle_input_path);
    }

    let test_input_path = get_test_input_path(year, day, 1);
    if std::path::Path::new(&test_input_path).exists() {
        println!("Test input file already exists => {}", test_input_path);
    } else {
        create_test_input_dummy(year, day);
        println!("Created test input file => {}", test_input_path);
    }

    for part in 1..=2 {
        let test_result_path = get_test_result_path(year, day, part, 1);
        if std::path::Path::new(&test_result_path).exists() {
            println!("Test result file already exists => {}", test_result_path);
        } else {
            create_test_result_dummy(year, day, part, 1);
            println!("Created test result file => {}", test_result_path);
        }
    }

    Ok(())
}

fn scaffold_year(year: u32) -> Result<(), std::io::Error> {
    let crate_path = format!("{}/../aoc-{}", env!("CARGO_MANIFEST_DIR"), year);

    if std::path::Path::new(&crate_path).exists() {
        println!("Crate already exists => {}", crate_path);
    } else {
        create_crate(year)?;
        println!("Created crate => {}", crate_path);
    }
    Ok(())
}

fn scaffold_day(year: u32, day: u8) -> Result<(), std::io::Error> {
    create_day(year, day)?;
    Ok(())
}

fn scaffold_data(year: u32, day: u8) -> Result<(), std::io::Error> {
    create_data(year, day)?;
    Ok(())
}

fn main() {
    let scaffold = Scaffold::parse();

    if let Err(e) = scaffold_year(scaffold.year) {
        println!("Failed to create year: {}", e);
        std::process::exit(1);
    }

    if let Err(e) = scaffold_day(scaffold.year, scaffold.day) {
        println!("Error: {}", e);
        std::process::exit(1);
    }

    if let Err(e) = scaffold_data(scaffold.year, scaffold.day) {
        println!("Error: {}", e);
        std::process::exit(1);
    }
}
