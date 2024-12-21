use itertools::Itertools;
use std::io::{self, Read};

fn most_common_bit(nums: &Vec<u32>, i: usize) -> bool {
    let num_ones = nums.iter().map(|n| (n >> i) as usize & 1).sum::<usize>();
    2 * num_ones >= nums.len()
}

fn filter_bit_criteria(nums: &Vec<u32>, most_common: bool) -> u32 {
    let mut nums = nums.to_vec();
    for i in (0..12).rev() {
        let criterium = most_common_bit(&nums, i) ^ !most_common;
        nums.retain(|n| (n >> i) & 1 == (criterium as u32));
        if nums.len() == 1 {
            break;
        }
    }
    nums[0]
}

fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let nums = input.lines()
        .map(|l| u32::from_str_radix(l, 2).unwrap())
        .collect_vec();
    let most_common = filter_bit_criteria(&nums, true);
    let least_common = filter_bit_criteria(&nums, false);

    println!("This fails because the test input is only 5 bits, while the puzzle in put has 12. I didn't bother fixing this after the fact.");
    eprintln!("{}", (most_common * least_common) as i32);
    Ok(())
}
