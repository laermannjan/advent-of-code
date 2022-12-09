use itertools::Itertools;

type Input = Vec<i32>;

fn parse_input(input: &str) -> Input {
    let numbers: Vec<i32> = input.lines().map(|s| s.parse::<i32>().unwrap()).collect();
    return numbers;
}

pub fn part_one(input: Input) -> Option<i32> {
    let result = input.iter().tuple_windows().filter(|(a, b)| a < b).count();
    return Some(result as i32);
}

pub fn part_two(input: Input) -> Option<i32> {
    let result = input
        .iter()
        .tuple_windows()
        .filter(|(a, _, _, d)| a < d)
        .count();
    return Some(result as i32);
}

utils::main!(2021, 1);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2021, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 1, 1);
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2021, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 1, 2);
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
