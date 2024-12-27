use std::io::{self, Read};

use itertools::Itertools;
use utils::grid::{Coord, Direction, Grid};

type Input = Vec<Motion>;

#[derive(Debug, Clone)]
pub struct Rope {
    knots: Vec<Coord>,
}

impl Rope {
    fn new(knots: usize) -> Self {
        Rope {
            knots: vec![Coord::new(0, 0); knots],
        }
    }

    #[allow(dead_code)]
    fn to_str(&self) -> String {
        let (max_x, max_y) = self.knots.iter().fold((0, 0), |acc, knot| {
            (knot.x.abs().max(acc.0), knot.y.abs().max(acc.1))
        });

        let width = (self.knots.len() * 2).max(max_x as usize + 1);
        let height = (self.knots.len() * 2).max(max_y as usize + 1);
        let mut grid = Grid::new(width, height);

        for (i, knot) in self.knots.iter().enumerate() {
            let viz_knot = Coord {
                x: knot.x + width as isize / 2,
                y: knot.y + height as isize / 2,
            };
            grid.set(&viz_knot, format!("{}", i));
        }

        grid.to_str(|c| {
            if c == "" {
                ".".to_string()
            } else {
                format!("{}", c)
            }
        })
    }

    fn step(&mut self, motion: &Motion) -> Vec<Vec<Coord>> {
        let mut knot_positions = self
            .knots
            .clone()
            .into_iter()
            .map(|coord| vec![coord])
            .collect_vec();

        for _ in 0..motion.steps {
            self.knots[0] = self.knots[0].move_once(&motion.direction, 1);
            knot_positions[0].push(self.knots[0]);

            for i in 1..self.knots.len() {
                if self.knots[i].chebyshev_distance(&self.knots[i - 1]) > 1 {
                    let direction = self.knots[i].direction_to(&self.knots[i - 1]).unwrap();
                    self.knots[i] = self.knots[i].move_once(&direction, 1);
                    knot_positions[i].push(self.knots[i]);
                }
            }
        }
        knot_positions
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

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    input.lines().map(|s| Motion::from(s)).collect()
}

pub fn main() -> io::Result<()> {
    let input = parse_input();
    let mut rope = Rope::new(10);

    let tail_positions = input
        .iter()
        .map(|motion| {
            let tp = rope.step(&motion)[9].clone();
            // println!("{:?}", motion);
            // println!("{}", rope.to_str());
            // println!();

            tp
        })
        .flatten()
        .unique()
        .collect_vec();

    eprintln!("{}", tail_positions.len() as i32);
    Ok(())
}
