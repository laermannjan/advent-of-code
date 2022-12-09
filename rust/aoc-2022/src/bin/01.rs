type Input = Vec<i32>;

pub fn parse_input(input: &str) -> Input {
    input
        .split("\n\n")
        .map(|elf_load| {
            elf_load
                .lines()
                .map(|calories| calories.parse::<i32>().unwrap())
                .sum()
        })
        .collect::<Vec<_>>()
}
pub fn part_one(input: Input) -> Option<i32> {
    input.into_iter().max()
}

pub fn part_two(input: Input) -> Option<i32> {
    let mut input = input.clone();
    input.sort();
    Some(input.into_iter().rev().take(3).sum())
}

utils::main!(2022, 1);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 1, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 1, 2).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
