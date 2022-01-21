use itertools::Itertools;

fn parse_input(input: &str) -> Vec<i32> {
    let numbers: Vec<i32> = input.lines().map(|s| s.parse::<i32>().unwrap()).collect();
    return numbers;
}

pub fn day01a(input: &str) -> usize {
    let numbers = parse_input(input);
    let result = numbers
        .iter()
        .tuple_windows()
        .filter(|(a, b)| a < b)
        .count();
    return result;
}

pub fn day01b(input: &str) -> usize {
    let numbers = parse_input(input);
    let result = numbers
        .iter()
        .tuple_windows()
        .filter(|(a, _, _, d)| a < d)
        .count();
    return result;
}
