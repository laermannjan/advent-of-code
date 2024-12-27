use std::io::{self, Read};

use itertools::Itertools;
use utils::grid::{Direction, Grid};

static DIRECTIONS: &[Direction] = &[
    Direction::North,
    Direction::East,
    Direction::South,
    Direction::West,
];

type Input = Grid<u32>;

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    Grid::from_str(&input[..], &mut |_, _, c| c.to_digit(10).unwrap())
}
pub fn main() -> io::Result<()> {
    let input = parse_input();
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

    eprintln!("{}", scenic_scores.max().unwrap() as i32);
    Ok(())
}
