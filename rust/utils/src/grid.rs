use itertools::Itertools;

#[derive(Clone, Debug)]
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

impl From<(usize, usize)> for Coord {
    fn from((x, y): (usize, usize)) -> Self {
        Self::new(x as isize, y as isize)
    }
}

impl From<(isize, isize)> for Coord {
    fn from((x, y): (isize, isize)) -> Self {
        Self::new(x, y)
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

    pub fn r#move(&self, directions: &[Direction], steps: isize) -> Self {
        directions
            .iter()
            .fold(self.clone(), |pos, dir| pos.move_once(dir, steps))
    }
}

#[derive(Clone)]
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

    pub fn from_str(input: &str, parser: fn(usize, usize, char) -> T) -> Self {
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
        predicate: impl Fn(&T) -> bool,
    ) -> Option<Coord> {
        let mut new_coord = coord.clone();
        while let Some(coord) = self.step(&new_coord, direction) {
            new_coord = coord;
            if predicate(self.get(&new_coord).unwrap()) {
                return Some(new_coord);
            }
        }
        return None;
    }

    pub fn walk<'a>(
        &'a self,
        coord: &Coord,
        direction: &'a Direction,
    ) -> impl Iterator<Item = Coord> + '_ {
        let mut new_coord = coord.clone();
        std::iter::from_fn(move || {
            if let Some(coord) = self.step(&new_coord, direction) {
                new_coord = coord;
                return Some(new_coord.clone());
            } else {
                return None;
            }
        })
    }

    pub fn walk_while<'a>(
        &'a self,
        coord: &Coord,
        direction: &'a Direction,
        predicate: impl Fn(&T) -> bool + 'a,
    ) -> impl Iterator<Item = Coord> + '_ {
        let mut new_coord = coord.clone();
        std::iter::from_fn(move || {
            if let Some(coord) = self.step_if(&new_coord, direction, &predicate) {
                new_coord = coord;
                return Some(new_coord.clone());
            } else {
                return None;
            }
        })
    }
}

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
