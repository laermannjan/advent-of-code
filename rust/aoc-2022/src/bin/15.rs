use itertools::Itertools;
use nom::{
    bytes::complete::tag,
    sequence::{preceded, separated_pair},
    IResult,
};

use utils::grid::{Coord, Grid};

type Input = Vec<Scan>;

#[derive(Debug, Clone)]
pub struct Scan {
    sensor: Coord,
    beacon: Coord,
}

#[derive(Clone, Default, PartialEq, Eq)]
pub enum ScanResult {
    Beacon,
    Sensor,
    BeaconFree,
    #[default]
    Unknown,
}

impl std::fmt::Debug for ScanResult {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            ScanResult::Beacon => write!(f, "B"),
            ScanResult::Sensor => write!(f, "S"),
            ScanResult::BeaconFree => write!(f, "#"),
            ScanResult::Unknown => write!(f, " "),
        }
    }
}

impl Scan {
    fn sensor_reach(&self) -> i32 {
        self.sensor.manhattan_distance(&self.beacon) as i32
    }

    fn scanned_cells(&self) -> Vec<Coord> {
        let sensor_range = self.sensor_reach() as isize;
        (self.sensor.x - sensor_range..=self.sensor.x + sensor_range)
            .cartesian_product(self.sensor.y - sensor_range..=self.sensor.y + sensor_range)
            .map(|(x, y)| Coord::new(x, y))
            .filter(|c| c.manhattan_distance(&self.sensor) as isize <= sensor_range)
            .collect()
    }

    fn beacon_free_cells(&self, row: Option<isize>) -> Vec<Coord> {
        self.scanned_cells()
            .into_iter()
            .filter(|cell| *cell != self.beacon)
            .filter(|cell| row.map(|r| cell.y == r).unwrap_or(true))
            .collect_vec()
    }

    fn parse_coord(input: &str) -> IResult<&str, (i32, i32)> {
        let (input, (x, y)) = separated_pair(
            preceded(tag("x="), nom::character::complete::i32),
            tag(", "),
            preceded(tag("y="), nom::character::complete::i32),
        )(input)?;
        Ok((input, (x, y)))
    }

    fn parse(input: &str) -> IResult<&str, Self> {
        let (input, _) = tag("Sensor at ")(input)?;
        let (input, (x, y)) = Scan::parse_coord(input)?;
        let sensor = Coord::new(x as isize, y as isize);
        let (input, _) = tag(": closest beacon is at ")(input)?;
        let (input, (x, y)) = Scan::parse_coord(input)?;
        let beacon = Coord::new(x as isize, y as isize);
        Ok((input, Scan { sensor, beacon }))
    }
}

pub fn parse_input(input: &str) -> Input {
    input
        .lines()
        .map(|line| Scan::parse(line).unwrap().1)
        .collect()
}

pub fn visualize(input: Input, row: Option<isize>) {
    let ((min_x, max_x), (min_y, max_y)) = input.iter().fold(
        ((usize::MAX, 0_usize), (usize::MAX, 0_usize)),
        |acc, sensor| {
            let min_x = sensor.sensor.x.min(sensor.beacon.x);
            let max_x = sensor.sensor.x.max(sensor.beacon.x);
            let min_y = sensor.sensor.y.min(sensor.beacon.y);
            let max_y = sensor.sensor.y.max(sensor.beacon.y);
            (
                (acc.0 .0.min(min_x as usize), acc.0 .1.max(max_x as usize)),
                (acc.1 .0.min(min_y as usize), acc.1 .1.max(max_y as usize)),
            )
        },
    );

    let mut grid = Grid::new((max_x - min_x + 1) as usize, (max_y - min_y + 1) as usize);

    input.iter().for_each(|sensor| {
        grid.set(
            &Coord::new(
                sensor.sensor.x - min_x as isize,
                sensor.sensor.y - min_y as isize,
            ),
            ScanResult::Sensor,
        );
        grid.set(
            &Coord::new(
                sensor.beacon.x - min_x as isize,
                sensor.beacon.y - min_y as isize,
            ),
            ScanResult::Beacon,
        );

        let free_locations = sensor.beacon_free_cells(row);

        for coord in free_locations {
            let x = coord.x - min_x as isize;
            let y = coord.y - min_y as isize;
            let grid_coord = Coord::new(x, y);
            if grid.get(&grid_coord) == Some(&ScanResult::Unknown) {
                grid.set(&grid_coord, ScanResult::BeaconFree);
            }
        }
    });

    for line in grid.to_str(|cell| format!("{:?}", cell)).lines() {
        println!("{}", line);
    }
}
pub fn part_one(input: Input) -> Option<usize> {
    let row = if cfg!(test) { 10 } else { 2_000_000 };

    // visualize(input.clone());

    let free_locations = input
        .iter()
        .filter(|scan| {
            // check if sensor range affects the row
            let reach = scan.sensor_reach();
            let lower = scan.sensor.y - reach as isize;
            let upper = scan.sensor.y + reach as isize;
            (lower..=upper).contains(&row)
        })
        .map(|s| s.beacon_free_cells(Some(row as isize)))
        .flatten()
        .unique()
        .collect_vec();

    Some(free_locations.len())
}

pub fn part_two(input: Input) -> Option<i32> {
    None
}

utils::main!(2022, 15);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_beacon_free_locs() {
        let input = parse_input("Sensor at x=2, y=2: closest beacon is at x=4, y=2");

        let mut free_locs = input[0].beacon_free_cells(None);

        let mut expected = vec![
            // up
            Coord { x: 2, y: 0 },
            // next level
            Coord { x: 1, y: 1 },
            Coord { x: 2, y: 1 },
            Coord { x: 3, y: 1 },
            // same level
            Coord { x: 0, y: 2 },
            Coord { x: 1, y: 2 },
            Coord { x: 2, y: 2 }, // sensor
            Coord { x: 3, y: 2 },
            // Coord { x: 4, y: 2 }, // beacon
            // below
            Coord { x: 1, y: 3 },
            Coord { x: 2, y: 3 },
            Coord { x: 3, y: 3 },
            // lowest
            Coord { x: 2, y: 4 },
        ];

        free_locs.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));
        expected.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));

        assert_eq!(free_locs, expected);
    }

    #[test]
    fn test_part_one() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 15, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 15, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 15, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 15, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
