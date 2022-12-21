extern crate nom;
use itertools::Itertools;


use nom::{
    bytes::complete::tag,
    character::complete::{i32 as parse_i32},
    IResult,
};

type Input = Vec<Scan>;

#[derive(PartialEq, Clone, Hash, Eq)]
struct Point {
    x: i32,
    y: i32,
}

impl std::fmt::Debug for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}


#[derive(PartialEq)]
struct ScanRanges {
    ranges: Vec<(i32, i32)>,
}

impl std::fmt::Debug for ScanRanges {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut s = String::new();
        for (a, b) in self.ranges.iter() {
            s.push_str(&format!("({}, {}), ", a, b));
        }
        write!(f, "[{}]", s)
    }
}

impl ScanRanges {
    fn new() -> ScanRanges {
        ScanRanges { ranges: Vec::new() }
    }

    fn len(&self) -> usize {
        self.ranges
            .iter()
            .map(|(start, end)| (end - start + 1) as usize)
            .sum::<usize>()

    }

    fn add_range(&mut self, mut new_range: (i32, i32)) {
    // Add the new range to the vector
    let mut i = 0;
    while i < self.ranges.len() {
        let curr_range = self.ranges[i];

        // If the new range is before the current range, insert it into the vector and break
        if new_range.1 < curr_range.0 {
            self.ranges.insert(i, new_range);
            break;
        }

        // If the new range is after the current range, move on to the next range
        if new_range.0 > curr_range.1 {
            i += 1;
            continue;
        }

        // If there is overlap, merge the two ranges
        let start = std::cmp::min(curr_range.0, new_range.0);
        let end = std::cmp::max(curr_range.1, new_range.1);
        let merged_range = (start, end);
        self.ranges.remove(i);
        new_range = merged_range;
    }

    // If we get to the end of the loop, add the new range to the vector
    if i == self.ranges.len() {
        self.ranges.push(new_range);
    }

    // Iterate through the vector again and merge any overlapping ranges
    i = 0;
    while i < self.ranges.len() - 1 {
        let curr_range = self.ranges[i];
        let next_range = self.ranges[i + 1];
        if curr_range.1 >= next_range.0 {
            let start = curr_range.0;
            let end = std::cmp::max(curr_range.1, next_range.1);
            let merged_range = (start, end);
            self.ranges.remove(i);
            self.ranges[i] = merged_range;
            self.ranges.remove(i + 1);
        } else {
            i += 1;
        }
    }
}
}

#[derive(Debug, Clone, PartialEq)]
struct Scan {
    sensor: Point,
    beacon: Point,
    scan_range: i32,
}

impl Scan {
    fn edge_points(&self) -> Vec<Point>{
       let x = self.sensor.x;
        let y = self.sensor.y;
        let range = self.scan_range + 1;

        let x_tops = (x - range)..=(x + range);
        let x_bots = (x - range)..=(x + range);

        let y_mid_top = y..=y+range;
        let y_top_mid = (y..y+range).rev();
        let y_tops = y_mid_top.chain(y_top_mid);

        let y_mid_bot = (y-range..=y).rev();
        let y_bot_mid = (y-range + 1)..=y;
        let y_bots = y_mid_bot.chain(y_bot_mid);

        let tops = x_tops.zip(y_tops);
        let bottoms = x_bots.zip(y_bots);
        let edge_points = tops.chain(bottoms).unique().map(|(x, y)| Point { x, y} ).collect_vec();
        edge_points
    }
    fn beacon_free_ranges(&self, y: i32, scan_ranges: &mut ScanRanges) {
        let dist = y.abs_diff(self.sensor.y) as i32;
        if dist > self.scan_range {
            return;
        }

        let scan_width = self.scan_range - dist;
        // Calculate the x-coordinates of the points in the scanned area for the given y-coordinate
        let x1 = self.sensor.x - scan_width;
        let x2 = self.sensor.x + scan_width;

        if y == self.beacon.y {
            if self.beacon.x == x1 {
                scan_ranges.add_range((x1 + 1, x2));
            } else if self.beacon.x == x2 {
                scan_ranges.add_range((x1, x2 - 1));
            } else {
                // Add the ranges on either side of the beacon if the beacon is on the same y-coordinate
                scan_ranges.add_range((x1, self.beacon.x - 1));
                scan_ranges.add_range((self.beacon.x + 1, x2));
            }
        } else {
            // Add the full range if the beacon is not on the same y-coordinate
            scan_ranges.add_range((x1, x2));
        }
    }

    fn from_string(input: &str) -> IResult<&str, Scan> {
        let (input, _) = tag("Sensor at x=")(input)?;
        let (input, x) = parse_i32(input)?;
        let (input, _) = tag(", y=")(input)?;
        let (input, y) = parse_i32(input)?;
        let (input, _) = tag(": closest beacon is at x=")(input)?;
        let (input, beacon_x) = parse_i32(input)?;
        let (input, _) = tag(", y=")(input)?;
        let (input, beacon_y) = parse_i32(input)?;

        let sensor = Point { x, y };
        let beacon = Point {
            x: beacon_x,
            y: beacon_y,
        };
        let scan_range = (x - beacon_x).abs() + (y - beacon_y).abs();
        Ok((input, Scan {
            sensor,
            beacon,
            scan_range,
        }))
    }
}

fn parse_input(input: &str) -> Vec<Scan> {
    let mut scans = Vec::new();
    for line in input.lines() {
        let scan = Scan::from_string(line).unwrap().1;
        scans.push(scan);
    }
    scans
}

fn part_one(scans: Input) -> Option<i32> {
    let y = if cfg!(test) { 10 } else { 2_000_000 };
    let mut scan_ranges = ScanRanges::new();
    for scan in scans {
        scan.beacon_free_ranges(y, &mut scan_ranges);
    }
    Some(scan_ranges.len() as i32)
}

fn part_two(scans: Input) -> Option<u64> {
    let lower = 0;
    let upper = if cfg!(test) { 20 } else { 4_000_000 };



    let counts = scans.iter().map(|scan| {
        scan.edge_points()
    }).flatten().counts();

    let distress_beacon = counts.iter().filter(|(p, _)| (lower..=upper).contains(&p.x) && (lower..=upper).contains(&p.y))
        .max_by(|a, b| a.1.cmp(&b.1))
        .map(|(p, _)| p).unwrap();

    let tuning_freq = distress_beacon.x as u64 * 4_000_000 + distress_beacon.y as u64;
    Some(tuning_freq)
}

utils::main!(2022, 15);

#[cfg(test)]
mod tests {
    use super::*;

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

    #[test]
    fn test_parse_input() {
        let input = "Sensor at x=2, y=18: closest beacon is at x=-2, y=15\nSensor at x=5, y=3: closest beacon is at x=2, y=4";
        let expected_scans = vec![
            Scan {
                sensor: Point { x: 2, y: 18 },
                beacon: Point { x: -2, y: 15 },
                scan_range: 7,
            },
            Scan {
                sensor: Point { x: 5, y: 3 },
                beacon: Point { x: 2, y: 4 },
                scan_range: 4,
            },
        ];
        assert_eq!(parse_input(input), expected_scans);
    }
    #[test]
    fn test_from_string() {
        let input = "Sensor at x=2, y=18: closest beacon is at x=-2, y=15";
        let expected_scan = Scan {
            sensor: Point { x: 2, y: 18 },
            beacon: Point { x: -2, y: 15 },
            scan_range: 7,
        };
        assert_eq!(Scan::from_string(input).unwrap().1, expected_scan);
    }

    #[test]
    fn test_add_range() {
        let mut sr = ScanRanges::new();
        sr.add_range((0, 1));
        assert_eq!(sr.ranges, [(0, 1)], "Test case 1: add range before first existing range");

        sr.add_range((2, 3));
        assert_eq!(sr.ranges, [(0, 1), (2, 3)], "Test case 2: add range after last existing range");

        sr.add_range((1, 2));
        assert_eq!(sr.ranges, [(0, 3)], "Test case 3: merge two ranges");

        sr.add_range((5, 6));
        assert_eq!(sr.ranges, [(0, 3), (5, 6)], "Test case 4: add range after last existing range");

        sr.add_range((4, 5));
        assert_eq!(sr.ranges, [(0, 3), (4, 6)], "Test case 5: merge two ranges");

        sr.add_range((3, 4));
        assert_eq!(sr.ranges, [(0, 6)], "Test case 6: merge two ranges");

        sr.add_range((-1, 0));
        assert_eq!(sr.ranges, [(-1, 6)], "Test case 7: merge two ranges");

        sr.add_range((7, 8));
        assert_eq!(sr.ranges, [(-1, 6), (7, 8)], "Test case 8: add range after last existing range");

        sr.add_range((6, 7));
        assert_eq!(sr.ranges, [(-1, 8)], "Test case 9: merge two ranges");
    }

    #[test]
fn test_beacon_free_ranges() {
        //let x1=-8 = self.sensor.x=2 - 10(self.scan_range=7 - -3(y = 15 - self.sensor.y = 18).abs());
        //let x2=12 = self.sensor.x + (self.scan_range - (y - self.sensor.y).abs());
    let y = 15;
    let scan = Scan {
        sensor: Point { x: 2, y: 18 },
        beacon: Point { x: -2, y: 15 },
        scan_range: 7,
    };
    let expected_scan_ranges = ScanRanges {
        // the beacon is on the same line as we test
        // we scan 3 lines above the sensor, so the scan_range is reduced by 3
        // thus the scan range is 4 and the range goes from -2 to 6
        // the beacon is at -2, so the range is adjusted to (-1, 6)
        ranges: vec![(-1, 6)],
    };
        let mut scan_ranges = ScanRanges::new();
        scan.beacon_free_ranges(y, &mut scan_ranges);
    assert_eq!(scan_ranges, expected_scan_ranges);

}

    #[test]
fn test_sum_lengths() {
    let scan_ranges = ScanRanges {
        ranges: vec![(1, 3), (5, 7), (9, 10)],
    };
    assert_eq!(scan_ranges.len(), 8);

    let scan_ranges = ScanRanges {
        ranges: vec![(-3, -1), (0, 1), (3, 5)],
    };
    assert_eq!(scan_ranges.len(), 8);
}
    #[test]
    fn test_edge_points(){
        let scan = Scan {
            sensor: Point { x: 0, y: 0 },
            beacon: Point { x: 0, y: 1 },
            scan_range: 1,
        };

       let expected_edge_points = vec![
            Point { x: -2, y: 0 },
            Point { x: -1, y: 1 },
            Point { x: 0, y: 2 },
            Point { x: 1, y: 1 },
            Point { x: 2, y: 0 },
            Point { x: -1, y: -1 },
            Point { x: 0, y: -2 },
            Point { x: 1, y: -1 },
        ];
        assert_eq!(scan.edge_points(), expected_edge_points);
    }
}