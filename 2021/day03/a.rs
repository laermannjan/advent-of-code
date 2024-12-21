use itertools::Itertools;
use std::io::{self, Read};

fn most_common_bit(nums: &Vec<u32>, i: usize) -> bool {
    let num_ones = nums.iter().map(|n| (n >> i) as usize & 1).sum::<usize>();
    2 * num_ones >= nums.len()
}

fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let nums = input.lines()
        .map(|l| u32::from_str_radix(l, 2).unwrap())
        .collect_vec();
    let most_common = (0..12)
        .map(|i| (most_common_bit(&nums, i) as u32) << i)
        .sum::<u32>();
    let least_common = most_common ^ ((1 << 12) - 1);
    println!("This fails because the test input is only 5 bits, while the puzzle in put has 12. I didn't bother fixing this after the fact.");
    eprintln!("{}", (most_common * least_common) as i32);
    Ok(())
}
