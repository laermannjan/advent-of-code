use utils::grid::Direction;

type Input = Vec<Direction>;

pub fn parse_input(input: &str) -> Input {
    input
        .chars()
        .map(|c| match c {
            '>' => Direction::East,
            '<' => Direction::West,
            _ => panic!("invalid input"),
        })
        .collect()
}
pub fn part_one(input: Input) -> Option<i32> {
    None
}

pub fn part_two(input: Input) -> Option<i32> {
    None
}

utils::main!(2022, 17);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 17, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 17, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 17, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 17, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
