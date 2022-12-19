use itertools::Itertools;
use utils::grid::{Coord, Direction, SparseGrid};

type Grid = SparseGrid<Content>;
type Input = (Grid, isize);

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum Content {
    Rock,
    Sand,
}

static SOURCE: Coord = Coord { x: 500, y: 0 };

pub fn parse_input(input: &str) -> Input {
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

pub fn part_one(input: Input) -> Option<usize> {
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

    Some(sand_units.len())
}

pub fn part_two(input: Input) -> Option<usize> {
    let (mut grid, cave_depth) = input;

    let mut sand_path = vec![SOURCE.clone()];
    loop {
        let sand_unit = &sand_path[sand_path.len() - 1];

        if let Some(next_coord) = fall(&grid, &sand_unit) {
            if next_coord.y < cave_depth + 2 {
                sand_path.push(next_coord);
                continue;
            }
        }

        grid.insert(sand_unit.clone(), Content::Sand);
        if sand_unit == &SOURCE {
            break;
        } else {
            sand_path.pop();
        }
    }

    let mut sand_units = grid
        .iter()
        .filter(|(_, v)| v == &&Content::Sand)
        .map(|(k, _)| k)
        .collect_vec();

    sand_units.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));

    Some(sand_units.len())
}

utils::main!(2022, 14);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 14, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 14, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 14, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 14, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
