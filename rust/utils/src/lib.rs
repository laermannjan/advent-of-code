use std::process::Command;

pub mod grid;
pub mod path_finding;

pub const ANSI_ITALIC: &str = "\x1b[3m";
pub const ANSI_BOLD: &str = "\x1b[1m";
pub const ANSI_RESET: &str = "\x1b[0m";

pub fn timed_func<I: Clone, R>(func: impl Fn(I) -> R, input: I) -> (R, String) {
    let start = std::time::Instant::now();
    let result = func(input.clone());
    let duration = start.elapsed();

    let formatted_duration = format!("{} µs", duration.as_micros());
    (result, formatted_duration)
}

#[macro_export]
macro_rules! solve {
    ($part:literal, $func:ident, $input:expr) => {{
        let (result, duration) = utils::timed_func($func, $input);

        match result {
            Some(r) => println!(
                "Part {}: {}{}{} {}(took: {}){}",
                $part,
                utils::ANSI_BOLD,
                r,
                utils::ANSI_RESET,
                utils::ANSI_ITALIC,
                duration,
                utils::ANSI_RESET
            ),
            None => println!(
                "Part {}: {}{}Not solved{}",
                $part,
                utils::ANSI_BOLD,
                utils::ANSI_ITALIC,
                utils::ANSI_RESET
            ),
        }
    }};
}

#[macro_export]
macro_rules! parse {
    ($func:ident, $input:expr) => {{
        let (parsed, duration) = utils::timed_func($func, $input);
        parsed
    }};
}

#[macro_export]
macro_rules! main {
    ($year:expr, $day:expr) => {
        pub fn main() {
            println!(
                " 🧝 Day {} {}[{}]{}",
                $day,
                utils::ANSI_ITALIC,
                $year,
                utils::ANSI_RESET
            );
            let input = utils::get_puzzle_input($year, $day);
            let parsed_input = utils::parse!(parse_input, &input);
            print!(" 🧩 ");
            utils::solve!("a", part_one, parsed_input.clone());
            print!(" 🧩 ");
            utils::solve!("b", part_two, parsed_input.clone());
        }
    };
}

pub fn run_day(year: u32, day: u8) -> Result<(), std::io::Error> {
    let aoc_year = format!("aoc-{}", year);
    let padded_day = format!("{:02}", day);
    let args = vec![
        "run",
        "-p",
        &aoc_year,
        "--release",
        "--quiet",
        "--bin",
        &padded_day,
    ];
    let mut cmd = Command::new("cargo").args(&args).spawn()?;

    cmd.wait()?;

    Ok(())
}

pub fn get_puzzle_bin(year: u32, day: u8) -> String {
    format!("aoc-{}/src/bin/{:02}.rs", year, day)
}

pub fn get_puzzle_input(year: u32, day: u8) -> String {
    read_data_file(&get_puzzle_input_path(year, day))
}

pub fn get_test_input(year: u32, day: u8, test_no: u8) -> String {
    read_data_file(&get_test_input_path(year, day, test_no))
}

pub fn get_test_result(year: u32, day: u8, part: u8, test_no: u8) -> String {
    let result = read_data_file(&get_test_result_path(year, day, part, test_no))
        .trim()
        .to_string();
    if result.is_empty() {
        return "1337".to_string();
    } else {
        return result;
    }
}

fn read_data_file(data_path: &str) -> String {
    std::fs::read_to_string(data_path.clone())
        .expect(&format!("could not open data file {}", &data_path))
}

pub fn get_puzzle_input_path(year: u32, day: u8) -> String {
    get_data_path(year, day, None, None, false)
}
pub fn create_puzzle_input_dummy(year: u32, day: u8) -> String {
    get_data_path(year, day, None, None, true)
}

pub fn get_test_input_path(year: u32, day: u8, test_no: u8) -> String {
    get_data_path(year, day, Some(test_no), None, false)
}
pub fn create_test_input_dummy(year: u32, day: u8) -> String {
    get_data_path(year, day, Some(1), None, true)
}
pub fn get_test_result_path(year: u32, day: u8, part: u8, test_no: u8) -> String {
    get_data_path(year, day, Some(test_no), Some(part), false)
}
pub fn create_test_result_dummy(year: u32, day: u8, part: u8, test_no: u8) -> String {
    get_data_path(year, day, Some(test_no), Some(part), true)
}

fn get_data_path(
    year: u32,
    day: u8,
    test: Option<u8>,
    part: Option<u8>,
    create_if_not_exist: bool,
) -> String {
    let bin_path = format!("{}/../../get_data_path", env!("CARGO_MANIFEST_DIR"));
    let output;

    match test {
        Some(test) => match part {
            Some(part) => {
                output = Command::new(bin_path)
                    .args(&[
                        "-y",
                        &format!("{}", year),
                        "-d",
                        &format!("{}", day),
                        "-t",
                        &format!("{}", test),
                        "-p",
                        &format!("{}", part),
                        "--create-if-not-exist",
                        &format!("{}", create_if_not_exist),
                    ])
                    .output()
                    .expect("Failed to execute get_data_path");
            }
            None => {
                output = Command::new(bin_path)
                    .args(&[
                        "-y",
                        &format!("{}", year),
                        "-d",
                        &format!("{}", day),
                        "-t",
                        &format!("{}", test),
                        "--create-if-not-exist",
                        &format!("{}", create_if_not_exist),
                    ])
                    .output()
                    .expect("Failed to execute get_data_path");
            }
        },
        None => {
            output = Command::new(bin_path)
                .args(&[
                    "-y",
                    &format!("{}", year),
                    "-d",
                    &format!("{}", day),
                    "--create-if-not-exist",
                    &format!("{}", create_if_not_exist),
                ])
                .output()
                .expect("Failed to execute get_data_path");
        }
    }

    if output.status.success() {
        String::from_utf8(output.stdout).unwrap().trim().to_string()
    } else {
        eprintln!(
            "Stderr: {}",
            String::from_utf8(output.stderr).unwrap().trim()
        );
        panic!("Failed to read output from get_data_path");
    }
}
