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

    fn is_sensor_reachable(&self, coord: &Coord) -> bool {
        self.sensor.manhattan_distance(coord) as i32 <= self.sensor_reach()
    }

    fn scan_shell(&self, dist: usize) -> impl Iterator<Item = Coord> {
        let x = self.sensor.x as isize;
        let y = self.sensor.y as isize;

        let xmin = x - dist as isize;
        let xmax = x + dist as isize;
        let ymin = y - dist as isize;
        let ymax = y + dist as isize;

        let upper = (xmin..=xmax).zip((y..ymax).chain((y..=ymax).rev()));
        let lower = (xmin + 1..xmax).zip((ymin..y).rev().chain((ymin..y).rev()));

        upper
            .chain(lower)
            .map(|(x, y)| Coord::new(x, y))
            .into_iter()
    }

    fn scanned_cells(&self, row: Option<isize>) -> impl Iterator<Item = Coord> + '_ {
        (0..=self.sensor_reach())
            .map(|dist| self.scan_shell(dist as usize))
            .flatten()
            .filter(move |cell| row.map(|r| cell.y == r).unwrap_or(true))
            .into_iter()
    }

    fn beacon_free_cells(&self, row: Option<isize>) -> impl Iterator<Item = Coord> + '_ {
        self.scanned_cells(row)
            .into_iter()
            .inspect(|cell| {
                // dbg!(cell);
            })
            .filter(|cell| *cell != self.beacon)
            .into_iter()
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
    return None;
    let row = if cfg!(test) { 10 } else { 2_000_000 };

    // visualize(input.clone(), None);

    let free_locations = input
        .iter()
        .filter(|scan| {
            // check if sensor range affects the row
            let reach = scan.sensor_reach();
            let lower = scan.sensor.y - reach as isize;
            let upper = scan.sensor.y + reach as isize;
            (lower..=upper).contains(&row)
        })
        // .inspect(|scan| {
        //     println!("Sensor at {:?}", scan.sensor);
        //     println!("Beacon at {:?}", scan.beacon);
        //     println!("Sensor range: {}", scan.sensor_reach());
        //     println!(
        //         "Free locations: {:?}",
        //         scan.beacon_free_cells(Some(row as isize)).collect_vec()
        //     );
        // })
        .map(|s| s.beacon_free_cells(Some(row as isize)))
        .flatten()
        .unique()
        .collect_vec();

    Some(free_locations.len())
}

pub fn part_two(input: Input) -> Option<isize> {
    println!("{:?}", input.len());
    let min_coord = 0;
    let max_coord = if cfg!(test) { 20 } else { 4_000_000 };
    let coord_range = min_coord..=max_coord;

    // visualize(input.clone(), None);

    // let scanned_cells = input
    //     .iter()
    //     .map(|scan| scan.scanned_cells(None))
    //     .flatten()
    //     .unique()
    //     .collect_vec();

    let potential_distress_cell = input
        .iter()
        .map(|scan| {
            println!("reach: {:?}", scan.sensor_reach());
            scan.scan_shell((scan.sensor_reach() + 1) as usize)
        })
        .flatten()
        .filter(|cell| coord_range.contains(&cell.x) && coord_range.contains(&cell.y))
        .counts()
        .into_iter()
        .max_by_key(|(_, v)| *v)
        .map(|(k, _)| k)
        .unwrap();

    Some(potential_distress_cell.x * 4_000_000 + potential_distress_cell.y)
}

utils::main!(2022, 15);

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::*;

    #[rstest]
    #[case(None, vec![
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
        ])]
    #[case(Some(1), vec![
            Coord { x: 1, y: 1 },
            Coord { x: 2, y: 1 },
            Coord { x: 3, y: 1 },
    ])]
    #[case(Some(2), vec![
            Coord { x: 0, y: 2 },
            Coord { x: 1, y: 2 },
            Coord { x: 2, y: 2 }, // sensor
            Coord { x: 3, y: 2 },
    ])]
    fn test_beacon_free_locs(#[case] row: Option<isize>, #[case] mut expected: Vec<Coord>) {
        let input = parse_input("Sensor at x=2, y=2: closest beacon is at x=4, y=2");

        let mut free_locs = input[0].beacon_free_cells(row).collect_vec();

        free_locs.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));
        expected.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));

        assert_eq!(free_locs, expected);
    }

    #[rstest]
    #[case(0, vec![(2, 3)])]
    #[case(1, vec![(3, 3), (2, 4), (1, 3), (2, 2)])]
    fn test_scan_shell(#[case] dist: usize, #[case] expected: Vec<(isize, isize)>) {
        let input = parse_input("Sensor at x=2, y=3: closest beacon is at x=4, y=2");

        let mut shell = input[0].scan_shell(dist).collect_vec();
        let mut expected = expected
            .into_iter()
            .map(|(x, y)| Coord { x, y })
            .collect_vec();

        shell.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));
        expected.sort_by(|a, b| a.x.cmp(&b.x).then(a.y.cmp(&b.y)));

        assert_eq!(shell, expected);
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
