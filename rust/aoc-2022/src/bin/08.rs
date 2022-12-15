use itertools::Itertools;
use utils::grid::{Coord, Direction, Grid};

static DIRECTIONS: &[Direction] = &[
    Direction::North,
    Direction::East,
    Direction::South,
    Direction::West,
];

type Input = Grid<u32>;

pub fn parse_input(input: &str) -> Input {
    Grid::from_str(input, |_, _, c| c.to_digit(10).unwrap())
}
pub fn part_one(input: Input) -> Option<i32> {
    // let mut s = String::new();
    // for (i, coord) in input.coords().enumerate() {
    //     if i % input.width() == 0 {
    //         s.push('\n');
    //     }
    //     s.push_str(&format!("{:?}", input.get(&coord).unwrap()));
    // }
    // println!("{}", s);

    let visible_trees = input
        .coords()
        .filter(|coord| {
            if input.on_edge(&coord) {
                true
            } else {
                let size = input.get(coord).unwrap();
                DIRECTIONS.iter().any(|direction| {
                    input
                        .walk(coord, direction)
                        .all(|other| input.get(&other).unwrap() < size)
                })
            }
        })
        .count();

    Some(visible_trees as i32)
}

pub fn part_two(input: Input) -> Option<i32> {
    let scenic_scores = input.coords().map(|coord| {
        DIRECTIONS
            .iter()
            .map(|direction| {
                let mut trees_in_line = input.walk(&coord, direction);
                let mut visible_trees = trees_in_line
                    .take_while_ref(|other| input.get(&other).unwrap() < input.get(&coord).unwrap())
                    .collect_vec();

                if let Some(coord) = trees_in_line.next() {
                    visible_trees.push(coord);
                }
                visible_trees.len()
            })
            .product::<usize>()
    });

    Some(scenic_scores.max().unwrap() as i32)
}

utils::main!(2022, 8);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 8);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 8, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 8);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 8, 2).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
