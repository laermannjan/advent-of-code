use std::collections::HashMap;

use itertools::Itertools;

#[derive(Clone, Debug, PartialEq, Eq)]
pub enum Direction {
    North,
    NorthEast,
    East,
    SouthEast,
    South,
    SouthWest,
    West,
    NorthWest,
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
pub struct Coord {
    pub x: isize,
    pub y: isize,
}

impl std::fmt::Debug for Coord {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

impl<T> From<(T, T)> for Coord
where
    T: TryInto<isize>,
{
    fn from((x, y): (T, T)) -> Self {
        Self::new(
            x.try_into()
                .unwrap_or_else(|_| panic!("number too large for architecture")),
            y.try_into()
                .unwrap_or_else(|_| panic!("number too large for architecture")),
        )
    }
}

impl Coord {
    pub fn new(x: isize, y: isize) -> Self {
        Self { x, y }
    }

    pub fn move_once(&self, direction: &Direction, steps: isize) -> Self {
        match direction {
            Direction::North => Self::new(self.x, self.y - steps),
            Direction::NorthEast => Self::new(self.x + steps, self.y - steps),
            Direction::East => Self::new(self.x + steps, self.y),
            Direction::SouthEast => Self::new(self.x + steps, self.y + steps),
            Direction::South => Self::new(self.x, self.y + steps),
            Direction::SouthWest => Self::new(self.x - steps, self.y + steps),
            Direction::West => Self::new(self.x - steps, self.y),
            Direction::NorthWest => Self::new(self.x - steps, self.y - steps),
        }
    }

    pub fn move_multiple<'a>(
        &self,
        directions: impl Iterator<Item = &'a Direction>,
        steps: isize,
    ) -> Self {
        directions.fold(self.clone(), |pos, dir| pos.move_once(&dir, steps))
    }

    pub fn path_to(&self, other: &Self) -> Vec<Self> {
        let mut pos = self.clone();
        let mut path = vec![pos];
        while pos != *other {
            let dir = pos.direction_to(&other).unwrap();
            pos = pos.move_once(&dir, 1);
            path.push(pos.clone());
        }
        path
    }

    pub fn direction_to(&self, target: &Coord) -> Option<Direction> {
        if self == target {
            return None;
        }

        let direction = if self.x == target.x {
            if self.y < target.y {
                Direction::South
            } else {
                Direction::North
            }
        } else if self.y == target.y {
            if self.x < target.x {
                Direction::East
            } else {
                Direction::West
            }
        } else if self.y < target.y {
            if self.x < target.x {
                Direction::SouthEast
            } else {
                Direction::SouthWest
            }
        } else if self.x < target.x {
            Direction::NorthEast
        } else {
            Direction::NorthWest
        };

        Some(direction)
    }

    /// Distance to target when only up, down, left, right movement is allowed
    pub fn manhattan_distance(&self, target: &Coord) -> usize {
        ((self.x - target.x).abs() + (self.y - target.y).abs()) as usize
    }

    /// Distance to target when diagonal movement is allowed
    pub fn chebyshev_distance(&self, target: &Coord) -> usize {
        (self.x - target.x).abs().max((self.y - target.y).abs()) as usize
    }
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct SparseGrid<T> {
    cells: HashMap<Coord, T>,
}

impl<T> SparseGrid<T> {
    pub fn get(&self, coord: &Coord) -> Option<&T> {
        self.cells.get(coord)
    }

    pub fn get_mut(&mut self, coord: &Coord) -> Option<&mut T> {
        self.cells.get_mut(coord)
    }

    pub fn insert(&mut self, coord: Coord, value: T) {
        self.cells.insert(coord, value);
    }

    pub fn remove(&mut self, coord: &Coord) -> Option<T> {
        self.cells.remove(coord)
    }

    pub fn contains(&self, coord: &Coord) -> bool {
        self.cells.contains_key(coord)
    }

    pub fn iter(&self) -> impl Iterator<Item = (&Coord, &T)> {
        self.cells.iter()
    }
}

impl<T> Default for SparseGrid<T> {
    fn default() -> Self {
        Self {
            cells: HashMap::new(),
        }
    }
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct Grid<T> {
    width: usize,
    height: usize,
    data: Vec<T>,
}

impl<T> Grid<T>
where
    T: Default + Clone,
{
    pub fn new(width: usize, height: usize) -> Self {
        Self {
            width,
            height,
            data: vec![T::default(); width * height],
        }
    }

    // fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
    //     let mut s = String::new();
    //     for (i, coord) in self.coords().enumerate() {
    //         if i == self.width {
    //             s.push('\n');
    //         }
    //         s.push_str(&format!("{:?}", self.get(&coord).unwrap()));
    //     }
    //     write!(f, "{}", s)
    // }

    pub fn height(&self) -> usize {
        self.height
    }

    pub fn width(&self) -> usize {
        self.width
    }

    pub fn from_str(input: &str, parser: &mut dyn FnMut(usize, usize, char) -> T) -> Self {
        let data: Vec<T> = input
            .lines()
            .enumerate()
            .map(|(y, line)| {
                line.chars()
                    .enumerate()
                    .map(|(x, c)| parser(x, y, c))
                    .collect::<Vec<T>>()
            })
            .flatten()
            .collect();
        let width = input.lines().next().unwrap().len();
        let height = data.len() / width;
        Self {
            width,
            height,
            data,
        }
    }

    pub fn to_str(&self, formatter: fn(&T) -> String) -> String {
        self.data
            .chunks(self.width)
            .map(|row| row.iter().map(formatter).collect::<String>())
            .join("\n")
    }

    pub fn in_bounds(&self, coord: &Coord) -> bool {
        0 <= coord.x
            && coord.x < self.width as isize
            && 0 <= coord.y
            && coord.y < self.height as isize
    }

    pub fn on_edge(&self, coord: &Coord) -> bool {
        coord.x == 0
            || coord.x == self.width as isize - 1
            || coord.y == 0
            || coord.y == self.height as isize - 1
    }

    pub fn get_index(&self, coord: &Coord) -> Option<usize> {
        if self.in_bounds(&coord) {
            return Some((coord.y * self.width as isize + coord.x) as usize);
        } else {
            return None;
        }
    }

    pub fn get(&self, coord: &Coord) -> Option<&T> {
        if let Some(index) = self.get_index(&coord) {
            return self.data.get(index);
        } else {
            return None;
        }
    }

    pub fn get_mut(&mut self, coord: &Coord) -> Option<&mut T> {
        if let Some(index) = self.get_index(&coord) {
            return self.data.get_mut(index);
        } else {
            return None;
        }
    }

    pub fn set(&mut self, coord: &Coord, value: T) {
        if let Some(index) = self.get_index(&coord) {
            self.data[index] = value;
        }
    }

    pub fn coords(&self) -> impl Iterator<Item = Coord> {
        (0..self.width as isize)
            .cartesian_product(0..self.height as isize)
            .map(|(x, y)| Coord::new(x, y))
    }

    pub fn step(&self, coord: &Coord, direction: &Direction) -> Option<Coord> {
        let new_coord = coord.move_once(direction, 1);

        if self.in_bounds(&new_coord) {
            return Some(new_coord);
        } else {
            return None;
        }
    }

    pub fn step_if(
        &self,
        coord: &Coord,
        direction: &Direction,
        predicate: impl Fn(&T, &T) -> bool,
    ) -> Option<Coord> {
        if let Some(new_coord) = self.step(&coord, direction) {
            let curr_cell = self.get(&coord).unwrap();
            let next_cell = self.get(&new_coord).unwrap();
            if predicate(curr_cell, next_cell) {
                return Some(new_coord);
            }
        }
        return None;
    }

    pub fn step_unless(
        &self,
        coord: &Coord,
        direction: &Direction,
        predicate: impl Fn(&T) -> bool,
    ) -> Option<Coord> {
        let mut new_coord = coord.clone();
        if let Some(coord) = self.step(&new_coord, direction) {
            new_coord = coord;
            if !predicate(self.get(&new_coord).unwrap()) {
                return Some(new_coord);
            }
        }
        return None;
    }

    pub fn walk<'a>(&'a self, coord: &Coord, direction: &'a Direction) -> WalkIterator<T> {
        WalkIterator {
            current: coord.clone(),
            direction,
            grid: self,
        }
    }
}

#[derive(Clone, Debug)]
pub struct WalkIterator<'a, T> {
    current: Coord,
    direction: &'a Direction,
    grid: &'a Grid<T>,
}

impl<'a, T> Iterator for WalkIterator<'_, T>
where
    T: Default + Clone,
{
    type Item = Coord;

    fn next(&mut self) -> Option<Self::Item> {
        if let Some(step) = self.grid.step(&self.current, self.direction) {
            self.current = step;
            return Some(self.current.clone());
        } else {
            return None;
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    mod coord {
        use super::*;
        use rstest::rstest;

        #[test]
        fn test_from() {
            let x = Coord::new(1, 2);
            assert_eq!(x.x, 1);
            assert_eq!(x.y, 2);
        }

        #[rstest]
        #[case(Direction::North, 1, Coord::new(0, -1))]
        #[case(Direction::South, 1, Coord::new(0, 1))]
        #[case(Direction::East, 1, Coord::new(1, 0))]
        #[case(Direction::West, 1, Coord::new(-1, 0))]
        #[case(Direction::North, 2, Coord::new(0, -2))]
        #[case(Direction::South, 3, Coord::new(0, 3))]
        #[case(Direction::East, 4, Coord::new(4, 0))]
        #[case(Direction::West, 5, Coord::new(-5, 0))]
        fn test_move_once(
            #[case] dir: Direction,
            #[case] steps: isize,
            #[case] expected_coord: Coord,
        ) {
            let x = Coord::new(0, 0);
            let y = x.move_once(&dir, steps);
            assert_eq!(y, expected_coord);
        }

        #[rstest]
        #[case(&[Direction::North, Direction::South], 2, Coord::new(0, 0))]
        #[case(&[Direction::North, Direction::South, Direction::East], 1,  Coord::new(1, 0))]
        fn test_move_multiple(
            #[case] directions: &[Direction],
            #[case] steps: isize,
            #[case] expected_coord: Coord,
        ) {
            let init = Coord::new(0, 0);
            let coord = init.move_multiple(directions.iter(), steps);
            assert_eq!(coord, expected_coord);
        }

        #[rstest]
        #[case(Coord::new(0, -1), Some(Direction::North))]
        #[case(Coord::new(0, 1), Some(Direction::South))]
        #[case(Coord::new(1, 0), Some(Direction::East))]
        #[case(Coord::new(-1, 0), Some(Direction::West))]
        #[case(Coord::new(0, 0), None)]
        #[case(Coord::new(1, 1), Some(Direction::SouthEast))]
        #[case(Coord::new(-1, -1), Some(Direction::NorthWest))]
        #[case(Coord::new(1, -1), Some(Direction::NorthEast))]
        #[case(Coord::new(-1, 1), Some(Direction::SouthWest))]
        fn test_direction_to(#[case] target: Coord, #[case] direction: Option<Direction>) {
            let x = Coord::new(0, 0);
            assert_eq!(x.direction_to(&target), direction);
        }

        #[rstest]
        #[case(Coord::new(0, 0), 0)]
        #[case(Coord::new(0, -1), 1)]
        #[case(Coord::new(0, 1), 1)]
        #[case(Coord::new(1, -1), 2)]
        #[case(Coord::new(-4, 4), 8)]
        fn test_manhattan_distance(#[case] target: Coord, #[case] distance: usize) {
            let init = Coord::new(0, 0);
            assert_eq!(init.manhattan_distance(&target), distance);
        }

        #[rstest]
        #[case(Coord::new(0, 0), 0)]
        #[case(Coord::new(0, -1), 1)]
        #[case(Coord::new(0, 1), 1)]
        #[case(Coord::new(1, -1), 1)]
        #[case(Coord::new(-4, 4), 4)]
        fn test_chebyshev_distance(#[case] target: Coord, #[case] distance: usize) {
            let init = Coord::new(0, 0);
            assert_eq!(init.chebyshev_distance(&target), distance);
        }
    }

    mod grid {

        use super::*;
        use rstest::{fixture, rstest};

        type TestGrid = Grid<u8>;

        #[fixture]
        fn test_grid() -> TestGrid {
            let mut grid = TestGrid::new(3, 3);
            grid.set(&Coord::new(0, 0), 1);
            grid.set(&Coord::new(1, 0), 2);
            grid.set(&Coord::new(2, 0), 3);
            grid.set(&Coord::new(0, 1), 4);
            grid.set(&Coord::new(1, 1), 5);
            grid.set(&Coord::new(2, 1), 6);
            grid.set(&Coord::new(0, 2), 7);
            grid.set(&Coord::new(1, 2), 8);
            grid.set(&Coord::new(2, 2), 9);
            assert_eq!(grid.data, vec![1, 2, 3, 4, 5, 6, 7, 8, 9]);
            grid
        }

        #[rstest]
        fn test_from_str(test_grid: TestGrid) {
            let input = "123\n456\n789";

            let grid = TestGrid::from_str(input, &mut |_, _, c| c.to_digit(10).unwrap() as u8);

            assert_eq!(grid.width, 3);
            assert_eq!(grid.height, 3);
            assert_eq!(grid.data, vec![1, 2, 3, 4, 5, 6, 7, 8, 9]);
            assert_eq!(grid, test_grid);
        }

        #[rstest]
        fn test_in_bounds(test_grid: TestGrid) {
            assert!(test_grid.in_bounds(&Coord::new(0, 0)));
            assert!(test_grid.in_bounds(&Coord::new(0, 2)));
            assert!(test_grid.in_bounds(&Coord::new(2, 0)));
            assert!(test_grid.in_bounds(&Coord::new(2, 2)));

            assert!(!test_grid.in_bounds(&Coord::new(-1, 0)));
            assert!(!test_grid.in_bounds(&Coord::new(0, -1)));
            assert!(!test_grid.in_bounds(&Coord::new(-1, -1)));

            assert!(!test_grid.in_bounds(&Coord::new(3, 0)));
            assert!(!test_grid.in_bounds(&Coord::new(0, 3)));
            assert!(!test_grid.in_bounds(&Coord::new(3, 3)));
        }

        #[rstest]
        #[case(Direction::North, Coord::new(0, 0), None)]
        #[case(Direction::South, Coord::new(0, 0), Some(Coord::new(0, 1)))]
        #[case(Direction::East, Coord::new(0, 0), Some(Coord::new(1, 0)))]
        #[case(Direction::West, Coord::new(0, 0), None)]
        #[case(Direction::NorthEast, Coord::new(0, 0), None)]
        #[case(Direction::SouthEast, Coord::new(0, 0), Some(Coord::new(1, 1)))]
        fn test_step(
            #[case] dir: Direction,
            #[case] init: Coord,
            #[case] expected: Option<Coord>,
            test_grid: TestGrid,
        ) {
            assert_eq!(test_grid.step(&init, &dir), expected);
        }

        #[rstest]
        fn test_walk(test_grid: TestGrid) {
            let init = Coord::new(0, 0);

            for (i, coord) in test_grid.walk(&init, &Direction::East).enumerate() {
                assert_eq!(coord, Coord::new((i + 1) as isize, 0));
            }
        }
    }
}
