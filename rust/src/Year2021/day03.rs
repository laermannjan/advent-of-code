use itertools::Itertools;

fn parse_input(input: &str) -> Vec<u32> {
    input
        .lines()
        .map(|l| u32::from_str_radix(l, 2).unwrap())
        .collect_vec()
}

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

pub fn day03a(input: &str) -> usize {
    let nums = parse_input(input);
    let most_common = (0..12)
        .map(|i| (most_common_bit(&nums, i) as u32) << i)
        .sum::<u32>();
    let least_common = most_common ^ ((1 << 12) - 1);
    (most_common * least_common) as usize
}

pub fn day03b(input: &str) -> usize {
    let nums = parse_input(input);
    let most_common = filter_bit_criteria(&nums, true);
    let least_common = filter_bit_criteria(&nums, false);
    (most_common * least_common) as usize
}
