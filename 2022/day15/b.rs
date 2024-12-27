extern crate nom;

use itertools::Itertools;
use nom::{bytes::complete::tag, character::complete::i32 as parse_i32, IResult};
use std::io::{self, Read};

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
    lower: i32,
    upper: i32,
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
    fn new(lower: i32, upper: i32) -> ScanRanges {
        ScanRanges {
            ranges: Vec::new(),
            lower,
            upper,
        }
    }

    fn len(&self) -> usize {
        self.ranges
            .iter()
            .map(|(start, end)| (end - start + 1) as usize)
            .sum::<usize>()
    }

    fn add_range(&mut self, mut new_range: (i32, i32)) {
        new_range.0 = new_range.0.max(self.lower);
        new_range.1 = new_range.1.min(self.upper);
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
        self.ranges = self
            .ranges
            .clone()
            .into_iter()
            .coalesce(|curr_range, next_range| {
                if next_range.0 <= curr_range.1 + 1 {
                    Ok((curr_range.0, next_range.1))
                } else {
                    Err((curr_range, next_range))
                }
            })
            .collect_vec();
    }
}

#[derive(Clone, PartialEq)]
struct Edge {
    start: Point,
    end: Point,
}

impl std::fmt::Debug for Edge {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "<{:?} -> {:?}>", self.start, self.end)
    }
}

#[derive(Debug, Clone, PartialEq)]
struct Scan {
    sensor: Point,
    beacon: Point,
    scan_reach: i32,
}

impl Scan {
    fn extract_scan_ranges(&self, y: i32, scan_ranges: &mut ScanRanges, exclude_beacon: bool) {
        let dist = y.abs_diff(self.sensor.y) as i32;
        if dist > self.scan_reach {
            return;
        }

        let scan_width = self.scan_reach - dist;
        // Calculate the x-coordinates of the points in the scanned area for the given y-coordinate
        let x1 = self.sensor.x - scan_width;
        let x2 = self.sensor.x + scan_width;

        if exclude_beacon && y == self.beacon.y {
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
        Ok((
            input,
            Scan {
                sensor,
                beacon,
                scan_reach: scan_range,
            },
        ))
    }
}

fn parse_input() -> Vec<Scan> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    let mut scans = Vec::new();
    for line in input.lines() {
        let scan = Scan::from_string(line).unwrap().1;
        scans.push(scan);
    }
    scans
}

fn main() -> io::Result<()> {
    let scans = parse_input();
    let lower = 0;
    let upper = if cfg!(test) { 20 } else { 4_000_000 };

    for y in lower..=upper {
        let mut scan_ranges = ScanRanges::new(lower, upper);
        for scan in &scans {
            // accumulate all the ranges for the current y-coordinate, not ignoring the beacon
            scan.extract_scan_ranges(y, &mut scan_ranges, false);
        }
        if scan_ranges.ranges.len() > 1 {
            // Since there will only be exactly one point that is not scanned by any sensor
            // ff there are more than one range, we have found the y-coordinate
            // the gap is between the two ranges and will be the adajcent point (x-coordinate + 1)
            eprintln!(
                "{}",
                y as i64 + (scan_ranges.ranges[0].1 as i64 + 1) * 4_000_000
            );
            return Ok(());
        }
    }
    return Ok(());
}
