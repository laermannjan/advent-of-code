use std::path::Path;
use utils::{get_puzzle_bin, run_day};

fn main() {
    for day in 1..=25 {
        if !Path::new(&get_puzzle_bin(2021, day)).exists() {
            println!(" 🚫 Day {} binary not found.", day);
        } else if let Err(_) = run_day(2021, day) {
            println!(" 🚫 Day {} failed to run", day)
        }
        println!("\n🎄🎊🎄✨🎄🎁🎄🎊🎄✨🎄🎁🎄\n");
    }
}

