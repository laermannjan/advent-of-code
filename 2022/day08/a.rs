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
    // let mut s = String::new();
    // for (i, coord) in input.coords().enumerate() {
    //     if i % input.width() == 0 {
    //         s.push('\n');
    //     }
    //     s.push_str(&format!("{:?}", input.get(&coord).unwrap()));
    // }
    // println!("{}", s);

    let input = parse_input();

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

    eprintln!("{}", visible_trees as i32);
    Ok(())
}
