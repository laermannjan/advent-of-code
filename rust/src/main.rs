use regex::Regex;
use std::env;
use std::fs;

#[allow(non_snake_case)]
mod Year2021;
use crate::Year2021::day01a;
use crate::Year2021::day01b;
use crate::Year2021::day02a;
use crate::Year2021::day02b;
use crate::Year2021::day03a;
use crate::Year2021::day03b;

fn main() {
    let args: Vec<String> = env::args().collect();
    let year = &args[1].parse::<u16>().expect("Couldn't parse <year>");

    let re = Regex::new(r"(2[0-5]|1[0-9]|[1-9])([a-z])").unwrap();
    let cap = re.captures(&args[2]).unwrap();
    let (day, part) = (
        &cap[1].parse::<u8>().expect("Coudn't parse <day>"),
        &cap[2].parse::<char>().expect("Coudn't parse <part>"),
    );

    let fname = format!("data/{0}/day{1:02}/day{1:02}_input_puzzle.txt", year, day);
    let err_msg = format!("Couldn't read the input file: {}!\n", fname);
    let puzzle_input = fs::read_to_string(fname).expect(&err_msg);

    let result = match args[2].as_str() {
        "1a" => day01a(&puzzle_input),
        "1b" => day01b(&puzzle_input),
        "2a" => day02a(&puzzle_input),
        "2b" => day02b(&puzzle_input),
        "3a" => day03a(&puzzle_input),
        "3b" => day03b(&puzzle_input),
        _ => 0,
    };
    println!("Solution for Advent of Code {year}, Day {day} - Part {part}: {result}");
}
