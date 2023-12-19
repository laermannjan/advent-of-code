use utils::Day;

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
pub fn part_one(input: &str) -> Option<i32> {
    parse_input(&input).into_iter().max()
}

pub fn part_two(input: &str) -> Option<i32> {
    let mut input = parse_input(&input);
    input.sort();
    Some(input.into_iter().rev().take(3).sum())
}

const DAY: Day = Day {
    loc: file!(),
    part_one,
    part_two,
};

pub fn main() {
    let _ = DAY.run();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = include_str!("./example.txt");
        assert_eq!(part_one(input), Some(24000));
    }

    #[test]
    fn test_part_two() {
        let input = include_str!("./example.txt");
        assert_eq!(part_two(input), Some(45000));
    }
}
