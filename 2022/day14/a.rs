use itertools::Itertools;
use std::io::{self, Read};
use utils::grid::{Coord, Direction, SparseGrid};

type Grid = SparseGrid<Content>;
type Input = (Grid, isize);

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum Content {
    Rock,
    Sand,
}

static SOURCE: Coord = Coord { x: 500, y: 0 };

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    input
        .lines()
        .fold((SparseGrid::default(), 0_isize), |mut acc, l| {
            l.split(" -> ")
                .filter_map(|coord| {
                    let (x, y) = coord.split_once(',')?;
                    Some(Coord {
                        x: x.parse().ok()?,
                        y: y.parse().ok()?,
                    })
                })
                .tuple_windows()
                .for_each(|(a, b)| {
                    for c in a.path_to(&b) {
                        acc.1 = acc.1.max(c.y);
                        acc.0.insert(c, Content::Rock);
                    }
                });
            acc
        })
}

pub fn fall(grid: &Grid, sand_unit: &Coord) -> Option<Coord> {
    for d in vec![Direction::South, Direction::SouthWest, Direction::SouthEast] {
        let next_coord = sand_unit.move_once(&d, 1);
        if grid.get(&next_coord).is_none() {
            return Some(next_coord);
        }
    }
    None
}

pub fn main() -> io::Result<()> {
    let input = parse_input();
    let (mut grid, cave_depth) = input;

    let mut sand_path = vec![SOURCE.clone()];
    loop {
        let sand_unit = &sand_path[sand_path.len() - 1];
        match fall(&grid, &sand_unit) {
            Some(next_coord) => {
                if next_coord.y < cave_depth {
                    sand_path.push(next_coord);
                } else {
                    break;
                }
            }
            None => {
                grid.insert(sand_unit.clone(), Content::Sand);
                sand_path.pop();
            }
        }
    }

    let mut sand_units = grid
        .iter()
        .filter(|(_, v)| v == &&Content::Sand)
        .map(|(k, _)| k)
        .collect_vec();

    sand_units.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));

    eprintln!("{}", sand_units.len());
    Ok(())
}
