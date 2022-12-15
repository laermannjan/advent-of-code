use itertools::Itertools;

pub fn parse_input(input: &str) -> &str {
    input
}
pub fn part_one(input: &str) -> Option<i32> {
    let pos = input
        .as_bytes()
        .windows(4)
        .position(|b| !(0..=2).any(|i| (i + 1..=3).any(|j| b[i] == b[j])))
        .unwrap();

    Some(pos as i32 + 4)
}

pub fn part_two(input: &str) -> Option<i32> {
    let pos = input
        .as_bytes()
        .windows(14)
        .position(|b| !(0..=12).any(|i| (i + 1..=13).any(|j| b[i] == b[j])))
        .unwrap();

    Some(pos as i32 + 14)
}

utils::main!(2022, 6);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 6, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 6, 1, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 6, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 6, 2, 1).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
