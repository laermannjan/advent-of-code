use utils::grid::Coord;
use utils::grid::Grid;
use utils::path_finding::shortest_path_grid;

type Input = (Grid<Square>, Coord, Coord);

#[derive(Clone, Copy, PartialEq, Eq)]
pub enum Square {
    Start,
    End,
    Height(u8),
}

impl Square {
    fn elevation(&self) -> u8 {
        match self {
            Square::Start => 0,
            Square::End => 25,
            Square::Height(x) => *x,
        }
    }
}
impl Default for Square {
    fn default() -> Self {
        Square::Height(0)
    }
}

impl std::fmt::Debug for Square {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Square::Start => write!(f, "S"),
            Square::End => write!(f, "E"),
            Square::Height(h) => write!(f, "{}", h),
        }
    }
}

impl From<char> for Square {
    fn from(c: char) -> Self {
        match c {
            'S' => Square::Start,
            'E' => Square::End,
            'a'..='z' => Square::Height(c as u8 - 'a' as u8),
            _ => panic!("bad square: {}", c),
        }
    }
}

pub fn parse_input(input: &str) -> Input {
    let mut start: Option<Coord> = None;
    let mut end: Option<Coord> = None;

    let grid = Grid::from_str(input, &mut |x, y, c| {
        if c == 'S' {
            start = Some(Coord {
                x: x as isize,
                y: y as isize,
            });
        } else if c == 'E' {
            end = Some(Coord {
                x: x as isize,
                y: y as isize,
            });
        }
        Square::from(c)
    });

    (grid, start.unwrap(), end.unwrap())
}
pub fn part_one(input: Input) -> Option<usize> {
    let (grid, start, end) = input;
    shortest_path_grid(
        &grid,
        &vec![start],
        &end,
        |curr, next| {
            (grid.get(&next).unwrap().elevation() as isize
                - grid.get(curr).unwrap().elevation() as isize)
                <= 1
        },
        |_| 1,
    )
}

pub fn part_two(input: Input) -> Option<usize> {
    let (grid, _, end) = input;

    let starts: Vec<Coord> = grid
        .coords()
        .filter(|c| [Square::Start, Square::Height(0)].contains(grid.get(&c).unwrap()))
        .collect();

    shortest_path_grid(
        &grid,
        &starts,
        &end,
        |curr, next| {
            (grid.get(&next).unwrap().elevation() as isize
                - grid.get(curr).unwrap().elevation() as isize)
                <= 1
        },
        |_| 1,
    )
}

utils::main!(2022, 12);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 12, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 12, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 12, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 12, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
