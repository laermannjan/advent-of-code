type Input = Vec<(i32, i32, i32, i32)>;

pub fn parse_input(input: &str) -> Input {
    input
        .lines()
        .map(|line| {
            let (elf_one, elf_two) = line.split_once(",").unwrap();
            let (one_lower, one_upper) = elf_one.split_once("-").unwrap();
            let (two_lower, two_upper) = elf_two.split_once("-").unwrap();
            (
                one_lower.parse::<i32>().unwrap(),
                one_upper.parse::<i32>().unwrap(),
                two_lower.parse::<i32>().unwrap(),
                two_upper.parse::<i32>().unwrap(),
            )
        })
        .collect::<Vec<_>>()
}
pub fn part_one(input: Input) -> Option<i32> {
    Some(
        input
            .into_iter()
            .filter(|(a, b, c, d)| (a <= c && b >= d) || (c <= a && d >= b))
            .count() as i32,
    )
}

pub fn part_two(input: Input) -> Option<i32> {
    Some(
        input
            .into_iter()
            .filter(|(a, b, c, d)| {
                (a <= c && b >= d) || (c <= a && d >= b) || (a >= c && a <= d) || (b >= c && a <= d)
            })
            .count() as i32,
    )
}

utils::main!(2022, 4);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 4);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 4, 1);
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 4);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 4, 2);
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
