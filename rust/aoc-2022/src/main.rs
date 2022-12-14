use std::path::Path;
use utils::{get_puzzle_bin, run_day};

fn main() {
    for day in 1..=25 {
        if !Path::new(&get_puzzle_bin(2022, day)).exists() {
            println!(" š« Day {} binary not found.", day);
        } else if let Err(_) = run_day(2022, day) {
            println!(" š« Day {} failed to run", day)
        }
        println!("\nšššāØšššššāØššš\n");
    }
}
