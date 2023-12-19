use utils::Day;

pub fn part_one(input: String) -> Option<i32> {
    None
}

pub fn part_two(input: String) -> Option<i32> {
    None
}

let DAY = Day {
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
        assert_eq!(part_one(input), Some(420));
    }

    #[test]
    fn test_part_two() {
        let input = include_str!("./example.txt");
        assert_eq!(part_two(input), Some(420));
    }
}

