use itertools::Itertools;

fn parse_input(input: &str) -> Vec<u32> {
    input
        .lines()
        .map(|l| u32::from_str_radix(l, 2).unwrap())
        .collect_vec()
}

pub fn day03a(input: &str) -> usize {
    let nums = parse_input(input);
    let mut most_common = 0;
    for i in 0..12 {
        let num_ones = nums.iter().map(|n| (n >> i) as usize & 1).sum::<usize>();
        if 2 * num_ones > nums.len() {
            most_common += 1 << i;
        }
    }
    (most_common * (most_common ^ ((1 << 12) - 1))) as usize
}
