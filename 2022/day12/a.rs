use std::io::{self, Read};
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

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    let mut start: Option<Coord> = None;
    let mut end: Option<Coord> = None;

    let grid = Grid::from_str(&input[..], &mut |x, y, c| {
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
pub fn main() -> io::Result<()> {
    let input = parse_input();
    let (grid, start, end) = input;
    eprintln!(
        "{}",
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
        .unwrap()
    );
    Ok(())
}
