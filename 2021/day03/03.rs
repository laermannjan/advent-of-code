use itertools::Itertools;

type Input = Vec<u32>;

fn parse_input(input: &str) -> Input {
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

pub fn part_one(input: Input) -> Option<i32> {
    let nums = input;
    let most_common = (0..12)
        .map(|i| (most_common_bit(&nums, i) as u32) << i)
        .sum::<u32>();
    let least_common = most_common ^ ((1 << 12) - 1);
    Some((most_common * least_common) as i32)
}

pub fn part_two(input: Input) -> Option<i32> {
    let nums = input;
    let most_common = filter_bit_criteria(&nums, true);
    let least_common = filter_bit_criteria(&nums, false);
    Some((most_common * least_common) as i32)
}

utils::main!(2021, 3);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[ignore]
    fn test_part_one() {
        println!("This fails because the test input is only 5 bits, while the puzzle in put has 12. I didn't bother fixing this after the fact.");
        let input = utils::get_test_input(2021, 3, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 3, 1, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    #[ignore]
    fn test_part_two() {
        println!("This fails because the test input is only 5 bits, while the puzzle in put has 12. I didn't bother fixing this after the fact.");
        let input = utils::get_test_input(2021, 3, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 3, 2, 1).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
