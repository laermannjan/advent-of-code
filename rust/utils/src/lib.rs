use std::{process::Command, time::Duration};

pub const ANSI_ITALIC: &str = "\x1b[3m";
pub const ANSI_BOLD: &str = "\x1b[1m";
pub const ANSI_RESET: &str = "\x1b[0m";

pub fn timed_func<I, R>(func: fn(I) -> R, input: I) -> (R, String) {
    let start = std::time::Instant::now();
    let result = func(input);
    let duration = start.elapsed();

    let formatted_duration = format!("{} µs", duration.as_micros());
    (result, formatted_duration)
}

#[macro_export]
macro_rules! solve {
    ($part:literal, $func:ident, $input:expr) => {
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
    };
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
            print!(" 🧩 ");
            utils::solve!("a", part_one, &input);
            print!(" 🧩 ");
            utils::solve!("b", part_two, &input);
        }
    };
}

pub fn get_puzzle_input(year: u32, day: u8) -> String {
    read_data_file(year, day, None, None)
}

fn read_data_file(year: u32, day: u8, test: Option<u8>, part: Option<u8>) -> String {
    let path = get_data_path(year, day, test, part);
    std::fs::read_to_string(path.clone()).expect(&format!("could not open data file {}", path))
}

fn get_data_path(year: u32, day: u8, test: Option<u8>, part: Option<u8>) -> String {
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
                    ])
                    .output()
                    .expect("Failed to execute get_data_path");
            }
        },
        None => {
            output = Command::new(bin_path)
                .args(&["-y", &format!("{}", year), "-d", &format!("{}", day)])
                .output()
                .expect("Failed to execute get_data_path");
        }
    }

    if output.status.success() {
        String::from_utf8(output.stdout).unwrap().trim().to_string()
    } else {
        panic!("Failed to read output from get_data_path");
    }
}
