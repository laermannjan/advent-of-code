use std::collections::HashSet;

type Input = Vec<(String, String)>;

pub fn parse_input(input: &str) -> Input {
    input
        .lines()
        .map(|rucksack| {
            let comp_size = rucksack.len() / 2;
            let comp_a = &rucksack[..comp_size].to_string();
            let comp_b = &rucksack[comp_size..].to_string();
            (comp_a.to_owned(), comp_b.to_owned())
        })
        .collect::<Vec<_>>()
}

fn convert_to_priority(c: char) -> i32 {
    ('a'..='z').chain('A'..='Z').position(|x| x == c).unwrap() as i32 + 1
}
pub fn part_one(input: Input) -> Option<i32> {
    let priorities = input
        .into_iter()
        .map(|(comp_a, comp_b)| {
            let common_char = comp_a.chars().find(|c| comp_b.contains(*c)).unwrap();
            convert_to_priority(common_char)
        })
        .sum::<i32>();
    Some(priorities)
}

pub fn part_two(input: Input) -> Option<i32> {
    let priorities = input
        .into_iter()
        .map(|(comp_a, comp_b)| format!("{}{}", comp_a, comp_b))
        .collect::<Vec<_>>()
        .chunks(3)
        .map(|rucksacks| {
            let common_char = rucksacks[0]
                .chars()
                .find(|c| rucksacks[1].contains(*c) && rucksacks[2].contains(*c))
                .unwrap();
            convert_to_priority(common_char)
        })
        .sum::<i32>();
    Some(priorities)
}

utils::main!(2022, 3);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 3);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 3, 1);
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 3);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 3, 2);
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
