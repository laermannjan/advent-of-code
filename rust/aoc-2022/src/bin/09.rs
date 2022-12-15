use itertools::Itertools;
use utils::grid::{Coord, Direction, Grid};

type Input = Vec<Motion>;

#[derive(Debug, Clone)]
pub struct Rope {
    tail: Coord,
    head: Coord,
}

impl Rope {
    fn new() -> Self {
        Rope {
            tail: Coord::new(0, 0),
            head: Coord::new(0, 0),
        }
    }

    fn step(&mut self, motion: &Motion) -> (Vec<Coord>, Vec<Coord>) {
        let mut head_positions = vec![self.head];
        let mut tail_positions = vec![self.tail];
        for _ in 0..motion.steps {
            self.head = self.head.move_once(&motion.direction, 1);
            head_positions.push(self.head);

            if self.tail.chebyshev_distance(&self.head) > 1 {
                let tail_direction = self.tail.get_direction(&self.head).unwrap();
                self.tail = self.tail.move_once(&tail_direction, 1);
                tail_positions.push(self.tail);
            }
        }
        (head_positions, tail_positions)
    }
}

#[derive(Debug, Clone)]
pub struct Motion {
    direction: Direction,
    steps: isize,
}

impl From<&str> for Motion {
    fn from(input: &str) -> Self {
        let (direction, steps) = input.split_once(" ").unwrap();
        let direction = match direction {
            "U" => Direction::North,
            "R" => Direction::East,
            "D" => Direction::South,
            "L" => Direction::West,
            _ => unreachable!(),
        };
        let steps = steps.parse().unwrap();
        Self { direction, steps }
    }
}

pub fn parse_input(input: &str) -> Input {
    input.lines().map(|s| Motion::from(s)).collect()
}
pub fn part_one(input: Input) -> Option<i32> {
    let mut rope = Rope::new();

    let tail_positions = input
        .iter()
        .map(|motion| rope.step(&motion).1)
        .flatten()
        .unique()
        .collect_vec();

    Some(tail_positions.len() as i32)
}

pub fn part_two(input: Input) -> Option<i32> {
    None
}

utils::main!(2022, 9);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 9);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 9, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 9);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 9, 2).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
